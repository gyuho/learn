package main

import (
	"fmt"
	"log"
	"sync/atomic"
	"time"
)

func main() {
	var refCounter int32 = 0
	fmt.Println(atomic.LoadInt32(&refCounter))
	fmt.Println(atomic.AddInt32(&refCounter, 1))
	fmt.Println(atomic.LoadInt32(&refCounter))
	fmt.Println(refCounter)

	go func() {
		time.Sleep(10 * time.Second)
		atomic.AddInt32(&refCounter, -1)
	}()

	for atomic.LoadInt32(&refCounter) != 0 {
		log.Println("Sleeping 20 seconds")
		time.Sleep(20 * time.Second)
		fmt.Println(refCounter)
	}
	atomic.AddInt32(&refCounter, 1)
	atomic.AddInt32(&refCounter, -1)
}
