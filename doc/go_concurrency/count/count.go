package count

import (
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

// NaiveCounter counts in a naive way. Do not use this with concurrency.
// It will cause race conditions. This is not thread-safe.
type NaiveCounter float64

func (c *NaiveCounter) Get() float64 {

	// return (*c).(float64)
	// (X) (*c).(float64) (non-interface type NaiveCounter on left)

	return float64(*c)
}

func (c *NaiveCounter) Add(delta float64) {
	*c += NaiveCounter(delta)
}

// MutexCounter implements Counter with sync.Mutex.
type MutexCounter struct {
	mu    sync.Mutex // guards the following
	value float64
}

func (c *MutexCounter) Get() float64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func (c *MutexCounter) Add(delta float64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value += delta
}

// RWMutexCounter implements Counter with sync.RWMutex.
type RWMutexCounter struct {
	mu    sync.RWMutex // guards the following sync.
	value float64
}

func (c *RWMutexCounter) Get() float64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.value
}

func (c *RWMutexCounter) Add(delta float64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value += delta
}

// AtomicIntCounter implements Counter with atomic package.
// Go has only int64 atomic variable.
// This truncates float value into integer.
type AtomicIntCounter int64

func (c *AtomicIntCounter) Get() float64 {
	return float64(atomic.LoadInt64((*int64)(c)))
}

// Add ignores the non-integer part of delta.
func (c *AtomicIntCounter) Add(delta float64) {
	atomic.AddInt64((*int64)(c), int64(delta))
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

// ChannelCounter counts through channels.
type ChannelCounter struct {
	valueChan chan float64
	deltaChan chan float64
	done      chan struct{}
}

func NewChannelCounter(buf int) *ChannelCounter {
	c := &ChannelCounter{
		make(chan float64),
		make(chan float64, buf), // only buffer the deltaChan
		make(chan struct{}),
	}
	go c.Run()
	return c
}

func (c *ChannelCounter) Run() {

	var value float64

	for {
		// "select" statement chooses which of a set of
		// possible send or receive operations will proceed.
		select {

		case delta := <-c.deltaChan:
			value += delta

		case <-c.done:
			return

		case c.valueChan <- value:
			// Do nothing.

			// If there is no default case, the "select" statement
			// blocks until at least one of the communications can proceed.
		}
	}
}

func (c *ChannelCounter) Get() float64 {
	return <-c.valueChan
}

func (c *ChannelCounter) Add(delta float64) {
	c.deltaChan <- delta
}

func (c *ChannelCounter) Done() {
	c.done <- struct{}{}
}

func (c *ChannelCounter) Close() {
	close(c.deltaChan)
}
