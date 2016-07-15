package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/jsonpb"
)

func main() {
	const input = `{"key": "Zm9v", "type": "READ"}`
	// works with
	// const input = `{"key": "Zm9v", "type": 0}`

	var req WatchCreateRequest
	if err := jsonpb.UnmarshalString(input, &req); err != nil {
		log.Fatal(err)
		// unknown value "READ" for enum main.Type
	}
	fmt.Printf("success with %+v\n", req)
}
