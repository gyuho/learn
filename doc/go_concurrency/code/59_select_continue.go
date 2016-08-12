package main

import "fmt"

func main() {
	donec := make(chan struct{})
	for i := range []int{1, 2, 3} {
		fmt.Println(i)
		if i == 1 {
			close(donec)
		}
		select {
		case <-donec:
			continue
		default:
		}
		fmt.Println("hey")
	}
}

/*
0
hey
1
2
*/
