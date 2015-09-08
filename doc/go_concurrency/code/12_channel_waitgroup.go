package main

import "sync"

func main() {
	ch := make(chan struct{})
	var wg sync.WaitGroup

	go func() {
		println(1)
		ch <- struct{}{}
	}()

	wg.Add(1)
	go func() {
		println(2)
		wg.Done()
	}()

	<-ch
	wg.Wait()

	// 1
	// 2
}
