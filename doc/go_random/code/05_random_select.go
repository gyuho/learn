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
{0.8 }
{0.15 }
{0.15 }
{0.8 }
{0.8 }
{0.8 }
*/

type weightEntry struct {
	weight float32
	name   string
}

type weightTable struct {
	entries    []weightEntry
	sumWeights float32
}

func (wt weightTable) Len() int           { return len(wt.entries) }
func (wt weightTable) Swap(i, j int)      { wt.entries[i], wt.entries[j] = wt.entries[j], wt.entries[i] }
func (wt weightTable) Less(i, j int) bool { return wt.entries[i].weight < wt.entries[j].weight }

func createWeightTable(entries []weightEntry) *weightTable {
	wt := weightTable{entries: entries}
	sort.Sort(wt)
	for _, entry := range wt.entries {
		wt.sumWeights += entry.weight
	}
	return &wt
}

func (wt *weightTable) choose() weightEntry {
	v := rand.Float32() * wt.sumWeights
	var sum float32
	var idx int
	for i := range wt.entries {
		sum += wt.entries[i].weight
		if sum >= v {
			idx = i
			break
		}
	}
	return wt.entries[idx]
}
