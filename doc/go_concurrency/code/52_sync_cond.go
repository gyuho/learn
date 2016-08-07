package main

import (
	"fmt"
	"sync"
)

type data struct {
	cond  *sync.Cond
	lines []string
}

var nums = []int{0, 1, 2}

func main() {
	d := data{
		cond:  &sync.Cond{L: &sync.Mutex{}},
		lines: make([]string, 0),
	}

	for i := range nums {
		go func(i int) {
			d.cond.L.Lock()
			d.lines = append(d.lines, fmt.Sprintf("%d: Hello World!", i))
			d.cond.L.Unlock()

			d.cond.Signal()
		}(i)
	}

	for {
		d.cond.L.Lock()
		if len(d.lines) != len(nums) {
			d.cond.Wait()
		} else {
			d.cond.L.Unlock()
			break
		}
		d.cond.L.Unlock()
	}

	fmt.Println(d.lines)
	// [2: Hello World! 0: Hello World! 1: Hello World!]
}
