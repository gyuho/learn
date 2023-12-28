package example

import (
	"context"
	"crypto/rand"
	"errors"
	"os"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/snow/consensus/snowball"
	"github.com/ava-labs/avalanchego/snow/consensus/snowman"
	"github.com/ava-labs/avalanchego/utils/logging"

	"go.uber.org/zap"
)

type Instance struct {
	Ctx        *snow.ConsensusContext
	Params     snowball.Parameters
	Consensus  snowman.Consensus
	GenesisBlk snowman.Block
}

// Creates a new snowman consensus instance with the provided parameters
// with an automatically generated genesis block.
func New(logLvl logging.Level, sbParams snowball.Parameters) (Instance, error) {
	consensus := &snowman.Topological{}

	ctx := snow.DefaultConsensusContextTest()
	ctx.Log = logging.NewLogger("", logging.NewWrappedCore(logLvl, os.Stdout, logging.Colors.ConsoleEncoder()))

	genesisBlkID := ids.ID{}
	if _, err := rand.Read(genesisBlkID[:]); err != nil {
		return Instance{}, err
	}

	genesisBlkHght := uint64(0)
	genesisBlk := NewBlock(genesisBlkID, ids.Empty, genesisBlkHght, time.Now())
	if err := genesisBlk.Accept(context.Background()); err != nil {
		return Instance{}, err
	}

	if err := consensus.Initialize(ctx, sbParams, genesisBlk.ID(), genesisBlk.Height(), genesisBlk.Timestamp()); err != nil {
		return Instance{}, err
	}
	ctx.Log.Info("initialized snowman", zap.Any("genesisBlk", genesisBlk))

	if consensus.LastAccepted() != genesisBlkID {
		return Instance{}, errors.New("last accepted block should be genesis block")
	}

	return Instance{
		Ctx:        ctx,
		Params:     sbParams,
		Consensus:  consensus,
		GenesisBlk: genesisBlk,
	}, nil
}
