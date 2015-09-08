package main

import "fmt"

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	panic("Panic!")

	fmt.Println("Hello, World!")
	// NOT printed
}

/*
Panic!
*/
