package main

import (
	"fmt"
	"sort"
)

// heap is from Go's standard container/heap package
// https://go.googlesource.com/go/+/master/src/container/heap/heap.go
type heap interface {
	sort.Interface          // embed sort then implement your own: Swap, Len, Less
	push(value interface{}) // add to heap
	pop() interface{}       // pop from heap (depends on Min,Max-Heap)
}

// down heapify downwards
func down(h heap, idx, heapSize int) {
	for {
		// minimum(maximum) element in the tree is the root, at index 0.

		// parent := idx / 2
		left := 2*idx + 1
		right := 2*idx + 2

		// no need to heapify (already heapified) / overflow
		if left >= heapSize || left < 0 {
			break
		}

		// Assume that you implement Min-Heap (right should be bigger)

		// just make it consistent that left is smaller in Min-Heap
		// (Max-Heap: left is bigger)
		if right < heapSize && h.Less(right, left) {
			left = right
		}

		// no need to heapify (already heapified)
		if h.Less(idx, left) {
			break
		}

		h.Swap(idx, left)
		idx = left

		// keep heapifying downwards
	}
}

// up heapify upwards
func up(h heap, idx int) {
	for {
		parent := (idx - 1) / 2

		// Assume that you implement Min-Heap (parent should be smaller)

		// no need to heapify (already heapified)
		if parent == idx || h.Less(parent, idx) {
			break
		}

		h.Swap(idx, parent)
		idx = parent

		// keep heapifying upwards
	}
}

// build : O(n) where n = h.Len()
func build(h heap) {
	heapSize := h.Len()
	for idx := heapSize/2 - 1; idx >= 0; idx-- {
		down(h, idx, heapSize)
	}
}

// to satisfy the heap interface
// any type that uses heap interface
// will need to implement push method.
// O(log(n)) where n = h.Len()
func push(h heap, val interface{}) {
	h.push(val)
	// heapify from bottom
	up(h, h.Len()-1)
}

// to satisfy the heap interface
// any type that uses heap interface
// will need to implement pop method.
// O(log(n)) where n = h.Len()
func pop(h heap) interface{} {
	lastIdx := h.Len() - 1

	// assume the min-heap
	// then pop returns the minimum

	// exchange the one to pop at root 0
	// with the one in the last node
	h.Swap(0, lastIdx)

	heapSize := lastIdx

	// heapify except the lastIdx node
	down(h, 0, heapSize)

	return h.pop()
}

// Note that Go source code heap implements "MIN" heap
// Go heap embeds sort.Interface
// Therefore, we need to define our custom Interface

type intMinHeap []int

func (s intMinHeap) Len() int           { return len(s) }
func (s intMinHeap) Less(i, j int) bool { return s[i] < s[j] }
func (s intMinHeap) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

//use pointer receivers
// because they modify the slice's length,
func (s *intMinHeap) push(val interface{}) {
	*s = append(*s, val.(int))
}

func (s *intMinHeap) pop() interface{} {
	heapSize := len(*s)
	lastNode := (*s)[heapSize-1]
	*s = (*s)[:heapSize-1]
	return lastNode
}

func main() {
	intSlice := &intMinHeap{12, 100, -15, 200, -5, 3, -12, 7}

	// make sure build Heap
	build(intSlice)

	// intSlice.push(-10) (X) is only interface without heapifying
	push(intSlice, -10)
	push(intSlice, 17)

	fmt.Println("after push and build:", intSlice)

	// heap sort on Min-Heap
	// keep popping the last Node which is the smallest in the Min-Heap
	// root is the smallest but had been exchanged for pop operation
	// biggest is popped at the end
	for intSlice.Len() != 0 {
		fmt.Print(pop(intSlice), " ")
	}
	println()

	println()
	itemSlice := &itemMaxHeap{
		item{value: "Apple", priority: 5},
		item{value: "Google", priority: 10},
	}
	build(itemSlice)

	push(itemSlice, item{value: "Amazon", priority: 3})

	// no need to build again
	// because push does the build heap
	// build(itemSlice)

	for _, elem := range *itemSlice {
		fmt.Println("after push:", elem.value)
		fmt.Println("after push:", elem.priority)
		println()
	}
	for itemSlice.Len() != 0 {
		fmt.Print(pop(itemSlice), " ")
	}
	println()
}

type item struct {
	value    string
	priority int
}

type itemMaxHeap []item

func (s itemMaxHeap) Len() int           { return len(s) }
func (s itemMaxHeap) Less(i, j int) bool { return s[i].priority > s[j].priority } // Max-Heap
func (s itemMaxHeap) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func (s *itemMaxHeap) push(val interface{}) {
	*s = append(*s, val.(item))
}

func (s *itemMaxHeap) pop() interface{} {
	heapSize := len(*s)
	lastNode := (*s)[heapSize-1]
	*s = (*s)[0 : heapSize-1]
	return lastNode
}

/*
after push and build: &[-15 -5 -12 7 12 3 -10 100 200 17]
-15 -12 -10 -5 3 7 12 17 100 200
after push: Google
after push: 10
after push: Apple
after push: 5
after push: Amazon
after push: 3
{Google 10} {Apple 5} {Amazon 3}
*/
