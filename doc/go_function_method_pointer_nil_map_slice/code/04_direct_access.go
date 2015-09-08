package main

import "fmt"

type myType struct{ val int }

func main() {
	one := myType{val: 1}

	// direct access
	one.val = 100
	fmt.Printf("%+v\n", one) // {val:100}

	// copy the value
	value := one
	value.val = 200
	fmt.Printf("%+v\n", one) // {val:100}

	// access to the original data through pointer
	pointer := &one
	pointer.val = 300
	fmt.Printf("%+v\n", one) // {val:300}
}
