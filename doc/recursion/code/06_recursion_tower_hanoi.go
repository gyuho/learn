package main

import "fmt"

// move num disks from src to dest
func hanoi(num int, src, aux, dest string) {
	if num > 0 {
		if num == 1 {
			fmt.Printf("%v -> %v\n", src, dest)
		} else {
			// order matters!
			hanoi(num-1, src, dest, aux)
			hanoi(1, src, aux, dest)
			hanoi(num-1, aux, src, dest)
		}
	}
}

/*
A -> C
A -> B
C -> B
A -> C
B -> A
B -> C
A -> C
*/

func main() {
	hanoi(3, "A", "B", "C")
}
