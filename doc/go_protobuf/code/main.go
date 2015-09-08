package main

import (
	"log"

	"github.com/golang/protobuf/proto"

	"github.com/gyuho/learn/doc/go_protobuf/example"
)

func main() {

	test := &example.Test{
		Label: proto.String("hello"),
		Type:  proto.Int32(17),
		Optionalgroup: &example.Test_OptionalGroup{
			RequiredField: proto.String("good bye"),
		},
	}
	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	newTest := &example.Test{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	// Now test and newTest contain the same data.
	if test.GetLabel() != newTest.GetLabel() {
		log.Fatalf("data mismatch %q != %q", test.GetLabel(), newTest.GetLabel())
	}

	log.Printf("Unmarshalled to: %+v", newTest)
	// 2015/07/09 00:58:49 Unmarshalled to: label:"hello" type:17 OptionalGroup{RequiredField:"good bye" }
}
