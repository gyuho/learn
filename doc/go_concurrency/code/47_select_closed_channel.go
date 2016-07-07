package main

import (
	"fmt"
	"time"
)

func main() {
	stream1 := make(chan string, 5000)
	stream2 := make(chan string, 5000)

	go func() {
		for i := 0; i < 5; i++ {
			stream1 <- fmt.Sprintf("%d", i)
			time.Sleep(time.Second)
		}
		// close(stream1)
	}()
	go func() {
		for i := 0; i < 5; i++ {
			stream2 <- fmt.Sprintf("%d", i)
			time.Sleep(time.Microsecond)
		}
		// close(stream2)
	}()

escape1:
	for {
		select {
		case s := <-stream1:
			fmt.Println("stream1:", s)

		case s := <-stream2:
			fmt.Println("stream2:", s)

		case <-time.After(500 * time.Millisecond):
			fmt.Println("escape1")
			break escape1
		}
	}

	fmt.Println()
	fmt.Println("first done")

escape2:
	for {
		select {
		case s := <-stream1:
			fmt.Println("stream1:", s)

		case s := <-stream2:
			fmt.Println("stream2:", s)

		case <-time.After(5 * time.Second):
			fmt.Println("escape2")
			break escape2
		}
	}
	/*
		stream1: 0
		stream2: 0
		stream2: 1
		stream2: 2
		stream2: 3
		stream2: 4
		escape1

		first done
		stream1: 1
		stream1: 2
		stream1: 3
		stream1: 4
		escape2
	*/
}
