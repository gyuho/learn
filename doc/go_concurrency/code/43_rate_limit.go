package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	st := &stresser{
		qps: 100000,
		N:   100,
	}

	go st.Start()
	time.Sleep(time.Second)
	st.StopAndWait()
	fmt.Println("st.success", st.success)

	println()
	go st.Start()
	time.Sleep(time.Second)
	st.StopAndWait()
	fmt.Println("st.success", st.success)

	println()
	go st.Start()
	time.Sleep(time.Second)
	st.StopAndWait()
	fmt.Println("st.success", st.success)
}

/*
s.cancel() 1
s.cancel() 2
Start finished with context canceled
wg.Wait() 1
wg.Wait() 2
st.success 100001

s.cancel() 1
s.cancel() 2
wg.Wait() 1
Start finished with context canceled
wg.Wait() 2
st.success 200002

s.cancel() 1
s.cancel() 2
wg.Wait() 1
Start finished with context canceled
wg.Wait() 2
st.success 300003
*/

type stresser struct {
	qps int
	N   int

	mu          sync.Mutex
	wg          *sync.WaitGroup
	rateLimiter *rate.Limiter
	cancel      func()

	canceled bool
	success  int
}

func (s *stresser) Start() {
	ctx, cancel := context.WithCancel(context.Background())

	wg := &sync.WaitGroup{}
	wg.Add(s.N)

	s.mu.Lock()
	s.wg = wg
	s.rateLimiter = rate.NewLimiter(rate.Every(time.Second), s.qps)
	s.cancel = cancel
	s.canceled = false
	s.mu.Unlock()

	for i := 0; i < s.N; i++ {
		go s.run(ctx)
	}

	<-ctx.Done()
	fmt.Println("Start finished with", ctx.Err())
}

func (s *stresser) run(ctx context.Context) {
	defer s.wg.Done()

	for {
		if err := s.rateLimiter.Wait(ctx); err == context.Canceled {
			return
		}

		s.mu.Lock()
		canceled := s.canceled
		s.mu.Unlock()
		if canceled {
			panic("canceled but got context...")
		}

		s.mu.Lock()
		s.success++
		s.mu.Unlock()
	}
}

func (s *stresser) StopAndWait() {
	s.mu.Lock()
	fmt.Println("s.cancel() 1")
	s.cancel()
	fmt.Println("s.cancel() 2")
	s.canceled = true
	wg := s.wg
	s.mu.Unlock()

	fmt.Println("wg.Wait() 1")
	wg.Wait()
	fmt.Println("wg.Wait() 2")
}
