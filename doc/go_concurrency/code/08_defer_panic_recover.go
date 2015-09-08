package main

import "fmt"

func main() {
	doPanic()
	// error: 1 and recovered
}

func doRecover() {
	if err := recover(); err != nil {
		fmt.Println("error:", err, "and recovered")
	}
}

func doPanic() {
	defer doRecover()
	panic(1)
}
