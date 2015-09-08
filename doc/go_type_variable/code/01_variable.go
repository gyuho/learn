package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var val0 string
	fmt.Println(reflect.TypeOf(val0), unsafe.Sizeof(val0), val0)
	// string 16

	var val1 = "A"
	fmt.Println(reflect.TypeOf(val1), unsafe.Sizeof(val1), val1)
	// string 16 A

	val2 := "B"
	fmt.Println(reflect.TypeOf(val2), unsafe.Sizeof(val2), val2)
	// string 16 B

	var data1 = struct{}{}
	fmt.Println(reflect.TypeOf(data1), unsafe.Sizeof(data1), data1)
	// struct {} 0 {}

	var data2 struct{}
	fmt.Println(reflect.TypeOf(data2), unsafe.Sizeof(data2), data2)
	// struct {} 0 {}
}
