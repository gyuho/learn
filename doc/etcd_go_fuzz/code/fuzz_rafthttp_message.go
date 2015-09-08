package rafthttp

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/coreos/etcd/raft/raftpb"
)

func Fuzz(data []byte) int {
	tests := []raftpb.Message{
		{
			Type:    raftpb.MsgApp,
			From:    1,
			To:      2,
			Term:    1,
			LogTerm: 1,
			Index:   3,
			Entries: []raftpb.Entry{{Term: 1, Index: 4}},
		},
		{
			Type: raftpb.MsgProp,
			From: 1,
			To:   2,
			Entries: []raftpb.Entry{
				{Data: data},
				{Data: data},
				{Data: data},
			},
		},
		linkHeartbeatMessage,
	}
	for i, tt := range tests {
		b := &bytes.Buffer{}
		enc := &messageEncoder{w: b}
		if err := enc.encode(tt); err != nil {
			panic(fmt.Errorf("#%d: unexpected encode message error: %v", i, err))
			continue
		}
		dec := &messageDecoder{r: b}
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
