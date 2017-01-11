package main

import (
	"container/heap"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	src := rand.NewSource(time.Now().UnixNano())
	random := rand.New(src)
	latencies := make([]float64, 5000000)
	for i := range latencies {
		latencies[i] = random.Float64()
	}

	var s1 []float64
	fmt.Println("s1 add")
	now := time.Now()
	for i := range latencies {
		s1 = append(s1, latencies[i])
	}
	fmt.Println("s1 add done; took", time.Since(now))
	fmt.Println("s1 sort start")
	now = time.Now()
	sort.Float64s(s1)
	fmt.Println("s1 sort done; took", time.Since(now))

	s2 := &float64Heap{}
	heap.Init(s2)
	fmt.Println("s2 add")
	now = time.Now()
	for i := range latencies {
		heap.Push(s2, latencies[i])
	}
	fmt.Println("s2 add done; took", time.Since(now))
	fmt.Println("s2 sort start")
	now = time.Now()
	s2rs := make([]float64, len(latencies))
	idx := 0
	for s2.Len() > 0 {
		s2rs[idx] = heap.Pop(s2).(float64)
		idx++
	}
	fmt.Println("s2 sort done; took", time.Since(now))

	fmt.Println("s2 is sorted?", sort.IsSorted(sort.Float64Slice(s2rs)))
	/*
	   s1 add
	   s1 add done; took 56.105544ms
	   s1 sort start
	   s1 sort done; took 1.261716659s
	   s2 add
	   s2 add done; took 349.446154ms
	   s2 sort start
	   s2 sort done; took 3.335775846s
	   s2 is sorted? true
	*/
	// just use sort.Sort unless it needs logN Top method
}

// min-heap of float64
type float64Heap []float64

func (h float64Heap) Len() int           { return len(h) }
func (h float64Heap) Less(i, j int) bool { return h[i] < h[j] }
func (h float64Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *float64Heap) Push(x interface{}) {
	*h = append(*h, x.(float64))
}

func (h *float64Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
