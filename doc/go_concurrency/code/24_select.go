package main

import (
	"fmt"
	"time"
)

func send(msg string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
			if i == 5 {
				fmt.Println("Sleeping 2 seconds...")
				time.Sleep(2 * time.Second)
			}
		}
	}()
	return ch
}

func main() {
	ch := send("Hello")
	for {
		select {
		case v := <-ch:
			fmt.Println("Received:", v)
		case <-time.After(time.Second):
			fmt.Println("Done!")
			return
		}
	}
}

/*
Received: Hello 0
Received: Hello 1
Received: Hello 2
Received: Hello 3
Received: Hello 4
Received: Hello 5
Sleeping 2 seconds...
Done!
*/
