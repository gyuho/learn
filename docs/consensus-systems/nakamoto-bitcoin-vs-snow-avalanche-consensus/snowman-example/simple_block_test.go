package example

import (
	"testing"
	"time"

	"github.com/ava-labs/avalanchego/ids"

	"github.com/stretchr/testify/require"
)

func TestBlock(t *testing.T) {
	require := require.New(t)

	blkID := ids.GenerateTestID()
	x := NewBlock(blkID, ids.Empty, 0, time.Now())
	require.Equal(x.ID(), blkID)

	encoded := x.Bytes()
	y, err := ParseBlock(encoded)
	require.NoError(err)

	require.Equal(x, y)
}
