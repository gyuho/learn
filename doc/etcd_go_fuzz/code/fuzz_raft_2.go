package raft

import (
	"fmt"

	pb "github.com/coreos/etcd/raft/raftpb"
)

// nextEnts returns the appliable entries and updates the applied index
func nextEnts(r *raft, s *MemoryStorage) (ents []pb.Entry) {
	// Transfer all unstable entries to "stable" storage.
	s.Append(r.raftLog.unstableEntries())
	r.raftLog.stableTo(r.raftLog.lastIndex(), r.raftLog.lastTerm())

	ents = r.raftLog.nextEnts()
	r.raftLog.appliedTo(r.raftLog.committed)
	return ents
}

type Interface interface {
	Step(m pb.Message) error
	readMessages() []pb.Message
}

func (r *raft) readMessages() []pb.Message {
	msgs := r.msgs
	r.msgs = make([]pb.Message, 0)

	return msgs
}

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

func newTestRaft(id uint64, peers []uint64, election, heartbeat int, storage Storage) *raft {
	return newRaft(newTestConfig(id, peers, election, heartbeat, storage))
}

func Fuzz(data []byte) int {
	r := newTestRaft(1, []uint64{1, 2}, 5, 1, NewMemoryStorage())
	r.becomeCandidate()
	r.becomeLeader()

	pr2 := r.prs[2]
	// force the progress to be in replicate state
	pr2.becomeReplicate()
	// fill in the inflights window
	for i := 0; i < r.maxInflight; i++ {
		r.Step(pb.Message{From: 1, To: 1, Type: pb.MsgProp, Entries: []pb.Entry{{Data: data}}})
		ms := r.readMessages()
		if len(ms) != 1 {
			panic(fmt.Errorf("#%d: len(ms) = %d, want 1", i, len(ms)))
		}
	}

	// ensure 1
	if !pr2.ins.full() {
		panic(fmt.Errorf("inflights.full = %t, want %t", pr2.ins.full(), true))
	}

	// ensure 2
	for i := 0; i < 10; i++ {
		r.Step(pb.Message{From: 1, To: 1, Type: pb.MsgProp, Entries: []pb.Entry{{Data: data}}})
		ms := r.readMessages()
		if len(ms) != 0 {
			panic(fmt.Errorf("#%d: len(ms) = %d, want 0", i, len(ms)))
		}
	}

	r = newTestRaft(1, []uint64{1, 2}, 5, 1, NewMemoryStorage())
	r.becomeCandidate()
	r.becomeLeader()

	pr2 = r.prs[2]
	// force the progress to be in replicate state
	pr2.becomeReplicate()
	// fill in the inflights window
	for i := 0; i < r.maxInflight; i++ {
		r.Step(pb.Message{From: 1, To: 1, Type: pb.MsgProp, Entries: []pb.Entry{{Data: data}}})
		r.readMessages()
	}

	// 1 is noop, 2 is the first proposal we just sent.
	// so we start with 2.
	for tt := 2; tt < r.maxInflight; tt++ {
		// move forward the window
		r.Step(pb.Message{From: 2, To: 1, Type: pb.MsgAppResp, Index: uint64(tt)})
		r.readMessages()

		// fill in the inflights window again
		r.Step(pb.Message{From: 1, To: 1, Type: pb.MsgProp, Entries: []pb.Entry{{Data: data}}})
		ms := r.readMessages()
		if len(ms) != 1 {
			panic(fmt.Errorf("#%d: len(ms) = %d, want 1", tt, len(ms)))
		}

		// ensure 1
		if !pr2.ins.full() {
			panic(fmt.Errorf("inflights.full = %t, want %t", pr2.ins.full(), true))
		}

		// ensure 2
		for i := 0; i < tt; i++ {
			r.Step(pb.Message{From: 2, To: 1, Type: pb.MsgAppResp, Index: uint64(i)})
			if !pr2.ins.full() {
				panic(fmt.Errorf("#%d: inflights.full = %t, want %t", tt, pr2.ins.full(), true))
			}
		}
	}

	r = newTestRaft(1, []uint64{1, 2}, 5, 1, NewMemoryStorage())
	r.becomeCandidate()
	r.becomeLeader()

	pr2 = r.prs[2]
	// force the progress to be in replicate state
	pr2.becomeReplicate()
	// fill in the inflights window
	for i := 0; i < r.maxInflight; i++ {
		r.Step(pb.Message{From: 1, To: 1, Type: pb.MsgProp, Entries: []pb.Entry{{Data: data}}})
		r.readMessages()
	}

	for tt := 1; tt < 5; tt++ {
		if !pr2.ins.full() {
			panic(fmt.Errorf("#%d: inflights.full = %t, want %t", tt, pr2.ins.full(), true))
		}

		// recv tt msgHeartbeatResp and expect one free slot
		for i := 0; i < tt; i++ {
			r.Step(pb.Message{From: 2, To: 1, Type: pb.MsgHeartbeatResp})
			r.readMessages()
			if pr2.ins.full() {
				panic(fmt.Errorf("#%d.%d: inflights.full = %t, want %t", tt, i, pr2.ins.full(), false))
			}
		}

		// one slot
		r.Step(pb.Message{From: 1, To: 1, Type: pb.MsgProp, Entries: []pb.Entry{{Data: data}}})
		ms := r.readMessages()
		if len(ms) != 1 {
			panic(fmt.Errorf("#%d: free slot = 0, want 1", tt))
		}

		// and just one slot
		for i := 0; i < 10; i++ {
			r.Step(pb.Message{From: 1, To: 1, Type: pb.MsgProp, Entries: []pb.Entry{{Data: data}}})
			ms1 := r.readMessages()
			if len(ms1) != 0 {
				panic(fmt.Errorf("#%d.%d: len(ms) = %d, want 0", tt, i, len(ms1)))
			}
		}

		// clear all pending messages.
		r.Step(pb.Message{From: 2, To: 1, Type: pb.MsgHeartbeatResp})
		r.readMessages()
	}
	return 1
}
