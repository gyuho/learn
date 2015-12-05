package ratelimit

import (
	"testing"
	"time"
)

func TestPush(t *testing.T) {
	q := NewQueue(10, time.Second/3)
	for range []int{0, 1, 2, 3, 4, 5, 6, 7, 8} {
		if q.Push(time.Now()) {
			t.Errorf("should not have exceeded the rate limit: %+v", q)
		}
	}
	for range []int{0, 1, 2, 3, 4, 5, 6, 7, 8} {
		if !q.Push(time.Now()) {
			t.Errorf("should have exceeded the rate limit: %+v", q)
		}
	}
	if q.slice.length() != 10 {
		t.Errorf("Queue should only have 10 timestamps: %+v", q)
	}
}

func TestRate(t *testing.T) {
	q := NewQueue(10, time.Second/10)
	tick := time.NewTicker(time.Second / 8)
	done := make(chan struct{})
	go func() {
		now := time.Now()
		for tk := range tick.C {
			t.Log("took:", time.Since(now))
			now = time.Now()
			isExceeded := q.Push(tk)
			if isExceeded {
				t.Log(tk, "has exceeded the rate limit", q.rate)
				tick.Stop()
				done <- struct{}{}
				break
			}
		}
	}()
	select {
	case <-time.After(3 * time.Second):
		t.Error("time out!")
	case <-done:
		t.Log("success")
	}
}
