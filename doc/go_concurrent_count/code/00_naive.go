package main

import (
	"fmt"
	"sync"
)

// Counter is an interface for counting.
// It contains counting data as long as a type
// implements all the methods in the interface.
type Counter interface {
	// Get returns the current count.
	Get() float64

	// Add adds the delta value to the counter.
	Add(delta float64)
}

// NaiveCounter counts in a naive way.
// Do not use this with concurrency.
// It will cause race conditions.
type NaiveCounter float64

func (c *NaiveCounter) Get() float64 {

	// return (*c).(float64)
	// (X) (*c).(float64) (non-interface type NaiveCounter on left)

	return float64(*c)
}

func (c *NaiveCounter) Add(delta float64) {
	*c += NaiveCounter(delta)
}

func main() {
	counter := new(NaiveCounter)
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			counter.Add(1.347)
			counter.Get()
			counter.Add(-5.5)
			counter.Get()
			counter.Add(0.340)
			counter.Get()
		}()
	}
	wg.Wait()

	fmt.Println(counter.Get())
	// -38.12999999999999
}

/*
go run -race 00_naive.go

==================
WARNING: DATA RACE
Read by goroutine 7:
  main.main.func1()
      /home/ubuntu/go/src/github.com/gyuho/learn/doc/go_concurrent_count/code/00_naive.go:43 +0x70

Previous write by goroutine 6:
  main.main.func1()
      /home/ubuntu/go/src/github.com/gyuho/learn/doc/go_concurrent_count/code/00_naive.go:43 +0x88

Goroutine 7 (running) created at:
  main.main()
      /home/ubuntu/go/src/github.com/gyuho/learn/doc/go_concurrent_count/code/00_naive.go:49 +0xc5

Goroutine 6 (finished) created at:
  main.main()
      /home/ubuntu/go/src/github.com/gyuho/learn/doc/go_concurrent_count/code/00_naive.go:49 +0xc5
==================
-38.12999999999999
Found 1 data race(s)
exit status 66

*/
