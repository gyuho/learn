package merge_sort

import (
	"math/rand"
	"sort"
	"testing"
)

func merge(s1, s2 []int) []int {
	final := make([]int, len(s1)+len(s2))
	i, j := 0, 0
	for i < len(s1) && j < len(s2) {
		if s1[i] <= s2[j] {
			final[i+j] = s1[i]
			i++
			continue
		}
		final[i+j] = s2[j]
		j++
	}
	for i < len(s1) {
		final[i+j] = s1[i]
		i++
	}
	for j < len(s2) {
		final[i+j] = s2[j]
		j++
	}
	return final
}

func mergeSort(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}
	idx := len(slice) / 2
	left := mergeSort(slice[:idx])
	right := mergeSort(slice[idx:])
	return merge(left, right)
}

func concurrentMergeSort(slice []int, result chan []int) {
	if len(slice) < 2 {
		result <- slice
		return
	}
	idx := len(slice) / 2
	ch1, ch2 := make(chan []int), make(chan []int)

	go concurrentMergeSort(slice[:idx], ch1)
	go concurrentMergeSort(slice[idx:], ch2)

	left := <-ch1
	right := <-ch2

	result <- merge(left, right)
}

func BenchmarkStandardPackage(b *testing.B) {
	var sampleIntSlice = []int{}
	sampleIntSlice = rand.New(rand.NewSource(123123)).Perm(999999)
	for i := 0; i < b.N; i++ {
		sort.Ints(sampleIntSlice)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	var sampleIntSlice = []int{}
	sampleIntSlice = rand.New(rand.NewSource(123123)).Perm(999999)
	for i := 0; i < b.N; i++ {
		mergeSort(sampleIntSlice)
	}
}

func BenchmarkConcurrentMergeSort(b *testing.B) {
	var sampleIntSlice = []int{}
	sampleIntSlice = rand.New(rand.NewSource(123123)).Perm(999999)
	for i := 0; i < b.N; i++ {
		result := make(chan []int)
		go concurrentMergeSort(sampleIntSlice, result)
		<-result
	}
}

/*
go get github.com/cespare/prettybench
go test -bench . -benchmem -cpu 1,2,4 | prettybench
benchmark                        iter       time/iter      bytes alloc              allocs
---------                        ----       ---------      -----------              ------
BenchmarkStandardPackage           10    131.37 ms/op      800929 B/op         1 allocs/op
BenchmarkStandardPackage-2         10    132.84 ms/op      800929 B/op         1 allocs/op
BenchmarkStandardPackage-4         10    131.81 ms/op      800929 B/op         1 allocs/op
BenchmarkMergeSort                  5    204.42 ms/op   166369507 B/op    999998 allocs/op
BenchmarkMergeSort-2                5    202.03 ms/op   166369507 B/op    999998 allocs/op
BenchmarkMergeSort-4                5    229.62 ms/op   166369507 B/op    999998 allocs/op
BenchmarkConcurrentMergeSort        1   3994.73 ms/op   488144848 B/op   3537113 allocs/op
BenchmarkConcurrentMergeSort-2      1   2134.87 ms/op   377522704 B/op   3199159 allocs/op
BenchmarkConcurrentMergeSort-4      1   1242.12 ms/op   377254480 B/op   3194968 allocs/op
ok  	github.com/gyuho/kway/benchmarks/merge_sort	18.784s
*/
