package main

import "fmt"

func main() {
	ch := make(chan int)

	for i := 0; i < 5; i++ {
		go func() {
			ch <- i
		}()
	}

	for v := range ch {
		fmt.Println(v)
	}
}

/*
5
5
5
5
5
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
main.main()
	/tmp/sandbox982202598/main.go:14 +0x1e0
*/
