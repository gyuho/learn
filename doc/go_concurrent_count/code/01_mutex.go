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

// MutexCounter implements Counter with sync.Mutex.
type MutexCounter struct {
	sync.Mutex
	value float64
}

func (c *MutexCounter) Get() float64 {
	c.Lock()
	defer c.Unlock()
	return c.value
}

func (c *MutexCounter) Add(delta float64) {
	c.Lock()
	defer c.Unlock()
	c.value += delta
}

func main() {
	counter := new(MutexCounter)
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
	// 962.0000000000002
}
