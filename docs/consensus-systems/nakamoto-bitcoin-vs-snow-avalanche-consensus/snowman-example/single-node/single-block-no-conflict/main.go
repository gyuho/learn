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
		Use:   "single-block-no-conflict",
		Short: "Simplest snowman example of a single block with no conflict",
		Long:  "Only one choice is in consensus with no conflict, which is finalized after beta virtuous rounds.",
		RunE:  runRunc,
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
	blk1 := example.NewBlock(ids.ID{1}, runner.GenesisBlk.ID(), runner.GenesisBlk.Height()+1, now)

	runner.Ctx.Log.Info("adding a single block to consensus",
		zap.Any("block1", blk1.ID()),
	)
	if err := runner.Consensus.Add(rootCtx, blk1); err != nil {
		return fmt.Errorf("failed to add block: %w", err)
	}

	// blk1 has no conflict, so it should finalize in beta virtuous
	finalizePolls := 0
	for i := 0; i < 2*runner.Params.BetaVirtuous; i++ { // 2*beta rounds to make sure it can terminate earlier
		votes := bag.Bag[ids.ID]{}
		votes.AddCount(blk1.ID(), runner.Params.Alpha)
		runner.Ctx.Log.Info("voting", zap.Any("blockId", blk1.ID()), zap.Any("polls", finalizePolls+1))

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
	if finalizePolls != runner.Params.BetaVirtuous {
		return fmt.Errorf("consensus should be finalized exactly after beta virtuous polls, took %d polls", finalizePolls)
	}

	if runner.Consensus.Preference() != blk1.ID() {
		return errors.New("consensus should prefer blk1")
	}
	if runner.Consensus.LastAccepted() != blk1.ID() {
		return errors.New("consensus should have accepted blk1")
	}
	if blk1.Status() != choices.Accepted {
		return errors.New("blk1 should have been accepted")
	}

	return nil
}
