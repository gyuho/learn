package rafthttp

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/coreos/etcd/etcdserver/stats"
	"github.com/coreos/etcd/pkg/types"
	"github.com/coreos/etcd/raft/raftpb"
)

func Fuzz(data []byte) int {
	tests := []raftpb.Message{
		linkHeartbeatMessage,
		{
			Type:    raftpb.MsgApp,
			From:    1,
			To:      2,
			Term:    1,
			LogTerm: 1,
			Index:   0,
			Entries: []raftpb.Entry{
				{Term: 1, Index: 1, Data: data},
				{Term: 1, Index: 2, Data: data},
				{Term: 1, Index: 3, Data: data},
			},
		},
		// consecutive MsgApp
		{
			Type:    raftpb.MsgApp,
			From:    1,
			To:      2,
			Term:    1,
			LogTerm: 1,
			Index:   3,
			Entries: []raftpb.Entry{
				{Term: 1, Index: 4, Data: data},
			},
		},
		linkHeartbeatMessage,
		// consecutive MsgApp after linkHeartbeatMessage
		{
			Type:    raftpb.MsgApp,
			From:    1,
			To:      2,
			Term:    1,
			LogTerm: 1,
			Index:   4,
			Entries: []raftpb.Entry{
				{Term: 1, Index: 5, Data: data},
			},
		},
		// MsgApp with higher term
		{
			Type:    raftpb.MsgApp,
			From:    1,
			To:      2,
			Term:    3,
			LogTerm: 1,
			Index:   5,
			Entries: []raftpb.Entry{
				{Term: 3, Index: 6, Data: data},
			},
		},
		linkHeartbeatMessage,
		// consecutive MsgApp
		{
			Type:    raftpb.MsgApp,
			From:    1,
			To:      2,
			Term:    3,
			LogTerm: 2,
			Index:   6,
			Entries: []raftpb.Entry{
				{Term: 3, Index: 7, Data: data},
			},
		},
		// consecutive empty MsgApp
		{
			Type:    raftpb.MsgApp,
			From:    1,
			To:      2,
			Term:    3,
			LogTerm: 2,
			Index:   7,
			Entries: nil,
		},
		linkHeartbeatMessage,
	}
	b := &bytes.Buffer{}
	enc := newMsgAppV2Encoder(b, &stats.FollowerStats{})
	dec := newMsgAppV2Decoder(b, types.ID(2), types.ID(1))

	for i, tt := range tests {
		if err := enc.encode(tt); err != nil {
			panic(fmt.Errorf("#%d: unexpected encode message error: %v", i, err))
			continue
		}
		m, err := dec.decode()
		if err != nil {
			panic(fmt.Errorf("#%d: unexpected decode message error: %v", i, err))
			continue
		}
		if !reflect.DeepEqual(m, tt) {
			panic(fmt.Errorf("#%d: message = %+v, want %+v", i, m, tt))
		}
	}
	return 1
}
