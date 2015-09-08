package main

import "fmt"

func main() {
	goto Here

	for {
		fmt.Println("Infinite Looping; Not Printing")
	}

Escape:
	fmt.Println("After goto Escape.")
	return

Here:
	fmt.Println("After goto Here.")
	for i := 0; i < 2; i++ {
		fmt.Println("Hello")
		goto Escape
	}
}

/*
After goto Here.
Hello
After goto Escape.
*/
