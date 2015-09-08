package main

import (
	"container/heap"
	"fmt"
)

type intMinHeap []int

func (s intMinHeap) Len() int           { return len(s) }
func (s intMinHeap) Less(i, j int) bool { return s[i] < s[j] }
func (s intMinHeap) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

//use pointer receivers
// because they modify the slice's length,
func (s *intMinHeap) Push(val interface{}) {
	*s = append(*s, val.(int))
}

func (s *intMinHeap) Pop() interface{} {
	heapSize := len(*s)
	lastNode := (*s)[heapSize-1]
	*s = (*s)[:heapSize-1]
	return lastNode
}

func main() {
	intSlice := &intMinHeap{12, 100, -15, 200, -5, 3, -12, 7}

	// make sure build Heap
	heap.Init(intSlice)

	heap.Push(intSlice, -10)
	heap.Push(intSlice, 17)

	fmt.Println("after build and push:", intSlice)

	// heap sort on Min-Heap
	// keep popping the last Node which is the smallest in the Min-Heap
	// root is the smallest but had been exchanged for pop operation
	// biggest is popped at the end
	for intSlice.Len() != 0 {
		fmt.Print(heap.Pop(intSlice), " ")
	}
	println()

	println()
	itemSlice := &itemMaxHeap{
		item{value: "Apple", priority: 5},
		item{value: "Google", priority: 10},
	}
	heap.Init(itemSlice)

	heap.Push(itemSlice, item{value: "Amazon", priority: 3})

	// no need to build again
	// because push does the build heap
	// heap.Init(itemSlice)

	for _, elem := range *itemSlice {
		fmt.Println("after push:", elem.value)
		fmt.Println("after push:", elem.priority)
		println()
	}
	for itemSlice.Len() != 0 {
		fmt.Print(heap.Pop(itemSlice), " ")
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

func (s *itemMaxHeap) Push(val interface{}) {
	*s = append(*s, val.(item))
}

func (s *itemMaxHeap) Pop() interface{} {
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
