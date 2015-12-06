package main

import (
	"fmt"
	"time"
)

func main() {
	func() {
		tick := time.NewTicker(time.Second / 2)
		// defer tick.Stop()

		cnt := 0
		now := time.Now()
		for t := range tick.C {
			fmt.Println("took:", t.Sub(now))
			fmt.Println("took:", time.Since(now))
			now = time.Now()
			fmt.Println()

			cnt++
			if cnt == 5 {
				tick.Stop()
				break
			}
		}
		/*
		   took: 499.814697ms
		   took: 499.881208ms

		   took: 500.001328ms
		   took: 500.102576ms

		   took: 499.859404ms
		   took: 499.918472ms

		   ...
		*/
	}()

	func() {
		rate := time.Second / 10
		tick := time.NewTicker(rate)
		defer tick.Stop()

		burstNum := 5
		throttle := make(chan time.Time, burstNum)
		go func() {
			for t := range tick.C {
				throttle <- t
			}
		}()
		now := time.Now()
		for range []int{0, 0, 0, 0, 0, 0, 0, 0, 0} {
			// rate limit
			<-throttle
			fmt.Println("took:", time.Since(now))
			now = time.Now()
		}
		/*
		   took: 100.175925ms
		   took: 100.00098ms
		   took: 100.001815ms
		   took: 99.967231ms
		   took: 99.926131ms
		   took: 99.954405ms
		   took: 99.98205ms
		   took: 99.934787ms
		   took: 99.957386ms
		*/
	}()
}
