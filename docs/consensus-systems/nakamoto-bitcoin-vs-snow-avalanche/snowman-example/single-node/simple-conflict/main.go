package main

import (
	"context"
	"crypto/rand"
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
		Use:   "simple-conflict",
		Short: "Single block with a conflict between two children",
		Long:  "Only two child blocks in conflict, one of which with alpha votes with beta rogue rounds is chosen and finalized.",
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

	blk1ID := ids.ID{}
	if _, err := rand.Read(blk1ID[:]); err != nil {
		return err
	}
	blk2ID := ids.ID{}
	copy(blk2ID[:], blk1ID[:])
	blk2ID[0] ^= byte(0xFF)

	now := time.Now()
	blk1 := example.NewBlock(blk1ID, runner.GenesisBlk.ID(), runner.GenesisBlk.Height()+1, now)
	blk2 := example.NewBlock(blk2ID, runner.GenesisBlk.ID(), runner.GenesisBlk.Height()+1, now)

	runner.Ctx.Log.Info("adding two blocks to consensus",
		zap.Any("block1", blk1.ID()),
		zap.Any("block2", blk2.ID()),
	)
	if err := runner.Consensus.Add(rootCtx, blk1); err != nil {
		return fmt.Errorf("failed to add block: %w", err)
	}
	if err := runner.Consensus.Add(rootCtx, blk2); err != nil {
		return fmt.Errorf("failed to add block: %w", err)
	}

	// blk1 and blk2 are in conflict, so need beta rogue threshold to finalize
	// note that this generates consecutive votes for the blk1
	// without its confidence reset
	finalizePolls := 0
	for i := 0; i < 2*runner.Params.BetaRogue; i++ { // 2*beta rounds to make sure it can terminate earlier
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
	if finalizePolls != runner.Params.BetaRogue {
		return fmt.Errorf("consensus should be finalized exactly after beta rogue polls, took %d polls", finalizePolls)
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
	if blk2.Status() != choices.Rejected {
		return errors.New("blk2 should have been rejected")
	}

	return nil
}
