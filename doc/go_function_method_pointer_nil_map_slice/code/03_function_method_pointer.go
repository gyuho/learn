package main

import "fmt"

type myType struct{ val int }

func noChange(t myType) {
	t.val = 100
	// this only changes the copied
}

func (t myType) noChange() {
	t.val = 200
	// this only changes the copied
}

func change(t *myType) {
	t.val = 300
	// this updates the original
	// that the pointer t points to
}

func (t *myType) change() {
	t.val = 400
	// this updates the original
	// that the pointer t points to
}

func main() {
	one := myType{val: 1}
	noChange(one)
	fmt.Printf("%+v\n", one) // {val:1}
	one.noChange()
	fmt.Printf("%+v\n", one) // {val:1}

	change(&one)
	fmt.Printf("%+v\n", one) // {val:300}
	(&one).change()
	fmt.Printf("%+v\n", one) // {val:400}
}
