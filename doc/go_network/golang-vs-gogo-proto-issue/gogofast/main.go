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
		// unknown field "create_request" in main.WatchRequest
	}
	fmt.Printf("success with %+v\n", req.RequestUnion)
}
