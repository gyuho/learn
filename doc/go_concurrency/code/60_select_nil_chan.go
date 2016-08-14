package main

import "fmt"

func main() {
	var c chan struct{}
	select { // select doesn't select nil channel, but without default it panics
	case <-c:
		panic(1)
	default:
		fmt.Println(c == nil) // true
	}
}
