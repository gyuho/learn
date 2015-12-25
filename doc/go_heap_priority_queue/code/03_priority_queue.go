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
	fmt.Printf("[AFTER]  heap.push: %+v %+v %+v\n", (*pq)[0], (*pq)[1], (*pq)[2])

	fmt.Println()

	// Insert a new item and then modify its priority.
	item := &Item{
		id:       "orange",
		priority: 1,
	}
	pq.push(item)
	(*pq)[0].priority = -10
	fmt.Printf("[BEFORE] heap.Fix: %+v %+v %+v %+v\n", (*pq)[0], (*pq)[1], (*pq)[2], (*pq)[3])

	heap.Fix(pq, (*pq)[0].index)
	fmt.Printf("[AFTER]  heap.Fix: %+v %+v %+v %+v\n", (*pq)[0], (*pq)[1], (*pq)[2], (*pq)[3])
	fmt.Printf("[AFTER]  heap.Fix pq.top(): %+v\n", pq.top())

	fmt.Println()

	// Take the items out; they arrive in decreasing priority order.
	for pq.Len() > 0 {
		// item := heap.Pop(pq).(*Item)
		item := pq.pop()
		fmt.Printf("%d:%s ", item.priority, item.id)
	}
}

/*
[AFTER]  heap.push: &{id:pear priority:10 index:0} &{id:banana priority:1 index:1} &{id:apple priority:5 index:2}

[BEFORE] heap.Fix: &{id:pear priority:-10 index:0} &{id:banana priority:1 index:1} &{id:apple priority:5 index:2} &{id:orange priority:1 index:3}
[AFTER]  heap.Fix: &{id:apple priority:5 index:0} &{id:banana priority:1 index:1} &{id:pear priority:-10 index:2} &{id:orange priority:1 index:3}
[AFTER]  heap.Fix pq.top(): &{id:apple priority:5 index:0}

5:apple 1:orange 1:banana -10:pear

*/
