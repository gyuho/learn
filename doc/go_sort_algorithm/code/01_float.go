package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []float64{5.4, 2.1, 3.5, 6.1, -10.5} // unsorted
	sort.Float64s(s)
	fmt.Println(s)
	// [-10.5 2.1 3.5 5.4 6.1]
}
