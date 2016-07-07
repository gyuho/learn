package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	cond := &sync.Cond{L: &sync.Mutex{}}
	sig := make(chan struct{})

	for i := 0; i < 3; i++ {
		go func(i int) {
			cond.L.Lock()
			sig <- struct{}{}
			fmt.Println("Wait begin:", i)
			cond.Wait()
			fmt.Println("Wait end:", i)
			cond.L.Unlock()
		}(i)
	}
	for range []int{0, 1, 2} {
		<-sig
	}

	// for i := 0; i < 3; i++ {
	// 	cond.L.Lock()
	// 	fmt.Println("Signal")
	// 	cond.Signal()
	// 	cond.L.Unlock()
	// }

	cond.L.Lock()
	fmt.Println("Broadcast")
	cond.Broadcast()
	cond.L.Unlock()

	fmt.Println("Sleep")
	time.Sleep(time.Second)
}

/*
Wait begin: 2
Wait begin: 1
Wait begin: 0
Broadcast
Sleep
Wait end: 2
Wait end: 1
Wait end: 0
*/
