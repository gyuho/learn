package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/jsonpb"
)

func main() {
	const input = `{"create_request": {"key": "Zm9v"}}`
	var req WatchRequest
	if err := jsonpb.UnmarshalString(input, &req); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("success with %+v\n", req.RequestUnion)
	// success with &{CreateRequest:key:"foo" }
}
