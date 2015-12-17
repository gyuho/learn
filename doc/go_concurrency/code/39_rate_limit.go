package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	func() {
		q := NewQueue(10, time.Second/3)
		for range []int{0, 1, 2, 3, 4, 5, 6, 7, 8} {
			if q.Push(time.Now()) {
				log.Fatalf("should not have exceeded the rate limit: %+v", q)
			}
		}
		for range []int{0, 1, 2, 3, 4, 5, 6, 7, 8} {
			if !q.Push(time.Now()) {
				log.Fatalf("should have exceeded the rate limit: %+v", q)
			}
		}
		if q.slice.length() != 10 {
			log.Fatalf("Queue should only have 10 timestamps: %+v", q)
		}
	}()

	func() {
		q := NewQueue(10, time.Second/10)
		tick := time.NewTicker(time.Second / 8)
		done := make(chan struct{})
		go func() {
			now := time.Now()
			for tk := range tick.C {
				log.Println("took:", time.Since(now))
				now = time.Now()
				isExceeded := q.Push(tk)
				if isExceeded {
					log.Println(tk, "has exceeded the rate limit", q.rate)
					tick.Stop()
					done <- struct{}{}
					break
				}
			}
		}()
		select {
		case <-time.After(3 * time.Second):
			log.Fatalln("time out!")
		case <-done:
			log.Println("success")
		}
	}()
}

// timeSlice stores a slice of time.Time
// in a thread-safe way.
type timeSlice struct {
	// RWMutex is more expensive
	// https://blogs.oracle.com/roch/entry/beware_of_the_performance_of
	// sync.RWMutex
	//
	// to synchronize access to shared state across multiple goroutines.
	//
	mu sync.Mutex

	times []time.Time
}

func newTimeSlice() *timeSlice {
	tslice := timeSlice{}
	sl := make([]time.Time, 0)
	tslice.times = sl
	return &tslice
}

func (t *timeSlice) push(ts time.Time) {
	t.mu.Lock()
	t.times = append(t.times, ts)
	t.mu.Unlock()
}

func (t *timeSlice) length() int {
	t.mu.Lock()
	d := len(t.times)
	t.mu.Unlock()
	return d
}

func (t *timeSlice) pop() {
	if t.length() != 0 {
		t.mu.Lock()
		t.times = t.times[1:len(t.times):len(t.times)]
		t.mu.Unlock()
	}
}

func (t *timeSlice) first() (time.Time, bool) {
	if t.length() == 0 {
		return time.Time{}, false
	}
	t.mu.Lock()
	v := t.times[0]
	t.mu.Unlock()
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
