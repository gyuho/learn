package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())

	// nr := rand.New(rand.NewSource(time.Now().UnixNano()))
	// fmt.Println(nr.Float32())

	// fmt.Println(rand.Float32())
}

func main() {
	ents := []weightEntry{
		{weight: 0.05},
		{weight: 0.8},
		{weight: 0.15},
	}
	wt := createWeightTable(ents)

	fmt.Println(wt.choose())
	fmt.Println(wt.choose())
	fmt.Println(wt.choose())
	fmt.Println(wt.choose())
	fmt.Println(wt.choose())
	fmt.Println(wt.choose())
	fmt.Println(wt.choose())
	fmt.Println(wt.choose())
}

/*
{0.8 }
{0.8 }
{0.15 }
{0.8 }
{0.8 }
{0.8 }
{0.8 }
{0.8 }
*/

type weightEntry struct {
	weight float32
	name   string
}

type weightTable struct {
	entries       []weightEntry
	distributions []float32
}

func (wt weightTable) Len() int           { return len(wt.entries) }
func (wt weightTable) Swap(i, j int)      { wt.entries[i], wt.entries[j] = wt.entries[j], wt.entries[i] }
func (wt weightTable) Less(i, j int) bool { return wt.entries[i].weight < wt.entries[j].weight }

func createWeightTable(entries []weightEntry) weightTable {
	wt := weightTable{entries: entries}
	sort.Sort(wt)
	var cw float32
	for _, entry := range wt.entries {
		cw += entry.weight
		wt.distributions = append(wt.distributions, cw)
	}
	return wt
}

func (wt weightTable) choose() weightEntry {
	entryN := len(wt.entries)
	lastWeight := wt.entries[len(wt.entries)-1].weight
	idx := sort.Search(entryN, func(i int) bool {
		// returns the smallest index, of which distribution is
		// greater than random weight value
		return wt.distributions[i] >= rand.Float32()*lastWeight
	})
	return wt.entries[idx]
}
