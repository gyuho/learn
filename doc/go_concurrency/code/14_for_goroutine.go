package main

import (
	"math/rand"
)

func main() {
	ch := make(chan int)

	for {
		go func() {
			ch <- rand.Intn(10)
		}()
	}

	<-ch

	// process took too long
}
