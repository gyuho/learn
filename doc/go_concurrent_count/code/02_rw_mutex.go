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

// RWMutexCounter implements Counter with sync.RWMutex.
type RWMutexCounter struct {
	sync.RWMutex
	value float64
}

func (c *RWMutexCounter) Get() float64 {
	c.RLock()
	defer c.RUnlock()
	return c.value
}

func (c *RWMutexCounter) Add(delta float64) {
	c.Lock()
	defer c.Unlock()
	c.value += delta
}

func main() {
	counter := new(RWMutexCounter)
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
