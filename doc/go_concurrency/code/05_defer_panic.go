package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		defer fmt.Println("Hello, playground")
		panic(1)
	}()

	time.Sleep(time.Second)
}

/*
Hello, playground
panic: 1
*/
