package main

import (
	"fmt"
	"reflect"
)

func main() {
	var slice = []interface{}{"A", 1, struct{}{}}
	fmt.Printf("slice %+v\n", slice)
	fmt.Printf("slice %+v\n", reflect.TypeOf(slice))

	fmt.Printf("slice %+v\n", reflect.ValueOf(slice).Kind())
	fmt.Printf("slice %+v\n", reflect.ValueOf(slice).String())
}

/*
slice [A 1 {}]
slice []interface {}
slice slice
slice <[]interface {} Value>
*/
