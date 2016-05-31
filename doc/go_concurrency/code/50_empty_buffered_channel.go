package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 1)
	donec := make(chan struct{})
	go func() {
		time.Sleep(time.Second)
		ch <- "hello"
		close(donec)
	}()

	fmt.Println(<-ch) // hello
	<-donec
}
