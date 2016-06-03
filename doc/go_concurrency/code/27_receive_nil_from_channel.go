package main

import (
	"fmt"
	"time"
)

func main() {
	{
		errChan := make(chan error)
		go func() {
			errChan <- nil
		}()
		select {
		case v := <-errChan:
			fmt.Println("even if nil, it still receives", v)
		case <-time.After(time.Second):
			fmt.Println("time-out!")
		}
		// even if nil, it still receives <nil>
	}

	{
		errChan := make(chan error)
		errChan = nil
		go func() {
			errChan <- nil
		}()
		select {
		case v := <-errChan:
			fmt.Println("even if nil, it still receives", v)
		case <-time.After(time.Second):
			fmt.Println("time-out!")
		}
		// time-out!
	}
}
