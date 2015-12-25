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

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, id string, priority int) {
	item.id = id
	item.priority = priority
	heap.Fix(pq, item.index)
}

// This example creates a PriorityQueue with some items, adds and manipulates an item,
// and then removes the items in priority order.
func main() {
	var unsortedItems = []Item{
		Item{id: "banana", priority: 1},
		Item{id: "apple", priority: 5},
		Item{id: "pear", priority: 10},
	}

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(PriorityQueue, len(unsortedItems))
	for i := range unsortedItems {
		pq[i] = &Item{
			id:       unsortedItems[i].id,
			priority: unsortedItems[i].priority,
			index:    i,
		}
	}

	fmt.Println()
	fmt.Printf("[BEFORE] heap.Init(&pq): %+v %+v %+v\n", pq[0], pq[1], pq[2])
	heap.Init(&pq)
	fmt.Printf("[AFTER]  heap.Init(&pq): %+v %+v %+v\n", pq[0], pq[1], pq[2])

	// Insert a new item and then modify its priority.
	item := &Item{
		id:       "orange",
		priority: 1,
	}
	heap.Push(&pq, item)

	fmt.Println()
	fmt.Printf("[BEFORE] pq.update: %+v %+v %+v %+v\n", pq[0], pq[1], pq[2], pq[3])
	pq.update(item, item.id, 5)
	fmt.Printf("[AFTER]  pq.update: %+v %+v %+v %+v\n", pq[0], pq[1], pq[2], pq[3])

	fmt.Println()
	// Take the items out; they arrive in decreasing priority order.
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%d:%s ", item.priority, item.id)
	}
	fmt.Println()
}

/*
[BEFORE] heap.Init(&pq): &{id:banana priority:1 index:0} &{id:apple priority:5 index:1} &{id:pear priority:10 index:2}
[AFTER]  heap.Init(&pq): &{id:pear priority:10 index:0} &{id:apple priority:5 index:1} &{id:banana priority:1 index:2}

[BEFORE] pq.update: &{id:pear priority:10 index:0} &{id:apple priority:5 index:1} &{id:banana priority:1 index:2} &{id:orange priority:1 index:3}
[AFTER]  pq.update: &{id:pear priority:10 index:0} &{id:apple priority:5 index:1} &{id:banana priority:1 index:2} &{id:orange priority:5 index:3}

10:pear 5:orange 5:apple 1:banana

*/
