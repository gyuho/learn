// A child is running up a staircase with n steps,
// and can hop either 1 step, 2 steps, or 3 steps at a time.
// Implement a method to count how many possible ways
// the child can run up the stairs.
package main

import "fmt"

func countRecursive(n int) int {
	if n < 0 {
		return 0
	} else if n == 0 {
		return 1
	}
	return countRecursive(n-1) + countRecursive(n-2) + countRecursive(n-3)
}

func countDynamic(n int, mm map[int]int) int {
	if n < 0 {
		return 0
	} else if n == 0 {
		return 1
	} else if mm[n] > 0 {
		return mm[n]
	}
	mm[n] = countDynamic(n-1, mm) +
		countDynamic(n-2, mm) +
		countDynamic(n-3, mm)
	return mm[n]
}

func main() {
	fmt.Println(countRecursive(10)) // 274

	nmap := make(map[int]int)
	fmt.Println(countDynamic(10, nmap)) // 274
	fmt.Println(nmap)
	// map[2:2 4:7 6:24 8:81 10:274 1:1 3:4 5:13 7:44 9:149]
}
