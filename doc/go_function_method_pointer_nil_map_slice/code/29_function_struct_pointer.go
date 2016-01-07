package main

import "fmt"

type aaa struct {
	Val int
}

func update(a *aaa) {
	a.Val = 100
}

func main() {
	a := &aaa{}
	update(a)
	fmt.Println(a) // &{100}
}
