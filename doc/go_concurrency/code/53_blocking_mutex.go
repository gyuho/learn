package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex

	fmt.Println("Holding Lock!")
	mu.Lock()
	time.Sleep(time.Second)
	mu.Unlock()
	fmt.Println("Released Lock!")

	donec := make(chan struct{})
	go func() {
		fmt.Println("goroutine is trying to holding the same Lock!")
		mu.Lock()
		fmt.Println("goroutine got the Lock!")
		mu.Unlock()
		fmt.Println("goroutine just released the Lock!")
		close(donec)
	}()

	<-donec
	fmt.Println("DONE")
}

/*
Holding Lock!
Released Lock!
goroutine is trying to holding the same Lock!
goroutine got the Lock!
goroutine just released the Lock!
DONE
*/
