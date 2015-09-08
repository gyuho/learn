package main

import "fmt"

type myType struct{ val int }

func a() *myType {
	one := myType{val: 1}
	return &one
}

func main() {
	one := a()
	fmt.Printf("%+v\n", one) // &{val:1}
}
