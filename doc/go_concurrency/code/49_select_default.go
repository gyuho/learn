package main

import (
	"fmt"
	"time"
)

func main() {
	select {
	case <-time.After(time.Nanosecond):
		fmt.Println("received from time.Nanosecond")
	default:
		fmt.Println("default")
	}
	// default

	done := make(chan struct{})
	close(done)
	select {
	case <-time.After(time.Nanosecond):
		fmt.Println("received from time.Nanosecond")
	case <-done:
		fmt.Println("received from done")
	default:
		fmt.Println("default")
	}
	// received from done
}
