package main

import (
	"fmt"
	"sort"
)

// key/value pair of map[string]float64
type MapSF struct {
	key   string
	value float64
}

// Sort map pairs implementing sort.Interface
// to sort by value
type MapSFList []MapSF

// sort.Interface
// Define our custom sort: Swap, Len, Less
func (p MapSFList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p MapSFList) Len() int      { return len(p) }
func (p MapSFList) Less(i, j int) bool {
	return p[i].value < p[j].value
}

// Sort the struct from a map and return a MapSFList
func sortMapByValue(m map[string]float64) MapSFList {
	p := make(MapSFList, len(m))
	i := 0
	for k, v := range m {
		p[i] = MapSF{k, v}
		i++
	}
	sort.Sort(p)
	return p
}

func main() {
	// with sort.Interface and struct
	// we can automatically handle the duplicates
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

	if v, ok := sfmap["California"]; !ok {
		fmt.Println("California does not exist")
	} else {
		fmt.Println(v, "exists")
	}
	// 9.9 exists

	fmt.Println(sfmap["California"]) // 9.9

	if v, ok := sfmap["California2"]; !ok {
		fmt.Println("California2 does not exist")
	} else {
		fmt.Println(v, "exists")
	}
	// California2 does not exist

	delete(sfmap, "Ohio")
	if v, ok := sfmap["Ohio"]; !ok {
		fmt.Println("Ohio does not exist")
	} else {
		fmt.Println(v, "exists")
	}
	// Ohio does not exist
}
