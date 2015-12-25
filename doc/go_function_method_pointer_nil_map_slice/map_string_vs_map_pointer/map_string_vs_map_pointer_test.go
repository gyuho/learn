package map_string_vs_map_pointer

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"testing"
)

func TestMap(t *testing.T) {
	func() {
		d := newMapString()
		ix := rand.Perm(testN)
		for _, v := range testVals {
			d.set(v)
		}
		for _, idx := range ix {
			d.delete(testVals[idx].id)
		}
		for _, idx := range ix {
			if d.exist(testVals[idx].id) {
				t.Errorf("%s should have not existed!", testVals[idx].id)
			}
		}
	}()

	func() {
		d := newMapPointer()
		ix := rand.Perm(testN)
		for _, v := range testVals {
			d.set(v)
		}
		for _, idx := range ix {
			d.delete(testVals[idx].id)
		}
		for _, idx := range ix {
			if d.exist(testVals[idx].id) {
				t.Errorf("%s should have not existed!", testVals[idx].id)
			}
		}
	}()
}

var (
	opt      string
	testN    = 10000
	testVals = make([]*node, testN)
)

func init() {
	flag.StringVar(&opt, "opt", "mapstring", "'mapstring' or 'mappointer'.")
	flag.Parse()
	opt = strings.TrimSpace(strings.ToLower(opt))
	if opt != "mapstring" && opt != "mappointer" {
		fmt.Fprintln(os.Stderr, fmt.Errorf("unknown option %s", opt))
		os.Exit(1)
	}
	log.Println("Running benchmarks with", opt)

	log.Println("Filling up the test data...")
	vs := multiRandBytes(15, testN)
	for i := 0; i < testN; i++ {
		m := node{}
		m.id = string(vs[i])
		testVals[i] = &m
	}
	log.Println("Done! Test data is ready!")
}

func BenchmarkSet(b *testing.B) {
	var d Interface
	if opt == "mapstring" {
		d = newMapString()
	} else {
		d = newMapPointer()
	}

	b.StartTimer()
	b.ReportAllocs()

	for _, v := range testVals {
		d.set(v)
	}
}

func BenchmarkExist(b *testing.B) {
	b.StopTimer()
	var d Interface
	if opt == "mapstring" {
		d = newMapString()
	} else {
		d = newMapPointer()
	}
	for _, v := range testVals {
		d.set(v)
	}

	// to make it not biased towards data structures
	// with an order, such as slice.
	ix := rand.Perm(testN)

	b.StartTimer()
	b.ReportAllocs()

	for _, idx := range ix {
		if !d.exist(testVals[idx].id) {
			b.Errorf("%s should have existed!", testVals[idx])
		}
	}
}

func BenchmarkDelete(b *testing.B) {
	b.StopTimer()
	var d Interface
	if opt == "mapstring" {
		d = newMapString()
	} else {
		d = newMapPointer()
	}
	for _, v := range testVals {
		d.set(v)
	}

	// to make it not biased towards data structures
	// with an order, such as slice.
	ix := rand.Perm(testN)

	b.StartTimer()
	b.ReportAllocs()

	for _, idx := range ix {
		d.delete(testVals[idx].id)
	}
}
