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
		go func(i int) {
			defer wg.Done()

			go func() {
				for limiter.Wait(ctx) == nil {
					mu.Lock()
					num++
					mu.Unlock()
				}
			}()
			// <-donec
		}(i)
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
