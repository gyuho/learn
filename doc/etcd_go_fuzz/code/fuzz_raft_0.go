package raft

import (
	"reflect"

	pb "github.com/coreos/etcd/raft/raftpb"
)

func Fuzz(data []byte) int {
	ents := []pb.Entry{{Index: 3, Term: 3}, {Index: 4, Term: 4}, {Index: 5, Term: 5}}
	cs := &pb.ConfState{Nodes: []uint64{1, 2, 3}}
	tests := []struct {
		i uint64

		werr  error
		wsnap pb.Snapshot
	}{
		{4, nil, pb.Snapshot{Data: data, Metadata: pb.SnapshotMetadata{Index: 4, Term: 4, ConfState: *cs}}},
		{5, nil, pb.Snapshot{Data: data, Metadata: pb.SnapshotMetadata{Index: 5, Term: 5, ConfState: *cs}}},
	}
	for _, tt := range tests {
		s := &MemoryStorage{ents: ents}
		snap, err := s.CreateSnapshot(tt.i, cs, data)
		if err != tt.werr {
			panic(err)
		}
		if !reflect.DeepEqual(snap, tt.wsnap) {
			panic("reflect.DeepEqual")
		}
	}
	return 1
}
