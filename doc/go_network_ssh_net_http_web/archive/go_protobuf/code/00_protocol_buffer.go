package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"

	"github.com/gyuho/learn/doc/go_protobuf/datapb"
)

func main() {
	d := &datapb.SampleData{
		Label: proto.String("hello"),
		Type:  proto.Int32(17),
		Optionalgroup: &datapb.SampleData_OptionalGroup{
			RequiredField: proto.String("good bye"),
		},
	}
	fmt.Println("d.GetLabel():", d.GetLabel())
	fmt.Println()
	// d.GetLabel(): hello

	data, err := proto.Marshal(d)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	fmt.Printf("data: %+v\n", data)
	fmt.Println()
	// data: [10 5 104 101 108 108 111 ...

	newD := &datapb.SampleData{}
	if err := proto.Unmarshal(data, newD); err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	fmt.Println("newD.GetLabel():", newD.GetLabel())
	fmt.Println()
	// newD.GetLabel(): hello

	fmt.Printf("newD: %+v\n", newD)
	// newD: label:"hello" type:17 OptionalGroup{RequiredField:"good bye" }
}
