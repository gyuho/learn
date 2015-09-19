package main

import (
	"fmt"
	"log"
)

func main() {
	bufferedSenderChan := make(chan<- int, 3)
	bufferedReceiverChan := make(<-chan int, 3)

	bufferedSenderChan <- 0
	bufferedSenderChan <- 1
	bufferedSenderChan <- 2

	// defer func() {
	// 	if err := recover(); err != nil {
	// 		fmt.Println(err)
	// 	}
	// }()
	// panic(1)

	// You cannot recover from deadlock!
	// <-bufferedReceiverChan
	// fatal error: all goroutines are asleep - deadlock!

	// 	close(bufferedReceiverChan) // (cannot close receive-only channel)
	// 	fmt.Println(<-bufferedReceiverChan)
	_ = bufferedReceiverChan

	bufferedChan := make(chan int, 3)
	bufferedChan <- 0
	bufferedChan <- 1
	bufferedChan <- 2
	fmt.Println(<-bufferedChan)
	fmt.Println(<-bufferedChan)
	fmt.Println(<-bufferedChan)
	/*
	   0
	   1
	   2
	*/

	fmt.Println()
	for i := 0; i < 10; i++ {
		go func(i int) {
			bufferedChan <- i
		}(i)
	}
	for i := 0; i < 10; i++ {
		fmt.Printf("%v ", <-bufferedChan)
	}
	fmt.Println()
	/*
	   9 0 1 6 7 5 2 3 8 4
	*/

	fmt.Println()
	slice := []float64{23.0, 23, 23, -123.2, 23, 123.2, -2.2, 23.1, -101.2, 17.2}
	sum := 0.0
	for _, elem := range slice {
		sum += elem
	}

	counter1 := NewChannelCounter(0)
	defer counter1.Done()
	defer counter1.Close()

	for _, elem := range slice {
		counter1.Add(elem)
	}
	val1 := counter1.Get()
	if val1 != sum {
		log.Fatalf("NewChannelCounter with No Buffer got wrong. Expected %v but got %v\n", sum, val1)
	}

	counter2 := NewChannelCounter(10)
	defer counter2.Done()
	defer counter2.Close()

	for _, elem := range slice {
		counter2.Add(elem)
	}
	val2 := counter2.Get()
	if val2 != sum {
		log.Fatalf("NewChannelCounter with Buffer got wrong. Expected %v but got %v\n", sum, val2)
	}

	// 2015/08/08 14:03:24 NewChannelCounter with Buffer got wrong. Expected 28.167699999999993 but got 23
}

// Counter is an interface for counting.
// It contains counting data as long as a type
// implements all the methods in the interface.
type Counter interface {
	// Get returns the current count.
	Get() float64

	// Add adds the delta value to the counter.
	Add(delta float64)
}

// ChannelCounter counts through channels.
type ChannelCounter struct {
	valueChan chan float64
	deltaChan chan float64
	done      chan struct{}
}

func NewChannelCounter(buf int) *ChannelCounter {
	c := &ChannelCounter{
		make(chan float64, buf),
		make(chan float64, buf),
		make(chan struct{}),
	}
	go c.Run()
	return c
}

func (c *ChannelCounter) Run() {

	var value float64

	for {
		// "select" statement chooses which of a set of
		// possible send or receive operations will proceed.
		select {

		case delta := <-c.deltaChan:
			value += delta

		case <-c.done:
			return

		case c.valueChan <- value:
			// Do nothing.

			// If there is no default case, the "select" statement
			// blocks until at least one of the communications can proceed.
		}
	}
}

func (c *ChannelCounter) Get() float64 {
	return <-c.valueChan
}

func (c *ChannelCounter) Add(delta float64) {
	c.deltaChan <- delta
}

func (c *ChannelCounter) Done() {
	c.done <- struct{}{}
}

func (c *ChannelCounter) Close() {
	close(c.deltaChan)
}
