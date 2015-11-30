package map_to_slice_vs_map

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"testing"
)

func TestMapTo(t *testing.T) {
	var data Interface
	key, val := "A", "HELLO"

	data = newMapToSlice()
	data.set(key, val)
	if !data.exist(key, val) {
		t.Errorf("%s should have existed but got %+v", val, data)
	}
	data.delete(key, val)
	if data.exist(key, val) {
		t.Errorf("%s should have not existed but got %+v", val, data)
	}
	data.set(key, val+val)
	data.set(key, val+val+val)
	if !data.exist(key, val+val) {
		t.Errorf("%s should have existed but got %+v", val+val, data)
	}
	data.delete(key, val+val)
	if data.exist(key, val+val) {
		t.Errorf("%s should have not existed but got %+v", val+val, data)
	}

	data = newMapToMap()
	data.set(key, val)
	if !data.exist(key, val) {
		t.Errorf("%s should have existed but got %+v", val, data)
	}
	data.delete(key, val)
	if data.exist(key, val) {
		t.Errorf("%s should have not existed but got %+v", val, data)
	}
	data.set(key, val+val)
	data.set(key, val+val+val)
	if !data.exist(key, val+val) {
		t.Errorf("%s should have existed but got %+v", val+val, data)
	}
	data.delete(key, val+val)
	if data.exist(key, val+val) {
		t.Errorf("%s should have not existed but got %+v", val+val, data)
	}

	func() {
		d := newMapToSlice()
		ix := rand.Perm(testSize)
		k := "A"
		for _, v := range testValues {
			d.set(k, v)
		}
		for _, idx := range ix {
			d.delete(k, testValues[idx])
		}
		for _, idx := range ix {
			if d.exist(k, testValues[idx]) {
				t.Errorf("%s should have not existed!", testValues[idx])
			}
		}
	}()

	func() {
		d := newMapToMap()
		ix := rand.Perm(testSize)
		k := "A"
		for _, v := range testValues {
			d.set(k, v)
		}
		for _, idx := range ix {
			d.delete(k, testValues[idx])
		}
		for _, idx := range ix {
			if d.exist(k, testValues[idx]) {
				t.Errorf("%s should have not existed!", testValues[idx])
			}
		}
	}()
}

var (
	opt        string
	testSize   = 30000
	testValues = []string{}
)

func init() {
	flag.StringVar(&opt, "opt", "slice", "'slice' or 'map'.")
	flag.Parse()
	opt = strings.TrimSpace(strings.ToLower(opt))
	if opt != "slice" && opt != "map" {
		fmt.Fprintln(os.Stderr, fmt.Errorf("unknown option", opt))
		os.Exit(1)
	}
	log.Println("Running benchmarks with", opt)

	log.Println("Filling up the test data...")
	for i := 0; i < testSize; i++ {
		testValues = append(testValues, string(randBytes(15)))
	}
	log.Println("Done! Test data is ready!")
}

func BenchmarkSet(b *testing.B) {
	var d Interface
	if opt == "slice" {
		d = newMapToSlice()
	} else {
		d = newMapToMap()
	}

	b.StartTimer()
	b.ReportAllocs()

	k := "A"
	for _, v := range testValues {
		d.set(k, v)
	}
}

func BenchmarkExist(b *testing.B) {
	b.StopTimer()
	var d Interface
	if opt == "slice" {
		d = newMapToSlice()
	} else {
		d = newMapToMap()
	}
	k := "A"
	for _, v := range testValues {
		d.set(k, v)
	}

	// to make it not biased towards data structures
	// with an order, such as slice.
	ix := rand.Perm(testSize)

	b.StartTimer()
	b.ReportAllocs()

	for _, idx := range ix {
		if !d.exist(k, testValues[idx]) {
			b.Errorf("%s should have existed!", testValues[idx])
		}
	}
}

func BenchmarkDelete(b *testing.B) {
	b.StopTimer()
	var d Interface
	if opt == "slice" {
		d = newMapToSlice()
	} else {
		d = newMapToMap()
	}
	k := "A"
	for _, v := range testValues {
		d.set(k, v)
	}

	// to make it not biased towards data structures
	// with an order, such as slice.
	ix := rand.Perm(testSize)

	b.StartTimer()
	b.ReportAllocs()

	for _, idx := range ix {
		d.delete(k, testValues[idx])
	}
}
