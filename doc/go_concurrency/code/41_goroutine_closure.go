package main

import (
	"fmt"
	"time"
)

func main() {
	for _, d := range []int{1, 2} {
		x := d
		func() {
			fmt.Printf("%d(x: %d)\n", d, x)
		}()
	}
	time.Sleep(time.Second)
	fmt.Println()
	/*
		1(x: 1)
		2(x: 2)
	*/

	// (X) DON'T DO THIS
	for _, d := range []int{10, 20} {
		x := d
		go func() {
			fmt.Printf("%d(x: %d)\n", d, x)
		}()
	}
	time.Sleep(time.Second)
	fmt.Println()
	/*
	   20(x: 10)
	   20(x: 20)
	*/

	for _, d := range []int{100, 200} {
		go func(d int) {
			fmt.Printf("%d\n", d)
		}(d)
	}
	time.Sleep(time.Second)
	fmt.Println()
	/*
	   200
	   100
	*/

	// https://github.com/coreos/etcd/pull/3880#issuecomment-157442671
	// 'forever' is first evaluated without creating a new goroutine.
	// And then 'wrap(forever())' is evaluated with a new goroutine.
	go func() { wrap(forever()) }()
	// calling forever...
	// this is running in the background(goroutine)

	// 'forever' is first evaluated without creating a new goroutine.
	// There is no goroutine created to run this in the background.
	// So this is blocking forever!!!
	go wrap(forever())
	// calling forever...
}

func forever() error {
	fmt.Println("calling forever...")
	time.Sleep(time.Hour)
	return nil
}

func wrap(err error) {
	_ = err
}
