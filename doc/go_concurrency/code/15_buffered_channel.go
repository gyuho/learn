package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(ch, len(ch), cap(ch))
	// 0x1052d080 2 2

	<-ch // 1 is retrieved and discarded
	fmt.Println(ch, len(ch), cap(ch))
	// 0x1052d080 1 2

	fmt.Println(<-ch) // 2
	fmt.Println(ch, len(ch), cap(ch))
	// 0x1052d080 0 2

	// fmt.Println(<-ch)
	// fatal error: all goroutines are asleep - deadlock!

	ch <- 5
	ch <- 10
	fmt.Println(ch, len(ch), cap(ch))
	// 0x1052d080 2 2
}
