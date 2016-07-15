package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	st := time.Now()
	i := 0
	limiter := rate.NewLimiter(rate.Every(time.Second), 100)
	ctx, cancel := context.WithTimeout(context.TODO(), 2*time.Second)
	for limiter.Wait(ctx) == nil {
		i++
	}
	cancel()
	fmt.Println(i, "DONE. Took", time.Since(st))
	// 101 DONE. Took 1.00013873s
}
