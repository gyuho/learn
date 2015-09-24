package slice_vs_map

import (
	"flag"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

const (
	// http://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func randBytes(n int) []byte {
	src := rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return b
}

var (
	testMessages = []*message{}
	mA           = newMessage("A")
	mB           = newMessage("B")
	mC           = newMessage("C")

	opt string
)

func init() {
	flag.StringVar(&opt, "opt", "Slice", "Slice or Map.")
	flag.Parse()
	fmt.Println()
	fmt.Println("testing with", opt)
	fmt.Println()

	for i := 0; i < 10000; i++ {
		rand.Seed(time.Now().UnixNano())
		m := newMessage(string(randBytes(10)))
		testMessages = append(testMessages, m)
	}
	testMessages = append(testMessages, mA)
	testMessages = append(testMessages, mB)
	testMessages = append(testMessages, mC)
}

func TestSlice(t *testing.T) {
	s := newSlice()
	for _, msg := range testMessages {
		s.Add(msg)
	}
	if !s.Exist(mA) {
		t.Fatalf("%s must be there", mA)
	}
	if !s.Exist(mB) {
		t.Fatalf("%s must be there", mB)
	}
	if !s.Exist(mC) {
		t.Fatalf("%s must be there", mC)
	}
}

func TestMap(t *testing.T) {
	m := newMap()
	for _, msg := range testMessages {
		m.Add(msg)
	}
	if !m.Exist(mA) {
		t.Fatalf("%s must be there", mA)
	}
	if !m.Exist(mB) {
		t.Fatalf("%s must be there", mB)
	}
	if !m.Exist(mC) {
		t.Fatalf("%s must be there", mC)
	}
}

func BenchmarkAdd(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	var d Interface
	if opt == "Slice" {
		d = newSlice()
	} else {
		d = newMap()
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for _, msg := range testMessages {
			d.Add(msg)
		}
	}
}

func BenchmarkDelete(b *testing.B) {
	b.ReportAllocs()
	var d Interface
	if opt == "Slice" {
		d = newSlice()
	} else {
		d = newMap()
	}
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for _, msg := range testMessages {
			d.Add(msg)
		}
		b.StartTimer()
		for _, msg := range testMessages {
			d.Delete(msg)
		}
	}
}
