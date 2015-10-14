package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	q := NewQueue(10, time.Second/10)
	tick := time.NewTicker(time.Second / 8)
	now := time.Now()
	for t := range tick.C {
		fmt.Println("took:", time.Since(now))
		now = time.Now()

		isExceeded := q.Push(t)
		if isExceeded {
			fmt.Println(t, "has exceeded the rate limit", q.rate)
			tick.Stop()
			break
		}
	}
	/*
	   took: 125.18325ms
	   took: 124.804029ms
	   took: 124.965582ms
	   took: 124.955137ms
	   took: 124.960788ms
	   took: 124.965522ms
	   took: 124.982369ms
	   took: 124.959646ms
	   took: 124.944575ms
	   took: 124.927333ms
	   2015-09-03 14:45:27.868522371 -0700 PDT has exceeded the rate limit 100ms
	*/
}

// timeSlice stores a slice of time.Time
// in a thread-safe way.
type timeSlice struct {
	times []time.Time

	// RWMutex is more expensive
	// https://blogs.oracle.com/roch/entry/beware_of_the_performance_of
	// sync.RWMutex
	//
	// to synchronize access to shared state across multiple goroutines.
	//
	sync.Mutex
}

func newTimeSlice() *timeSlice {
	tslice := timeSlice{}
	sl := make([]time.Time, 0)
	tslice.times = sl
	return &tslice
}

func (t *timeSlice) push(ts time.Time) {
	t.Lock()
	t.times = append(t.times, ts)
	t.Unlock()
}

func (t *timeSlice) length() int {
	t.Lock()
	d := len(t.times)
	t.Unlock()
	return d
}

func (t *timeSlice) pop() {
	if t.length() != 0 {
		t.Lock()
		t.times = t.times[1:len(t.times):len(t.times)]
		t.Unlock()
	}
}

func (t *timeSlice) first() (time.Time, bool) {
	if t.length() == 0 {
		return time.Time{}, false
	}
	t.Lock()
	v := t.times[0]
	t.Unlock()
	return v, true
}

// Queue contains the slice of timestamps
// and other rate limiter configurations.
type Queue struct {
	slice *timeSlice

	// burstSize is like a buffer.
	// If burstSize is 5, it allows rate exceeding
	// for the fist 5 elements.
	burstSize int
	rate      time.Duration
}

// NewQueue returns a new Queue.
func NewQueue(burstSize int, rate time.Duration) *Queue {
	tslice := newTimeSlice()
	q := Queue{}
	q.slice = tslice
	q.burstSize = burstSize
	q.rate = rate
	return &q
}

// Push appends the timestamp to the Queue.
// It return true if rate has exceeded.
// We need a pointer of Queue, where it defines
// timeSlice with pointer as well. To append to slice
// and update struct members, we need pointer types.
func (q *Queue) Push(ts time.Time) bool {
	if q.slice.length() == q.burstSize {
		q.slice.pop()
	}
	q.slice.push(ts)
	if q.slice.length() < q.burstSize {
		return false
	}
	ft, ok := q.slice.first()
	if !ok {
		return false
	}
	diff := ft.Sub(ts)
	return q.rate > diff
}

func (q *Queue) String() string {
	return fmt.Sprintf("times: %+v / burstSize: %d / rate: %v", q.slice.times, q.burstSize, q.rate)
}
