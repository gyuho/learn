package main

import (
	"fmt"
	"time"
)

func main() {
	stream1 := make(chan string, 5000)
	stream2 := make(chan string, 5000)

	go func() {
		for i := 0; i < 3; i++ {
			stream1 <- fmt.Sprintf("%d", i)
			time.Sleep(time.Second)
		}
		close(stream1)
	}()
	go func() {
		for i := 0; i < 5; i++ {
			stream2 <- fmt.Sprintf("%d", i)
			time.Sleep(time.Microsecond)
		}
		close(stream2)
	}()

escape:
	for {
		select {
		case s, ok := <-stream1:
			if !ok {
				fmt.Println("stream1 closed")
				break escape // when the channel is closed
				// without escape, infinite loop
			}
			fmt.Println("stream1:", s)

		case s, ok := <-stream2:
			if !ok {
				fmt.Println("stream2 closed")
				break escape // when the channel is closed
			}
			fmt.Println("stream2:", s)

		case <-time.After(time.Second):
			// drain channel until it takes longer than 1 second
			fmt.Println("escaping")
			break escape
		}
	}
	/*
		stream1 couldn't finish!

		stream1: 0
		stream2: 0
		stream2: 1
		stream2: 2
		stream2: 3
		stream2: 4
		stream2 closed
	*/
}
