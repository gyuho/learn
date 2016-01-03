// This example demonstrates a priority queue built using the heap interface.
// Source: https://go.googlesource.com/go/+/master/src/container/heap/example_pq_test.go
//
package main

import (
	"container/heap"
	"fmt"
)

// An Item is something we manage in a priority queue.
type Item struct {
	id       string // The id of the item.
	priority int    // The priority of the item in the queue.

	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	// Highest priority comes at first in the array.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[:n-1]
	return item
}

func newPriorityQueue() *PriorityQueue {
	pq := &PriorityQueue{}
	heap.Init(pq)
	return pq
}

func (pq *PriorityQueue) push(x interface{}) {
	heap.Push(pq, x)
}

func (pq *PriorityQueue) top() *Item {
	if pq.Len() != 0 {
		return (*pq)[0]
	}
	return nil
}

func (pq *PriorityQueue) pop() *Item {
	x := heap.Pop(pq)
	n, _ := x.(*Item)
	return n
}

func (pq *PriorityQueue) replace(it *Item) bool {
	for i := range *pq {
		if (*pq)[i].id != it.id {
			continue
		}
		(*pq)[i] = it
		heap.Fix(pq, i)
		return true
	}
	return false
}

// This example creates a PriorityQueue with some items, adds and manipulates an item,
// and then removes the items in priority order.
func main() {
	var unsortedItems = []Item{
		Item{id: "banana", priority: 1},
		Item{id: "apple", priority: 5},
		Item{id: "pear", priority: 10},
	}

	pq := newPriorityQueue()
	for i := range unsortedItems {
		pq.push(&Item{
			id:       unsortedItems[i].id,
			priority: unsortedItems[i].priority,
			index:    i,
		})
	}
	fmt.Printf("After push all, pq.top(): %+v\n", pq.top())
	// After push all, pq.top(): &{id:pear priority:10 index:0}

	// Insert a new item and then modify its priority.
	pq.push(&Item{
		id:       "orange",
		priority: 1000,
	})
	fmt.Printf("After push new item, pq.top(): %+v\n", pq.top())
	// After push new item, pq.top(): &{id:orange priority:1000 index:0}

	(*pq)[0].priority = -10
	fmt.Printf("Before fix, pq.top(): %+v\n", pq.top())
	// Before fix, pq.top(): &{id:orange priority:-10 index:0}

	heap.Fix(pq, (*pq)[0].index)
	fmt.Printf("After fix, pq.top(): %+v\n", pq.top())
	// After fix, pq.top(): &{id:pear priority:10 index:0}

	(*pq)[0].priority = -100
	fmt.Printf("Before replace, pq.top(): %+v\n", pq.top())
	// Before replace, pq.top(): &{id:pear priority:-100 index:0}

	pq.replace((*pq)[0])
	fmt.Printf("After replace, pq.top(): %+v\n", pq.top())
	// After replace, pq.top(): &{id:apple priority:5 index:0}

	fmt.Println()

	// Take the items out; they arrive in decreasing priority order.
	for pq.Len() > 0 {
		// item := heap.Pop(pq).(*Item)
		item := pq.pop()
		fmt.Printf("%d:%s ", item.priority, item.id)
	}
	fmt.Println()
	// 5:apple 1:banana -10:orange -100:pear
}
