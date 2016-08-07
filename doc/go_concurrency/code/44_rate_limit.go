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
		qps: 100,
		N:   10,
	}

	go st.start()

	time.Sleep(2 * time.Second)
	st.StopAndWait()
	fmt.Println("success", st.success) // success 299
}

type stresser struct {
	qps int
	N   int

	mu          sync.Mutex
	wg          *sync.WaitGroup
	rateLimiter *rate.Limiter
	cancel      func()

	success int
}

func (s *stresser) start() {
	ctx, cancel := context.WithCancel(context.Background())

	wg := &sync.WaitGroup{}
	wg.Add(s.N)

	s.mu.Lock()
	s.wg = wg
	s.rateLimiter = rate.NewLimiter(rate.Limit(s.qps), s.qps)
	s.cancel = cancel
	s.mu.Unlock()

	for i := 0; i < s.N; i++ {
		go s.run(ctx)
	}

	<-ctx.Done()
}

func (s *stresser) run(ctx context.Context) {
	defer s.wg.Done()

	for {
		if err := s.rateLimiter.Wait(ctx); err == context.Canceled || ctx.Err() == context.Canceled {
			return
		}

		s.mu.Lock()
		s.success++
		s.mu.Unlock()
	}
}

func (s *stresser) StopAndWait() {
	s.mu.Lock()
	s.cancel()
	wg := s.wg
	s.mu.Unlock()

	wg.Wait()
}
