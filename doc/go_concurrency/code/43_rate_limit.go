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
		qps: 50000,
		N:   300,
	}

	go st.Start()
	time.Sleep(time.Second)
	fmt.Println("After Start:       st.success", st.success)
	st.StopAndWait()
	fmt.Println("After StopAndWait: st.success", st.success)

	fmt.Println()
	go st.Start()

	func() {
		st.StopAndWait()
		defer func() {
			go st.Start()
		}()
	}()

	time.Sleep(2 * time.Second)

	fmt.Println()
	st.StopAndWait()
	fmt.Println("st.success", st.success)
}

/*
After Start:       st.success 50001
StopAndWait: canceled!
Start finished with {}
StopAndWait: stopped!
After StopAndWait: st.success 50001

StopAndWait: canceled!
StopAndWait: stopped!

StopAndWait: canceled!
Start finished with {}
StopAndWait: stopped!
st.success 100018
*/

type stresser struct {
	qps int
	N   int

	mu          sync.Mutex
	wg          *sync.WaitGroup
	rateLimiter *rate.Limiter
	cancel      func()

	success int
}

func (s *stresser) Start() {
	ctx, cancel := context.WithCancel(context.Background())

	wg := &sync.WaitGroup{}
	wg.Add(s.N)

	s.mu.Lock()
	s.wg = wg
	s.rateLimiter = rate.NewLimiter(rate.Every(time.Second), s.qps)
	s.cancel = cancel
	s.mu.Unlock()

	for i := 0; i < s.N; i++ {
		go s.run(ctx)
	}

	fmt.Println("Start finished with", <-ctx.Done())
}

func (s *stresser) run(ctx context.Context) {
	defer s.wg.Done()

	for {
		if err := s.rateLimiter.Wait(ctx); err == context.Canceled {
			return
		}

		s.mu.Lock()
		s.success++
		s.mu.Unlock()
	}
}

func (s *stresser) StopAndWait() {
	s.mu.Lock()
	fmt.Println("StopAndWait: canceled!")
	s.cancel()
	s.wg.Wait()
	fmt.Println("StopAndWait: stopped!")
	s.mu.Unlock()

	// cancel, wg := s.cancel, s.wg
	// cancel()
	// wg.Wait()
}
