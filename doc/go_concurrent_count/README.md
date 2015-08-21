[*back to contents*](https://github.com/gyuho/learn#contents)
<br>

# Go: concurrent count

- [Reference](#reference)
- [Counting Problem](#counting-problem)
- [Simulate web requests](#simulate-web-requests)
- [Count: `NaiveCounter`](#count-naivecounter)
- [Count: `MutexCounter`](#count-mutexcounter)
- [Count: `RWMutexCounter`](#count-rwmutexcounter)
- [Count: `AtomicIntCounter`](#count-atomicintcounter)
- [Count: `AtomicCounter`](#count-atomiccounter)
- [Count: `ChannelCounter` (No Buffer)](#count-channelcounter-no-buffer)
- [Count: `ChannelCounter` (Buffer)](#count-channelcounter-buffer)
- [Benchmark Results](#benchmark-results)

[↑ top](#go-concurrent-count)
<br><br><br><br>
<hr>








#### Reference

- [Bjorn Rabenstein - Prometheus: Designing and Implementing a Modern Monitoring Solution in Go](https://www.youtube.com/watch?v=1V7eJ0jN8-E)
- [beorn7/concurrentcount](https://github.com/beorn7/concurrentcount)
- [gyuho/count](https://github.com/gyuho/count)

[↑ top](#go-concurrent-count)
<br><br><br><br>
<hr>








#### Counting Problem

Suppose millions of **concurrent** web requests coming to your web application. 
And you want to *count* visits, or any other metrics per request. Counting should
not hurt the performance of your application. **Counting** is an **inherently 
sequential** problem. There's one resource to be updated while **concurrent**, 
multiple requests can cause contentions. Then what would be the best way to **count**
with concurrency?

[↑ top](#go-concurrent-count)
<br><br><br><br>
<hr>








#### Simulate web requests

Here's how I would simulate the web requests:

```go
func RunCountHandler(b *testing.B, isDebug bool, counter Counter, delta float64) {
	countHandler := func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			original := counter.Get()
			counter.Add(delta)
			fmt.Fprintf(w, "Original: %v / Added: %v / Current: %v", original, delta, counter.Get())
		default:
			http.Error(w, "Method Not Allowed", 405)
		}
	}

	mainRouter := http.NewServeMux()
	mainRouter.HandleFunc("/", httpLog(isDebug, countHandler))

	numberOfRequests := 100
	// don't do this at Travis

	for i := 0; i < b.N; i++ {

		b.StopTimer()
		ts := httptest.NewServer(mainRouter)

		var wg sync.WaitGroup
		wg.Add(numberOfRequests)

		for i := 0; i < numberOfRequests; i++ {
			go func() {
				defer wg.Done()
				if resp, err := http.Get(ts.URL); err != nil {
					panic(err)
				} else {
					// bstr, err := ioutil.ReadAll(resp.Body)
					resp.Body.Close()
					// if err != nil {
					// 	panic(err)
					// }
					// fmt.Println(string(bstr))

					// without Close
					// 2015/08/02 16:49:00 http: Accept error: accept tcp6 [::1]:38096: accept4: too many open files; retrying in 1s
					// 2015/08/02 16:49:01 http: Accept error: accept tcp6 [::1]:38096: accept4: too many open files; retrying in 1s
				}

			}()
		}

		b.StartTimer()
		wg.Wait()
		ts.Close()
	}
}
```

<br>
Counting operation takes only about nanoseconds while `http` request
takes much milliseconds. Benchmarking by mocking web server won't be able to isolate
the performance of `counting` as below, except that `channel` method is slower because
it allocates more memory:

```
BenchmarkServer_NaiveCounter              	     100	  13641425 ns/op	 1724319 B/op	   10551 allocs/op
BenchmarkServer_NaiveCounter-2            	     200	   5577538 ns/op	 1761024 B/op	   10465 allocs/op
BenchmarkServer_NaiveCounter-4            	     300	   3970441 ns/op	 1736143 B/op	   10392 allocs/op
BenchmarkServer_NaiveCounter-8            	     500	   3054495 ns/op	 1636052 B/op	    9846 allocs/op
BenchmarkServer_NaiveCounter-16           	     500	   2754022 ns/op	 1446608 B/op	    8784 allocs/op

BenchmarkServer_MutexCounter              	     100	  10334728 ns/op	 1739715 B/op	   10570 allocs/op
BenchmarkServer_MutexCounter-2            	     200	   6533533 ns/op	 1737853 B/op	   10466 allocs/op
BenchmarkServer_MutexCounter-4            	     300	   4217715 ns/op	 1703817 B/op	   10349 allocs/op
BenchmarkServer_MutexCounter-8            	     500	   3072379 ns/op	 1599124 B/op	    9745 allocs/op
BenchmarkServer_MutexCounter-16           	     500	   2721123 ns/op	 1417956 B/op	    8579 allocs/op

BenchmarkServer_RWMutexCounter            	     100	  11248896 ns/op	 1736902 B/op	   10579 allocs/op
BenchmarkServer_RWMutexCounter-2          	     200	   7160659 ns/op	 1759653 B/op	   10481 allocs/op
BenchmarkServer_RWMutexCounter-4          	     300	   4439413 ns/op	 1718228 B/op	   10390 allocs/op
BenchmarkServer_RWMutexCounter-8          	     500	   3340555 ns/op	 1679569 B/op	   10077 allocs/op
BenchmarkServer_RWMutexCounter-16         	     500	   3053389 ns/op	 1438662 B/op	    8698 allocs/op

BenchmarkServer_AtomicIntCounter          	     100	  12053604 ns/op	 1743955 B/op	   10590 allocs/op
BenchmarkServer_AtomicIntCounter-2        	     200	   8204060 ns/op	 1750468 B/op	   10477 allocs/op
BenchmarkServer_AtomicIntCounter-4        	     300	   4443112 ns/op	 1710413 B/op	   10370 allocs/op
BenchmarkServer_AtomicIntCounter-8        	     500	   3961467 ns/op	 1630977 B/op	    9897 allocs/op
BenchmarkServer_AtomicIntCounter-16       	     500	   2926347 ns/op	 1441098 B/op	    8780 allocs/op

BenchmarkServer_AtomicCounter             	     100	  11159504 ns/op	 1736091 B/op	   10570 allocs/op
BenchmarkServer_AtomicCounter-2           	     200	   7661146 ns/op	 1741652 B/op	   10482 allocs/op
BenchmarkServer_AtomicCounter-4           	     300	   4450239 ns/op	 1725751 B/op	   10406 allocs/op
BenchmarkServer_AtomicCounter-8           	     500	   3121161 ns/op	 1627260 B/op	    9925 allocs/op
BenchmarkServer_AtomicCounter-16          	     500	   2963900 ns/op	 1465410 B/op	    8873 allocs/op

BenchmarkServer_ChannelCounter_NoBuffer   	     100	 113879946 ns/op	 1801659 B/op	   10602 allocs/op
BenchmarkServer_ChannelCounter_NoBuffer-2 	      20	 111064393 ns/op	 1742514 B/op	   10512 allocs/op
BenchmarkServer_ChannelCounter_NoBuffer-4 	      20	 110180521 ns/op	 1801574 B/op	   10566 allocs/op
BenchmarkServer_ChannelCounter_NoBuffer-8 	     100	  30717707 ns/op	 1990469 B/op	   10692 allocs/op
BenchmarkServer_ChannelCounter_NoBuffer-16	     100	  24029631 ns/op	 1689640 B/op	    9902 allocs/op

BenchmarkServer_ChannelCounter_Buffer     	       2	1126995870 ns/op	 1576520 B/op	   10680 allocs/op
BenchmarkServer_ChannelCounter_Buffer-2   	       3	 684545001 ns/op	 1710218 B/op	   10609 allocs/op
BenchmarkServer_ChannelCounter_Buffer-4   	       3	 417227794 ns/op	 1782202 B/op	   10636 allocs/op
BenchmarkServer_ChannelCounter_Buffer-8   	      10	 188985058 ns/op	 1850097 B/op	   10654 allocs/op
BenchmarkServer_ChannelCounter_Buffer-16  	      10	 119519447 ns/op	 1591680 B/op	    9212 allocs/op
```


[↑ top](#go-concurrent-count)
<br><br><br><br>
<hr>









#### Count: `NaiveCounter`

[**_`NaiveCounter`_**](https://godoc.org/github.com/gyuho/count#NaiveCounter) is the
fastest way to count but subject to race conditions, as [here](http://play.golang.org/p/wBW-vMCSLl). 
This is not thread-safe:

```go
package main

import (
	"fmt"
	"sync"
)

// Counter is an interface for counting.
// It contains counting data as long as a type
// implements all the methods in the interface.
type Counter interface {
	// Get returns the current count.
	Get() float64

	// Add adds the delta value to the counter.
	Add(delta float64)
}

// NaiveCounter counts in a naive way.
// Do not use this with concurrency.
// It will cause race conditions.
type NaiveCounter float64

func (c *NaiveCounter) Get() float64 {

	// return (*c).(float64)
	// (X) (*c).(float64) (non-interface type NaiveCounter on left)

	return float64(*c)
}

func (c *NaiveCounter) Add(delta float64) {
	*c += NaiveCounter(delta)
}

func main() {
	counter := new(NaiveCounter)
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			counter.Add(1.347)
			counter.Get()
			counter.Add(-5.5)
			counter.Get()
			counter.Add(0.340)
			counter.Get()
		}()
	}
	wg.Wait()

	fmt.Println(counter.Get())
	// -38.12999999999999
}

/*
go run -race 00_naive.go

==================
WARNING: DATA RACE
Read by goroutine 7:
  main.main.func1()
      /home/ubuntu/go/src/github.com/gyuho/learn/doc/go_concurrent_count/code/00_naive.go:43 +0x70

Previous write by goroutine 6:
  main.main.func1()
      /home/ubuntu/go/src/github.com/gyuho/learn/doc/go_concurrent_count/code/00_naive.go:43 +0x88

Goroutine 7 (running) created at:
  main.main()
      /home/ubuntu/go/src/github.com/gyuho/learn/doc/go_concurrent_count/code/00_naive.go:49 +0xc5

Goroutine 6 (finished) created at:
  main.main()
      /home/ubuntu/go/src/github.com/gyuho/learn/doc/go_concurrent_count/code/00_naive.go:49 +0xc5
==================
-38.12999999999999
Found 1 data race(s)
exit status 66

*/

```


[↑ top](#go-concurrent-count)
<br><br><br><br>
<hr>









#### Count: `MutexCounter`

Try [this](http://play.golang.org/p/gxD7rxQ1b7):

```go
package main

import (
	"fmt"
	"sync"
)

// Counter is an interface for counting.
// It contains counting data as long as a type
// implements all the methods in the interface.
type Counter interface {
	// Get returns the current count.
	Get() float64

	// Add adds the delta value to the counter.
	Add(delta float64)
}

// MutexCounter implements Counter with sync.Mutex.
type MutexCounter struct {
	mu    sync.Mutex // guards the following
	value float64
}

func (c *MutexCounter) Get() float64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func (c *MutexCounter) Add(delta float64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value += delta
}

func main() {
	counter := new(MutexCounter)
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			counter.Add(1.347)
			counter.Get()
			counter.Add(-5.5)
			counter.Get()
			counter.Add(0.340)
			counter.Get()
		}()
	}
	wg.Wait()

	fmt.Println(counter.Get())
	// 962.0000000000002
}

```

[↑ top](#go-concurrent-count)
<br><br><br><br>
<hr>











#### Count: `RWMutexCounter`

Try [this](http://play.golang.org/p/1cWPcFVvPA):

```go
package main

import (
	"fmt"
	"sync"
)

// Counter is an interface for counting.
// It contains counting data as long as a type
// implements all the methods in the interface.
type Counter interface {
	// Get returns the current count.
	Get() float64

	// Add adds the delta value to the counter.
	Add(delta float64)
}

// RWMutexCounter implements Counter with sync.RWMutex.
type RWMutexCounter struct {
	mu    sync.RWMutex // guards the following sync.
	value float64
}

func (c *RWMutexCounter) Get() float64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.value
}

func (c *RWMutexCounter) Add(delta float64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value += delta
}

func main() {
	counter := new(RWMutexCounter)
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			counter.Add(1.347)
			counter.Get()
			counter.Add(-5.5)
			counter.Get()
			counter.Add(0.340)
			counter.Get()
		}()
	}
	wg.Wait()

	fmt.Println(counter.Get())
	// -38.12999999999999
}

```

[↑ top](#go-concurrent-count)
<br><br><br><br>
<hr>











#### Count: `AtomicIntCounter`

Try [this](http://play.golang.org/p/HlGY7UnZDg):

```go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Counter is an interface for counting.
// It contains counting data as long as a type
// implements all the methods in the interface.
type Counter interface {
	// Get returns the current count.
	Get() float64

	// Add adds the delta value to the counter.
	Add(delta float64)
}

// AtomicIntCounter implements Counter with atomic package.
// Go has only int64 atomic variable.
// This truncates float value into integer.
type AtomicIntCounter int64

func (c *AtomicIntCounter) Get() float64 {
	return float64(atomic.LoadInt64((*int64)(c)))
}

// Add ignores the non-integer part of delta.
func (c *AtomicIntCounter) Add(delta float64) {
	atomic.AddInt64((*int64)(c), int64(delta))
}

func main() {
	counter := new(AtomicIntCounter)
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			counter.Add(1.347)
			counter.Get()
			counter.Add(-5.5)
			counter.Get()
			counter.Add(0.340)
			counter.Get()
		}()
	}
	wg.Wait()

	fmt.Println(counter.Get())
	// -40
}

```

[↑ top](#go-concurrent-count)
<br><br><br><br>
<hr>










#### Count: `AtomicCounter`

Try [this](http://play.golang.org/p/cbMs7C-zSH):

```go
package main

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
)

// Counter is an interface for counting.
// It contains counting data as long as a type
// implements all the methods in the interface.
type Counter interface {
	// Get returns the current count.
	Get() float64

	// Add adds the delta value to the counter.
	Add(delta float64)
}

// AtomicCounter implements Counter with atomic package.
// Go has only int64 atomic variable.
// This uses math.Float64frombits package for the floating
// point number corresponding the IEEE 754 binary representation
type AtomicCounter uint64

func (c *AtomicCounter) Get() float64 {
	return math.Float64frombits(atomic.LoadUint64((*uint64)(c)))
}

// Add ignores the non-integer part of delta.
func (c *AtomicCounter) Add(delta float64) {
	for {
		oldBits := atomic.LoadUint64((*uint64)(c))
		newBits := math.Float64bits(math.Float64frombits(oldBits) + delta)
		if atomic.CompareAndSwapUint64((*uint64)(c), oldBits, newBits) {
			return
		}
	}
}

func main() {
	counter := new(AtomicCounter)
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			counter.Add(1.347)
			counter.Get()
			counter.Add(-5.5)
			counter.Get()
			counter.Add(0.340)
			counter.Get()
		}()
	}
	wg.Wait()

	fmt.Println(counter.Get())
	// -38.12999999999999
}

```

[↑ top](#go-concurrent-count)
<br><br><br><br>
<hr>








#### Count: `ChannelCounter` (No Buffer)

Try this:

```go
package main

import (
	"fmt"
	"sync"
)

// Counter is an interface for counting.
// It contains counting data as long as a type
// implements all the methods in the interface.
type Counter interface {
	// Get returns the current count.
	Get() float64

	// Add adds the delta value to the counter.
	Add(delta float64)
}

// ChannelCounter counts through channels.
type ChannelCounter struct {
	valueChan chan float64
	deltaChan chan float64
	done      chan struct{}
}

func NewChannelCounter(buf int) *ChannelCounter {
	c := &ChannelCounter{
		make(chan float64),
		make(chan float64, buf), // only buffer the deltaChan
		make(chan struct{}),
	}
	go c.Run()
	return c
}

func (c *ChannelCounter) Run() {

	var value float64

	for {
		// "select" statement chooses which of a set of
		// possible send or receive operations will proceed.
		select {

		case delta := <-c.deltaChan:
			value += delta

		case <-c.done:
			return

		case c.valueChan <- value:
			// Do nothing.

			// If there is no default case, the "select" statement
			// blocks until at least one of the communications can proceed.
		}
	}
}

func (c *ChannelCounter) Get() float64 {
	return <-c.valueChan
}

func (c *ChannelCounter) Add(delta float64) {
	c.deltaChan <- delta
}

func (c *ChannelCounter) Done() {
	c.done <- struct{}{}
}

func (c *ChannelCounter) Close() {
	close(c.deltaChan)
}

func main() {
	counter := NewChannelCounter(0)
	defer counter.Done()
	defer counter.Close()
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			counter.Add(1.347)
			counter.Get()
			counter.Add(-5.5)
			counter.Get()
			counter.Add(0.340)
			counter.Get()
		}()
	}
	wg.Wait()

	fmt.Println(counter.Get())
	// -38.12999999999997
}

```

[↑ top](#go-concurrent-count)
<br><br><br><br>
<hr>









#### Count: `ChannelCounter` (Buffer)

Try this:

```go
package main

import (
	"fmt"
	"sync"
)

// Counter is an interface for counting.
// It contains counting data as long as a type
// implements all the methods in the interface.
type Counter interface {
	// Get returns the current count.
	Get() float64

	// Add adds the delta value to the counter.
	Add(delta float64)
}

// ChannelCounter counts through channels.
type ChannelCounter struct {
	valueChan chan float64
	deltaChan chan float64
	done      chan struct{}
}

func NewChannelCounter(buf int) *ChannelCounter {
	c := &ChannelCounter{
		make(chan float64),
		make(chan float64, buf), // only buffer the deltaChan
		make(chan struct{}),
	}
	go c.Run()
	return c
}

func (c *ChannelCounter) Run() {

	var value float64

	for {
		// "select" statement chooses which of a set of
		// possible send or receive operations will proceed.
		select {

		case delta := <-c.deltaChan:
			value += delta

		case <-c.done:
			return

		case c.valueChan <- value:
			// Do nothing.

			// If there is no default case, the "select" statement
			// blocks until at least one of the communications can proceed.
		}
	}
}

func (c *ChannelCounter) Get() float64 {
	return <-c.valueChan
}

func (c *ChannelCounter) Add(delta float64) {
	c.deltaChan <- delta
}

func (c *ChannelCounter) Done() {
	c.done <- struct{}{}
}

func (c *ChannelCounter) Close() {
	close(c.deltaChan)
}

func main() {
	counter := NewChannelCounter(10)
	defer counter.Done()
	defer counter.Close()
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			counter.Add(1.347)
			counter.Get()
			counter.Add(-5.5)
			counter.Get()
			counter.Add(0.340)
			counter.Get()
		}()
	}
	wg.Wait()

	fmt.Println(counter.Get())
	// -38.12999999999997
}

```


[↑ top](#go-concurrent-count)
<br><br><br><br>
<hr>








#### Benchmark Results

For the full results, please take a look [gyuho/count](https://github.com/gyuho/count/blob/master/benchmark_results.txt).

`Add`, in the descending order of time per operation:

1. [**_`NaiveCounter`_**](https://godoc.org/github.com/gyuho/count#NaiveCounter) but should be ignored. Not thread-safe
2. [**_`AtomicIntCounter`_**](https://godoc.org/github.com/gyuho/count#AtomicIntCounter) but only supports `int64` type
3. [**_`AtomicCounter`_**](https://godoc.org/github.com/gyuho/count#AtomicCounter)
4. [**_`MutexCounter`_**](https://godoc.org/github.com/gyuho/count#MutexCounter)
5. [**_`RWMutexCounter`_**](https://godoc.org/github.com/gyuho/count#RWMutexCounter)
6. [**_`ChannelCounter_Buffer`_**](https://godoc.org/github.com/gyuho/count#ChannelCounter) is faster than non-buffered channel
7. [**_`ChannelCounter_NoBuffer`_**](https://godoc.org/github.com/gyuho/count#ChannelCounter)


<br>
`Get`, in the descending order of time per operation:

1. [**_`NaiveCounter`_**](https://godoc.org/github.com/gyuho/count#NaiveCounter) but should be ignored. Not thread-safe
2. [**_`AtomicIntCounter`_**](https://godoc.org/github.com/gyuho/count#AtomicIntCounter) but only supports `int64` type
3. [**_`AtomicCounter`_**](https://godoc.org/github.com/gyuho/count#AtomicCounter)
4. [**_`MutexCounter`_**](https://godoc.org/github.com/gyuho/count#MutexCounter)
5. [**_`RWMutexCounter`_**](https://godoc.org/github.com/gyuho/count#RWMutexCounter)
6. [**_`ChannelCounter_Buffer`_**](https://godoc.org/github.com/gyuho/count#ChannelCounter) is faster than non-buffered channel
7. [**_`ChannelCounter_NoBuffer`_**](https://godoc.org/github.com/gyuho/count#ChannelCounter)

<br>

And `channel` is slower than `sync.Mutex` because it allocates more memory.

[↑ top](#go-concurrent-count)
<br><br><br><br>
<hr>
