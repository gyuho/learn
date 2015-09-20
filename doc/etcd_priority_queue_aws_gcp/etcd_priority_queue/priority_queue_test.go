package main

import (
	"container/heap"
	"testing"
)

func TestHeap(t *testing.T) {
	job1 := &job{}
	job1.Action = "action1"
	job1.ETCDIndex = 1
	job1.Priority = 11.0

	job2 := &job{}
	job2.Action = "action2"
	job2.ETCDIndex = 2
	job2.Priority = 0.0

	job3 := &job{}
	job3.Action = "action3"
	job3.ETCDIndex = 3
	job3.Priority = 0.0

	job4 := &job{}
	job4.Action = "action4"
	job4.ETCDIndex = 4
	job4.Priority = 1000.0

	job5 := &job{}
	job5.Action = "action5"
	job5.ETCDIndex = 5
	job5.Priority = 2.3

	var pq priorityQueue
	heap.Push(&pq, job1)
	heap.Push(&pq, job2)
	heap.Push(&pq, job3)
	heap.Push(&pq, job4)
	heap.Push(&pq, job5)

	heap.Init(&pq)

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*job)
		t.Logf("%.2f:%s\n", item.Priority, item.Action)
		if pq.Len() == 0 {
			if item.Action != "action3" {
				t.Fatalf("The least prioritized item should be action3 but %+v", item)
			}
		}
	}
}
