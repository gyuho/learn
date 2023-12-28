package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/choices"
	"github.com/ava-labs/avalanchego/snow/consensus/snowball"
	"github.com/ava-labs/avalanchego/snow/consensus/snowman"
	"github.com/ava-labs/avalanchego/snow/consensus/snowman/example"
	"github.com/ava-labs/avalanchego/utils/bag"
	"github.com/ava-labs/avalanchego/utils/logging"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func init() {
	cobra.EnablePrefixMatching = true
}

var (
	logLvl  string
	timeout time.Duration
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "transitive-vote",
		Short: "Transitive voting example",
		Long: `
Voting for child blocks transitively endorses all the blocks in its transitive path.
Transitive voting rejects all blocks outside of its transitive paths.
Parent blocks do not need direct votes to be accepted, as long as its child blocks receive enough votes.

(ref. 'RecordPollRejectTransitivelyTest' and 'RecordPollTransitivelyResetConfidenceTest').
`,
		RunE: runRunc,
	}

	rootCmd.PersistentFlags().StringVar(&logLvl, "log-level", logging.Verbo.String(), "log level")
	rootCmd.PersistentFlags().DurationVar(&timeout, "timeout", time.Hour, "timeout for command")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "failed %v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}

func runRunc(cmd *cobra.Command, args []string) error {
	lvl, err := logging.ToLevel(logLvl)
	if err != nil {
		return err
	}
	rootCtx, rootCancel := context.WithTimeout(context.Background(), timeout)
	defer rootCancel()

	runner, err := example.New(lvl, snowball.DefaultParameters)
	if err != nil {
		return err
	}

	now := time.Now()

	// Current graph structure:
	//   G
	//  / \
	// 1   2
	//    / \
	//   3   4
	blk1 := example.NewBlock(ids.ID{1}, runner.GenesisBlk.ID(), runner.GenesisBlk.Height()+1, now)
	blk2 := example.NewBlock(ids.ID{2}, runner.GenesisBlk.ID(), runner.GenesisBlk.Height()+1, now)
	blk3 := example.NewBlock(ids.ID{3}, blk2.ID(), blk2.Height()+1, now)
	blk4 := example.NewBlock(ids.ID{4}, blk2.ID(), blk2.Height()+1, now)

	for _, blk := range []snowman.Block{blk1, blk2, blk3, blk4} {
		runner.Ctx.Log.Info("adding", zap.Any("blockId", blk.ID()))
		if err := runner.Consensus.Add(rootCtx, blk); err != nil {
			return fmt.Errorf("failed to add block: %w", err)
		}
	}

	// blk3 is in conflict with blk4, so it requires beta rogue rounds to finalize
	// if voted less than beta rogue, it should not finalize
	for i := 0; i < runner.Params.BetaRogue/2; i++ {
		votesFor3 := bag.Bag[ids.ID]{}
		votesFor3.AddCount(blk3.ID(), runner.Params.Alpha)

		if err := runner.Consensus.RecordPoll(rootCtx, votesFor3); err != nil {
			return fmt.Errorf("failed to record poll: %w", err)
		}
	}
	if runner.Consensus.NumProcessing() == 0 {
		return errors.New("consensus should not be finalized yet")
	}
	if runner.Consensus.Preference() != blk3.ID() {
		return fmt.Errorf("expected preference to be blk3, got %s", runner.Consensus.Preference())
	}

	// note that this generates consecutive votes for the blk4
	// without its confidence reset
	finalizePolls := 0
	for i := 0; i < 2*runner.Params.BetaRogue; i++ { // 2*beta rounds to make sure it can terminate earlier
		votes := bag.Bag[ids.ID]{}
		votes.AddCount(blk4.ID(), runner.Params.Alpha)
		runner.Ctx.Log.Info("voting", zap.Any("blockId", blk4.ID()), zap.Any("polls", finalizePolls+1))

		if err := runner.Consensus.RecordPoll(rootCtx, votes); err != nil {
			return fmt.Errorf("failed to record poll: %w", err)
		}
		finalizePolls++

		runner.Ctx.Log.Info("status",
			zap.Any("preference", runner.Consensus.Preference()),
			zap.Any("lastAccepted", runner.Consensus.LastAccepted()),
			zap.Any("finalized", runner.Consensus.NumProcessing() == 0),
		)

		if runner.Consensus.NumProcessing() == 0 {
			break
		}
	}
	if runner.Consensus.NumProcessing() > 0 {
		return errors.New("consensus should be finalized")
	}
	if finalizePolls != runner.Params.BetaRogue {
		return fmt.Errorf("consensus should be finalized exactly after beta rogue polls, took %d polls", finalizePolls)
	}

	if runner.Consensus.Preference() != blk4.ID() {
		return errors.New("consensus should prefer blk4")
	}
	if runner.Consensus.LastAccepted() != blk4.ID() {
		return errors.New("consensus should have accepted blk4")
	}
	if blk1.Status() != choices.Rejected {
		return errors.New("blk1 should have been rejected")
	}
	if blk2.Status() != choices.Accepted {
		return errors.New("blk2 should have been accepted")
	}
	if blk3.Status() != choices.Rejected {
		return errors.New("blk3 should have been rejected")
	}
	if blk4.Status() != choices.Accepted {
		return errors.New("blk4 should have been accepted")
	}

	return nil
}
