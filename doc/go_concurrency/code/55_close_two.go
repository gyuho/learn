package main

import (
	"fmt"
	"time"
)

func main() {
	closed, donec := make(chan struct{}), make(chan struct{})
	go func() {
		select {
		case <-time.After(3 * time.Second):
			fmt.Println("close(closed) took too long")
		case <-closed:
		}
		close(donec)
	}()

	close(closed)

	select {
	case <-time.After(3 * time.Second):
		fmt.Println("close(donec) took too long")
	case <-donec:
	}

	fmt.Println("DONE!")
	// DONE!
}
