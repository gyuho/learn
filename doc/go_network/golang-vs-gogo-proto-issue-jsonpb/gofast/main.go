package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/jsonpb"
)

func main() {
	const input = `{"key": "Zm9v", "type": "READ"}`
	var req WatchCreateRequest
	if err := jsonpb.UnmarshalString(input, &req); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("success with %+v\n", req)
	// success with {Key:[102 111 111] Type:READ}
}
