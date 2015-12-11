package main

import "fmt"

type List struct{ root Element }
type Element struct{ val int }

func change(l List) { l.root.val = 100 }

func main() {
	l := List{}
	l.root = Element{val: 1}
	l.root.val = 2
	fmt.Printf("%+v\n", l) // {root:{val:2}}
	// this updates because we are not passing the copy
	// it's not in the function or method

	change(l)
	fmt.Printf("%+v", l) // {root:{val:2}}
	// passing the non-pointer
	// and only the copied data is passed
	// so it can't update the original value
}
