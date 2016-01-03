package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
)

func strToFloat64(str string) float64 {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func main() {
	rows := [][]string{
		[]string{"cdomain.com", "3", "-5.02", "aaa", "aaa"},
		[]string{"cdomain.com", "2", "133.02", "aaa", "aaa"},
		[]string{"cdomain.com", "1", "1.02", "aaa", "aaa"},
		[]string{"bdomain.com", "2", "23.02", "aaa", "aaa"},
		[]string{"bdomain.com", "1", "12.02", "aaa", "aaa"},
		[]string{"bdomain.com", "3", "53.02", "aaa", "aaa"},
		[]string{"adomain.com", "5", "32.1232", "aaa", "aaa"},
		[]string{"adomain.com", "3", "2.02202", "aaa", "aaa"},
		[]string{"adomain.com", "1", "511.02", "aaa", "aaa"},
	}
	ascendingName0 := func(row1, row2 *[]string) bool {
		return (*row1)[0] < (*row2)[0]
	}
	descendingVal := func(row1, row2 *[]string) bool {
		return strToFloat64((*row1)[2]) > strToFloat64((*row2)[2])
	}
	ascendingName1 := func(row1, row2 *[]string) bool {
		return (*row1)[1] < (*row2)[1]
	}
	by(rows, ascendingName0, descendingVal, ascendingName1).Sort(rows)
	rs := fmt.Sprintf("%v", rows)
	if rs != "[[adomain.com 1 511.02 aaa aaa] [adomain.com 5 32.1232 aaa aaa] [adomain.com 3 2.02202 aaa aaa] [bdomain.com 3 53.02 aaa aaa] [bdomain.com 2 23.02 aaa aaa] [bdomain.com 1 12.02 aaa aaa] [cdomain.com 2 133.02 aaa aaa] [cdomain.com 1 1.02 aaa aaa] [cdomain.com 3 -5.02 aaa aaa]]" {
		fmt.Errorf("%v", rows)
	}
}

// by returns a multiSorter that sorts using the less functions
func by(rows [][]string, lesses ...lessFunc) *multiSorter {
	return &multiSorter{
		data: rows,
		less: lesses,
	}
}

// lessFunc compares between two string slices.
type lessFunc func(p1, p2 *[]string) bool

func makeAscendingFunc(idx int) func(row1, row2 *[]string) bool {
	return func(row1, row2 *[]string) bool {
		return (*row1)[idx] < (*row2)[idx]
	}
}

// multiSorter implements the Sort interface
// , sorting the two dimensional string slices within.
type multiSorter struct {
	data [][]string
	less []lessFunc
}

// Sort sorts the rows according to lessFunc.
func (ms *multiSorter) Sort(rows [][]string) {
	sort.Sort(ms)
}

// Len is part of sort.Interface.
func (ms *multiSorter) Len() int {
	return len(ms.data)
}

// Swap is part of sort.Interface.
func (ms *multiSorter) Swap(i, j int) {
	ms.data[i], ms.data[j] = ms.data[j], ms.data[i]
}

// Less is part of sort.Interface.
func (ms *multiSorter) Less(i, j int) bool {
	p, q := &ms.data[i], &ms.data[j]
	var k int
	for k = 0; k < len(ms.less)-1; k++ {
		less := ms.less[k]
		switch {
		case less(p, q):
			// p < q
			return true
		case less(q, p):
			// p > q
			return false
		}
		// p == q; try next comparison
	}
	return ms.less[k](p, q)
}
