package main

// priorityQueue is a min-heap of Jobs.
type priorityQueue []*job

func (pq priorityQueue) Len() int {
	return len(pq)
}

func (pq priorityQueue) Less(i, j int) bool {
	p1 := pq[i].Priority
	idx1 := pq[i].ETCDIndex

	p2 := pq[j].Priority
	idx2 := pq[j].ETCDIndex

	if p1 == p2 {
		// min-heap returns the lowest priority first
		// when the Priority's were same, we want to return the one with lower index.
		return idx1 < idx2
	}

	// max-heap
	return p1 > p2
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*job))
}

func (pq *priorityQueue) Pop() interface{} {
	heapSize := len(*pq)
	lastNode := (*pq)[heapSize-1]
	*pq = (*pq)[:heapSize-1]
	return lastNode
}
