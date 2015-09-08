package raft

import (
	"bytes"
	"fmt"

	"golang.org/x/net/context"
)

func newTestConfig(id uint64, peers []uint64, election, heartbeat int, storage Storage) *Config {
	return &Config{
		ID:              id,
		peers:           peers,
		ElectionTick:    election,
		HeartbeatTick:   heartbeat,
		Storage:         storage,
		MaxSizePerMsg:   noLimit,
		MaxInflightMsgs: 256,
	}
}

func Fuzz(data []byte) int {
	mn := newMultiNode(1)
	go mn.run()
	s := NewMemoryStorage()
	mn.CreateGroup(1, newTestConfig(1, nil, 10, 1, s), []Peer{{ID: 1}})
	mn.Campaign(context.TODO(), 1)
	proposed := false
	for {
		rds := <-mn.Ready()
		rd := rds[1]
		s.Append(rd.Entries)
		// Once we are the leader, propose a command.
		if !proposed && rd.SoftState.Lead == mn.id {
			mn.Propose(context.TODO(), 1, data)
			proposed = true
		}
		mn.Advance(rds)

		// Exit when we have three entries: one ConfChange, one no-op for the election,
		// and our proposed command.
		lastIndex, err := s.LastIndex()
		if err != nil {
			panic(err)
		}
		if lastIndex >= 3 {
			break
		}
	}
	mn.Stop()

	lastIndex, err := s.LastIndex()
	if err != nil {
		panic(err)
	}
	entries, err := s.Entries(lastIndex, lastIndex+1, noLimit)
	if err != nil {
		panic(err)
	}
	if len(entries) != 1 {
		panic(fmt.Errorf("len(entries) = %d, want %d", len(entries), 1))
	}
	if !bytes.Equal(entries[0].Data, data) {
		panic(fmt.Errorf("entries[0].Data = %v, want %v", entries[0].Data, data))
	}
	return 1
}
