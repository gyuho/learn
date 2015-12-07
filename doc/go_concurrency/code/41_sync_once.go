package main

import (
	"fmt"
	"sync"
)

type Mine struct {
	createOnce *sync.Once
}

func main() {
	m := Mine{}
	m.createOnce = &sync.Once{}
	onceBody := func() {
		fmt.Println("Only once")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {

			// m.createOnce = &sync.Once{}

			m.createOnce.Do(onceBody)

			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}

	fmt.Println()

	for i := 0; i < 10; i++ {
		go func() {

			m.createOnce = &sync.Once{}

			m.createOnce.Do(onceBody)

			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
}

/*
Only once

Only once
Only once
Only once
Only once
Only once
Only once
Only once
Only once
Only once
Only once
*/
