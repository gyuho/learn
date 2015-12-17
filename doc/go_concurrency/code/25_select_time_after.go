package main

import (
	"log"
	"time"
)

func main() {
	chs := make([]chan struct{}, 100)

	// init
	for i := range chs {
		chs[i] = make(chan struct{}, 1)
	}

	// close
	for _, ch := range chs {
		close(ch)
	}

	// receive
	for _, ch := range chs {
		select {
		case <-ch:
			// https://golang.org/ref/spec#Close
			// After calling close, and after any previously sent values
			// have been received, receive operations will return the zero
			// value for the channel's type without blocking.
			log.Println("Succeed")

			// http://golang.org/ref/spec#Select_statements
			// time.After _is_ evaluated each time.
			// https://groups.google.com/d/msg/golang-nuts/1tjcV80ccq8/hcoP9uMNiUcJ
		case <-time.After(time.Millisecond):
			log.Fatalf("Receive Delayed!")
		}
	}
}

/*
...
2015/06/27 14:34:48 Succeed
2015/06/27 14:34:48 Succeed
2015/06/27 14:34:48 Succeed
2015/06/27 14:34:48 Succeed
*/
