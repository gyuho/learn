package main

import "fmt"

func myFunc(num int) {
	fmt.Println(num + 1)
}

type Int int

func (num Int) myMethod() {
	fmt.Println(num + 2)
}

func main() {
	myFunc(1)         // 2
	Int(1).myMethod() // 3
}
