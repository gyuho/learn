package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strings"
)

func open(filename string) *os.File {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Unable to read file: %+v", err)
	}
	return f
}

func main() {
	pFunc := Probability(open("sample.txt"))
	rs := Get("IamCodingwithGO.Ireallyenjoythislanguage.", pFunc)
	fmt.Println(rs)
	// [IamCodingwithGO.I really enjoy this language .]
}

// probability calculates frequency probability.
// For each word, divide frequency by the number of words.
func probability(reader io.Reader) map[string]float64 {
	scanner := bufio.NewScanner(reader)
	//
	// This must be called before Scan.
	// The default split function is bufio.ScanLines.
	scanner.Split(bufio.ScanWords)
	//
	pmap := make(map[string]float64)
	//
	var length float64
	for scanner.Scan() {
		// Remove all leading and trailing Unicode code points.
		word := strings.Trim(scanner.Text(), ",-!;:\"?.")
		if _, exist := pmap[word]; exist {
			pmap[word]++
		} else {
			pmap[word] = 1
		}
		// keep increasing while reading(scanning)
		length++
	}
	for k, v := range pmap {
		pmap[k] = v / length
	}
	return pmap
}

// Probability returns the word probability,
// with smoothing.
func Probability(reader io.Reader) func(string) float64 {
	pmap := probability(reader)
	return func(word string) float64 {
		if score, ok := pmap[word]; ok {
			return score
		}
		// if the word has never showed up, smooth.
		return 10 / (float64(len(pmap)) * math.Pow(10, float64(len(word))))
	}
}

type split struct {
	Head string
	Tail string
}

func doSplit(txt string) []split {
	splits := []split{}
	for i := range txt {
		splits = append(splits, split{txt[:i+1], txt[i+1:]})
	}
	return splits
}

func mostPlausible(chunks [][]string, probFunc func(string) float64) []string {
	chunk := []string{}
	min := -1 * math.MaxFloat64
	for _, words := range chunks {
		score := 1.0
		for _, elem := range words {
			score *= probFunc(elem)
		}
		if min < score {
			min = score
			chunk = words
		}
	}
	return chunk
}

// prev stores previously found segmentations.
var prev = map[string][]string{}

// Get returns the highest-scoring segmentation.
func Get(txt string, probFunc func(string) float64) []string {
	if len(txt) == 0 {
		return []string{}
	}
	if result, ok := prev[txt]; ok {
		return result
	}
	chunks := [][]string{}
	for _, split := range doSplit(txt) {
		chunks = append(chunks,
			append([]string{split.Head},
				Get(split.Tail, probFunc)...,
			),
		)
	}
	rs := mostPlausible(chunks, probFunc)
	prev[txt] = rs
	return rs
}
