package main

import (
	"fmt"
	"sync"
	"time"
)

/*

Ian Lance Taylor (https://groups.google.com/d/msg/golang-nuts/7Xi2APcqpM0/UzHJnabiDQAJ):

You are looking at this incorrectly in some way that I don't
understand.  A sync.Mutex is a value with two methods: Lock and
Unlock.  Lock acquires a lock on the mutex.  Unlock releases it.  Only
one goroutine can acquire the lock on the mutex at a time.

That's all there is.  A mutex doesn't have a scope.  It can be a field
of a struct but it doesn't have to be.  A mutex doesn't protect
anything in particular by itself.  You have to write your code to call
Lock, do the protected operations, and then call Unlock.

Your example code looks fine.
*/

func main() {
	var hits struct {
		sync.Mutex
		n int
	}
	hits.Lock()
	hits.n++
	hits.Unlock()
	fmt.Println(hits)
	// {{0 0} 1}

	m := map[string]time.Time{}

	// without this:
	// Found 1 data race(s)
	var mutex sync.Mutex

	done := make(chan struct{})
	for range []int{0, 1} {
		go func() {
			mutex.Lock()
			m[time.Now().String()] = time.Now()
			mutex.Unlock()
			done <- struct{}{}
		}()
	}
	cn := 0
	for range done {
		cn++
		if cn == 2 {
			close(done)
		}
	}
	fmt.Println(m)
	/*
	   map[2015-11-05 20:42:36.516629792 -0800 PST:2015-11-05 20:42:36.516678634 -0800 PST 2015-11-05 20:42:36.516685141 -0800 PST:2015-11-05 20:42:36.516686379 -0800 PST]
	*/
}
