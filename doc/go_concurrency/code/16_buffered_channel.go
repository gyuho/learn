package main

import "fmt"

func main() {
	ch := make(chan int, 100)
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println(ch, len(ch), cap(ch)) // 0xc420072380 3 100

	fmt.Println(<-ch) // 1

	fmt.Println(ch, len(ch), cap(ch)) // 0xc420072380 2 100

	fmt.Println(<-ch) // 2

	fmt.Println(ch, len(ch), cap(ch)) // 0xc420072380 1 100

	fmt.Println(<-ch) // 3

	// fmt.Println(<-ch)
	// fatal error: all goroutines are asleep - deadlock!

	ch <- 5
	ch <- 10
	fmt.Println(ch, len(ch), cap(ch)) // 0xc42005e380 2 100
}
