package map_pointer_vs_map_int

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
		d := newMapPointer()
		ix := rand.Perm(testN)
		for _, v := range testVals {
			d.set(v)
		}
		for _, idx := range ix {
			d.delete(testVals[idx])
		}
		for _, idx := range ix {
			if d.exist(testVals[idx]) {
				t.Errorf("%s should have not existed!", testVals[idx])
			}
		}
	}()

	func() {
		d := newMapInt()
		ix := rand.Perm(testN)
		for _, v := range testVals {
			d.set(v)
		}
		for _, idx := range ix {
			d.delete(testVals[idx])
		}
		for _, idx := range ix {
			if d.exist(testVals[idx]) {
				t.Errorf("%s should have not existed!", testVals[idx])
			}
		}
	}()
}

var (
	opt      string
	testN    = 3000000
	testVals = make([]*metric, testN)
)

func init() {
	flag.StringVar(&opt, "opt", "mappointer", "'mappointer' or 'mapint'.")
	flag.Parse()
	opt = strings.TrimSpace(strings.ToLower(opt))
	if opt != "mappointer" && opt != "mapint" {
		fmt.Fprintln(os.Stderr, fmt.Errorf("unknown option", opt))
		os.Exit(1)
	}
	log.Println("Running benchmarks with", opt)

	log.Println("Filling up the test data...")
	vs := multiRandBytes(15, testN)
	for i := 0; i < testN; i++ {
		m := metric{}
		m.id = i
		m.s = string(vs[i])
		testVals[i] = &m
	}
	log.Println("Done! Test data is ready!")
}

func BenchmarkSet(b *testing.B) {
	var d Interface
	if opt == "mappointer" {
		d = newMapPointer()
	} else {
		d = newMapInt()
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
	if opt == "mappointer" {
		d = newMapPointer()
	} else {
		d = newMapInt()
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
		if !d.exist(testVals[idx]) {
			b.Errorf("%s should have existed!", testVals[idx])
		}
	}
}

func BenchmarkDelete(b *testing.B) {
	b.StopTimer()
	var d Interface
	if opt == "mappointer" {
		d = newMapPointer()
	} else {
		d = newMapInt()
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
		d.delete(testVals[idx])
	}
}
