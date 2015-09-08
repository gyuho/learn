package main

import (
	"fmt"
	"reflect"
)

// type for numeric value
type Int int

// type for struct
type data struct {
	Value string
}

// type for interface
type Interface interface {
	Len() int
}

func (d data) Len() int {
	return len(d.Value)
}

func main() {
	v0 := Int(0)
	fmt.Println(reflect.TypeOf(v0))
	// main.Int

	v1 := data{Value: "A"}
	fmt.Println(reflect.TypeOf(v1))
	// main.data

	v2 := interface{}(v1).(Interface)
	fmt.Println(reflect.TypeOf(v2))
	// main.data
}
