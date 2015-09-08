package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

func main() {
	num := 100000000

	sendOneTo := func(c chan int) {
		for i := 0; i < num; i++ {
			c <- 1
		}
	}

	connect := func(cin, cout chan int) {
		for {
			x := <-cin
			cout <- x
		}
	}

	round := func(ch1, ch2, ch3, ch4 chan int) {
		go connect(ch1, ch2)
		go connect(ch2, ch3)
		go connect(ch3, ch4)
		go sendOneTo(ch1)

		for i := 0; i < num; i++ {
			_ = <-ch4
		}
	}

	startBfCh := time.Now()
	bfCh1 := make(chan int, num)
	bfCh2 := make(chan int, num)
	bfCh3 := make(chan int, num)
	bfCh4 := make(chan int, num)
	round(bfCh1, bfCh2, bfCh3, bfCh4)
	fmt.Println("[Asynchronous, Non-Blocking] Buffered   took", time.Since(startBfCh))

	startUnCh := time.Now()
	unCh1 := make(chan int)
	unCh2 := make(chan int)
	unCh3 := make(chan int)
	unCh4 := make(chan int)
	round(unCh1, unCh2, unCh3, unCh4)
	fmt.Println("[Synchronous,  Blocking]      UnBuffered took", time.Since(startUnCh))
}

/*
[Asynchronous, Non-Blocking] Buffered   took 32.96282781s     (30 seconds)
[Synchronous,  Blocking]     UnBuffered took 3m17.140920286s  (3 minutes)
*/

func init() {
	maxCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.Println("Concurrent execution with", maxCPU, "CPUs.")
}
