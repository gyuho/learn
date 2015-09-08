package main

import (
	"fmt"
	"sort"
)

type keyValue struct {
	key   string
	value float64
}

type keyValueSlice []keyValue

func (p keyValueSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p keyValueSlice) Len() int      { return len(p) }
func (p keyValueSlice) Less(i, j int) bool {
	return p[i].value < p[j].value
}

func sortMapByValue(m map[string]float64) keyValueSlice {
	p := make(keyValueSlice, len(m))
	i := 0
	for k, v := range m {
		p[i] = keyValue{k, v}
		i++
	}
	sort.Sort(p)
	return p
}

func main() {
	sfmap := map[string]float64{
		"California":    9.9,
		"Japan":         7.23,
		"Korea":         -.3,
		"Hello":         1.5,
		"USA":           8.4,
		"San Francisco": 8.4,
		"Ohio":          -1.10,
		"New York":      1.23,
		"Los Angeles":   23.1,
		"Mountain View": 9.9,
	}
	fmt.Println(sortMapByValue(sfmap), len(sortMapByValue(sfmap)))
	// [{Ohio -1.1} {Korea -0.3} {New York 1.23} {Hello 1.5}
	// {Japan 7.23} {USA 8.4} {San Francisco 8.4} {California 9.9}
	// {Mountain View 9.9} {Los Angeles 23.1}] 10
}
