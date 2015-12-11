package main

import "fmt"

func main() {
	// make sure to assign step by step!!!
	m1 := make(map[string]map[string]struct{})
	m1["A"] = make(map[string]struct{})
	m1["A"]["B"] = struct{}{}
	fmt.Println(m1) // map[A:map[B:{}]]}]

	m1["X"]["C"] = struct{}{}
	// panic: assignment to entry in nil map
}
