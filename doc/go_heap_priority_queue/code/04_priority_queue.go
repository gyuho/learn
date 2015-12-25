package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	id       string // The id of the item.
	priority int    // The priority of the item in the queue.
	index    int    // The index of the item in the heap.
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
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

func main() {
	var unsortedItems = []Item{
		Item{id: "banana", priority: 1},
		Item{id: "apple", priority: 5},
		Item{id: "pear", priority: 10},
	}

	pq := PriorityQueue{}
	heap.Init(&pq)

	for i := range unsortedItems {
		item := &Item{
			id:       unsortedItems[i].id,
			priority: unsortedItems[i].priority,
			index:    i,
		}
		heap.Push(&pq, item)
	}
	fmt.Printf("[AFTER]  heap.Init(&pq): %+v %+v %+v\n", pq[0], pq[1], pq[2])

	// Insert a new item and then modify its priority.
	item := &Item{
		id:       "orange",
		priority: 1,
	}
	heap.Push(&pq, item)

	fmt.Println()
	pq[0].priority = -10
	fmt.Printf("[BEFORE] pq.update: %+v %+v %+v %+v\n", pq[0], pq[1], pq[2], pq[3])

	heap.Fix(&pq, pq[0].index)
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
[AFTER]  heap.Init(&pq): &{id:pear priority:10 index:0} &{id:banana priority:1 index:1} &{id:apple priority:5 index:2}

[BEFORE] pq.update: &{id:pear priority:-10 index:0} &{id:banana priority:1 index:1} &{id:apple priority:5 index:2} &{id:orange priority:1 index:3}
[AFTER]  pq.update: &{id:apple priority:5 index:0} &{id:banana priority:1 index:1} &{id:pear priority:-10 index:2} &{id:orange priority:1 index:3}

5:apple 1:orange 1:banana -10:pear

*/
