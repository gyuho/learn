package main

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
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

// AtomicCounter implements Counter with atomic package.
// Go has only int64 atomic variable.
// This uses math.Float64frombits package for the floating
// point number corresponding the IEEE 754 binary representation
type AtomicCounter uint64

func (c *AtomicCounter) Get() float64 {
	return math.Float64frombits(atomic.LoadUint64((*uint64)(c)))
}

// Add ignores the non-integer part of delta.
func (c *AtomicCounter) Add(delta float64) {
	for {
		oldBits := atomic.LoadUint64((*uint64)(c))
		newBits := math.Float64bits(math.Float64frombits(oldBits) + delta)
		if atomic.CompareAndSwapUint64((*uint64)(c), oldBits, newBits) {
			return
		}
	}
}

func main() {
	counter := new(AtomicCounter)
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
