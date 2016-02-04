package main

import "fmt"

func channelClose() <-chan int {
	ch := make(chan int)
	close(ch)
	return ch
}

func channelCloseArg(ch chan int) {
	close(ch)
}

func main() {
	cc := make(chan int, 1)
	cc <- 1
	v, open := <-cc
	fmt.Println(v, open) // 1 true

	close(cc)
	v, open = <-cc
	fmt.Println(v, open) // 0 false
	v, open = <-cc
	fmt.Println(v, open) // 0 false

	fmt.Println()

	rc := channelClose()
	v, open = <-rc
	fmt.Println(v, open) // 0 false
	v, open = <-rc
	fmt.Println(v, open) // 0 false
	v, open = <-rc

	fmt.Println()

	ch := make(chan int)
	channelCloseArg(ch)
	v, open = <-ch
	fmt.Println(v, open) // 0 false
	v, open = <-ch
	fmt.Println(v, open) // 0 false
	v, open = <-ch
}
