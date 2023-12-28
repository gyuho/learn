package example

import (
	"context"
	"encoding/json"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/choices"
	"github.com/ava-labs/avalanchego/snow/consensus/snowman"
)

var _ snowman.Block = (*simpleBlock)(nil)

func NewBlock(id ids.ID, parent ids.ID, height uint64, timestamp time.Time) snowman.Block {
	return &simpleBlock{
		Id:         id,
		ParentId:   parent,
		Hght:       height,
		BlkTimeUtc: timestamp.UTC(),
		BlkStatus:  choices.Processing,
	}
}

func ParseBlock(b []byte) (snowman.Block, error) {
	blk := &simpleBlock{}
	if err := json.Unmarshal(b, blk); err != nil {
		return nil, err
	}
	return blk, nil
}

type simpleBlock struct {
	Id        ids.ID         `json:"id"`
	Accepted  bool           `json:"accepted"`
	Rejected  bool           `json:"rejected"`
	BlkStatus choices.Status `json:"block_status"`

	ParentId   ids.ID    `json:"parent_id"`
	Hght       uint64    `json:"height"`
	BlkTimeUtc time.Time `json:"block_time_utc"`
}

func (b *simpleBlock) ID() ids.ID {
	return b.Id
}

func (b *simpleBlock) Accept(context.Context) error {
	b.Accepted = true
	b.Rejected = false
	b.BlkStatus = choices.Accepted
	return nil
}

func (b *simpleBlock) Reject(context.Context) error {
	b.Accepted = false
	b.Rejected = true
	b.BlkStatus = choices.Rejected

	// TODO: error if we go from accept to reject?
	return nil
}

func (b *simpleBlock) Status() choices.Status {
	return b.BlkStatus
}

func (b *simpleBlock) Parent() ids.ID {
	return b.ParentId
}

func (b *simpleBlock) Verify(context.Context) error {
	return nil
}

func (b *simpleBlock) Bytes() []byte {
	encoded, err := json.Marshal(b)
	if err != nil {
		return nil
	}
	return encoded
}

func (b *simpleBlock) Height() uint64 {
	return b.Hght
}

func (b *simpleBlock) Timestamp() time.Time {
	return b.BlkTimeUtc
}
