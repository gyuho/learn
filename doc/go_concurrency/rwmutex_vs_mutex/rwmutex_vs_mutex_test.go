package rwmutex_vs_mutex

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"testing"
)

func TestInterface(t *testing.T) {
	func() {
		d := newDataRWMutex()
		ix := rand.Perm(testSize)
		for _, v := range testValues {
			d.set(v)
		}
		for _, idx := range ix {
			d.delete(testValues[idx])
		}
		for _, idx := range ix {
			if d.exist(testValues[idx]) {
				t.Errorf("%s should have not existed!", testValues[idx])
			}
		}
	}()

	func() {
		d := newDataMutex()
		ix := rand.Perm(testSize)
		for _, v := range testValues {
			d.set(v)
		}
		for _, idx := range ix {
			d.delete(testValues[idx])
		}
		for _, idx := range ix {
			if d.exist(testValues[idx]) {
				t.Errorf("%s should have not existed!", testValues[idx])
			}
		}
	}()
}

var (
	opt        string
	testSize   = 3000000
	testValues = []string{}
)

func init() {
	flag.StringVar(&opt, "opt", "rwmutex", "'rwmutex' or 'mutex'.")
	flag.Parse()
	opt = strings.TrimSpace(strings.ToLower(opt))
	if opt != "rwmutex" && opt != "mutex" {
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
	if opt == "rwmutex" {
		d = newDataRWMutex()
	} else {
		d = newDataMutex()
	}

	b.StartTimer()
	b.ReportAllocs()

	for _, v := range testValues {
		d.set(v)
	}
}

func BenchmarkExist(b *testing.B) {
	b.StopTimer()
	var d Interface
	if opt == "rwmutex" {
		d = newDataRWMutex()
	} else {
		d = newDataMutex()
	}
	for _, v := range testValues {
		d.set(v)
	}

	// to make it not biased towards data structures
	// with an order, such as slice.
	ix := rand.Perm(testSize)

	b.StartTimer()
	b.ReportAllocs()

	for _, idx := range ix {
		if !d.exist(testValues[idx]) {
			b.Errorf("%s should have existed!", testValues[idx])
		}
	}
}

func BenchmarkDelete(b *testing.B) {
	b.StopTimer()
	var d Interface
	if opt == "rwmutex" {
		d = newDataRWMutex()
	} else {
		d = newDataMutex()
	}
	for _, v := range testValues {
		d.set(v)
	}

	// to make it not biased towards data structures
	// with an order, such as slice.
	ix := rand.Perm(testSize)

	b.StartTimer()
	b.ReportAllocs()

	for _, idx := range ix {
		d.delete(testValues[idx])
	}
}
