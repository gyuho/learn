package main

import "fmt"

type aaa struct {
	Val int
}

func update(a *aaa) {
	a.Val = 100
}

func updateNil(a *aaa) {
	a = nil // only updates pointer value
}

func updateNilVal(a *aaa) {
	*a = aaa{}
}

func main() {
	a := &aaa{}
	update(a)
	fmt.Println(a) // &{100}

	updateNil(a)
	fmt.Println(a) // &{100}

	a = &aaa{Val: 200}
	updateNilVal(a)
	fmt.Println(a) // &{0}
}
