package main

import (
	"fmt"
	"log"

	"github.com/coreos/etcd/etcdserver/etcdserverpb"
	"github.com/golang/protobuf/jsonpb"
)

func main() {
	const input = `{"create_request": {"key": "Zm9v"}}`
	var protoReq etcdserverpb.WatchRequest
	if err := jsonpb.UnmarshalString(input, &protoReq); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("success with %+v\n", protoReq.RequestUnion)
}
