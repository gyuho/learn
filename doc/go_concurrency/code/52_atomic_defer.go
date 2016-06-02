package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var proposeCounter int32

func main() {
	atomic.AddInt32(&proposeCounter, 1)
	defer atomic.AddInt32(&proposeCounter, -1)

	fmt.Println(atomic.LoadInt32(&proposeCounter)) // 1
	time.Sleep(time.Second)
}
