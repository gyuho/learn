package count

import (
	"sync"
	"testing"
)

var (
	isDebug              = false
	delta                = 0.5
	testNumberOfRequests = 1000
)

// func BenchmarkServer_NaiveCounter(b *testing.B) {
// 	counter := new(NaiveCounter)
// 	RunCountHandler(b, isDebug, counter, delta)
// }

// func BenchmarkServer_MutexCounter(b *testing.B) {
// 	counter := new(MutexCounter)
// 	RunCountHandler(b, isDebug, counter, delta)
// }

// func BenchmarkServer_RWMutexCounter(b *testing.B) {
// 	counter := new(RWMutexCounter)
// 	RunCountHandler(b, isDebug, counter, delta)
// }

// func BenchmarkServer_AtomicIntCounter(b *testing.B) {
// 	counter := new(AtomicIntCounter)
// 	RunCountHandler(b, isDebug, counter, delta)
// }

// func BenchmarkServer_AtomicCounter(b *testing.B) {
// 	counter := new(AtomicCounter)
// 	RunCountHandler(b, isDebug, counter, delta)
// }

// func BenchmarkServer_ChannelCounter_NoBuffer(b *testing.B) {
// 	counter := NewChannelCounter(0)
// 	defer counter.Close()
// 	RunCountHandler(b, isDebug, counter, delta)
// }

// func BenchmarkServer_ChannelCounter_Buffer(b *testing.B) {
// 	counter := NewChannelCounter(testNumberOfRequests)
// 	defer counter.Close()
// 	RunCountHandler(b, isDebug, counter, delta)
// }

func runAdd(b *testing.B, counter Counter, delta float64, numberOfRequests int) {
	b.StopTimer()
	var wg sync.WaitGroup
	wg.Add(numberOfRequests)
	for i := 0; i < numberOfRequests; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < b.N; i++ {
				counter.Add(delta)
			}
		}()
	}
	b.StartTimer()
	wg.Wait()
}

func runGet(b *testing.B, counter Counter, delta float64, numberOfRequests int) {
	b.StopTimer()
	for i := 0; i < numberOfRequests; i++ {
		counter.Add(delta)
	}
	var wg sync.WaitGroup
	wg.Add(numberOfRequests)
	for i := 0; i < numberOfRequests; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < b.N; i++ {
				counter.Get()
			}
		}()
	}
	b.StartTimer()
	wg.Wait()
}

func BenchmarkAdd_NaiveCounter(b *testing.B) {
	counter := new(NaiveCounter)
	runAdd(b, counter, delta, testNumberOfRequests)
}

func BenchmarkAdd_MutexCounter(b *testing.B) {
	counter := new(MutexCounter)
	runAdd(b, counter, delta, testNumberOfRequests)
}

func BenchmarkAdd_RWMutexCounter(b *testing.B) {
	counter := new(RWMutexCounter)
	runAdd(b, counter, delta, testNumberOfRequests)
}

func BenchmarkAdd_AtomicIntCounter(b *testing.B) {
	counter := new(AtomicIntCounter)
	runAdd(b, counter, delta, testNumberOfRequests)
}

func BenchmarkAdd_AtomicCounter(b *testing.B) {
	counter := new(AtomicCounter)
	runAdd(b, counter, delta, testNumberOfRequests)
}

func BenchmarkAdd_ChannelCounter_NoBuffer(b *testing.B) {
	counter := NewChannelCounter(0)
	defer counter.Close()
	runAdd(b, counter, delta, testNumberOfRequests)
}

func BenchmarkAdd_ChannelCounter_Buffer(b *testing.B) {
	counter := NewChannelCounter(testNumberOfRequests)
	defer counter.Done()
	defer counter.Close()
	runAdd(b, counter, delta, testNumberOfRequests)
}

func BenchmarkGet_NaiveCounter(b *testing.B) {
	counter := new(NaiveCounter)
	runGet(b, counter, delta, testNumberOfRequests)
}

func BenchmarkGet_MutexCounter(b *testing.B) {
	counter := new(MutexCounter)
	runGet(b, counter, delta, testNumberOfRequests)
}

func BenchmarkGet_RWMutexCounter(b *testing.B) {
	counter := new(RWMutexCounter)
	runGet(b, counter, delta, testNumberOfRequests)
}

func BenchmarkGet_AtomicIntCounter(b *testing.B) {
	counter := new(AtomicIntCounter)
	runGet(b, counter, delta, testNumberOfRequests)
}

func BenchmarkGet_AtomicCounter(b *testing.B) {
	counter := new(AtomicCounter)
	runGet(b, counter, delta, testNumberOfRequests)
}

func BenchmarkGet_ChannelCounter_NoBuffer(b *testing.B) {
	counter := NewChannelCounter(0)
	defer counter.Done()
	defer counter.Close()
	runGet(b, counter, delta, testNumberOfRequests)
}

func BenchmarkGet_ChannelCounter_Buffer(b *testing.B) {
	counter := NewChannelCounter(testNumberOfRequests)
	defer counter.Done()
	defer counter.Close()
	runGet(b, counter, delta, testNumberOfRequests)
}
