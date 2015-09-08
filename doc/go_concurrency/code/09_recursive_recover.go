package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	keepRunning(5)
}

/*
Restarting after error: 2009-11-10 23:00:00 +0000 UTC
Restarting after error: 2009-11-10 23:00:00.001 +0000 UTC
Restarting after error: 2009-11-10 23:00:00.002 +0000 UTC
Restarting after error: 2009-11-10 23:00:00.003 +0000 UTC
Restarting after error: 2009-11-10 23:00:00.004 +0000 UTC
Too much panic: 5
2009/11/10 23:00:00 2009-11-10 23:00:00.004 +0000 UTC
*/

var count int

func keepRunning(limit int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Restarting after error:", err)

			time.Sleep(time.Millisecond)

			count++
			if count == limit {
				fmt.Printf("Too much panic: %d\n", count)
				log.Fatal(err)
			}
			keepRunning(limit)
		}
	}()
	run()
}

func run() {
	panic(time.Now().String())
}
