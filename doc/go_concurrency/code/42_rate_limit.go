package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	var (
		num int
		mu  sync.Mutex

		qps = 10
		wg  sync.WaitGroup
		N   = 10000
	)

	wg.Add(N)

	ctx, cancel := context.WithCancel(context.Background())

	limiter := rate.NewLimiter(rate.Every(time.Second), qps)

	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()

			for {
				if err := limiter.Wait(ctx); err == context.Canceled {
					return
				}

				mu.Lock()
				num++
				mu.Unlock()
			}
		}()
	}

	time.Sleep(time.Second)
	mu.Lock()
	fmt.Println("num:", num)
	mu.Unlock()

	fmt.Println("burst:", limiter.Burst())

	fmt.Println("canceling...")
	cancel()
	wg.Wait()
	fmt.Println("Done!")
}

/*
num: 11
burst: 10
canceling...
Done!
*/
