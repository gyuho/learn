package main

import "fmt"

func main() {
	defer println("Defer 1")
	defer println("Defer 2")
	defer println("Defer 3")

	defer func() {
		fmt.Println("Recover:", recover())
	}()
	panic("Panic!!!")

	/*
		Recover: Panic!!!
		Defer 3
		Defer 2
		Defer 1
	*/

	// recover stops the panic
	// recover returns the value from panic
	// panic function is to cause a run time error
	// for "cannot happen" situations
	// And stops the program to begin panicking
	// So even if it's recovered
	// the next lines after panic won't be run.
	for {
		fmt.Println("This does not print! Anything below not being run!")
	}
}
