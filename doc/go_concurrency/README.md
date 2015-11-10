[*back to contents*](https://github.com/gyuho/learn#contents)
<br>

# Go: concurrency

<br>

> If you look at the programming languages of today, you probably get this idea
> that the world is objected-oriented. But it’s not. It’s actually parallel.
> Multi-core machines, users, networking, etc. All these things are happening
> simultaneously and yet computing tools that we have are not good at
> expressing this kind of world-views. […] **Go is a concurrent language.**
>
> [**_Concurrency is not parallelism_**](https://www.youtube.com/watch?v=cN_DpYBzKso) *by Rob Pike*

<br>

- [Reference](#reference)
- [**Concurrency is not parallelism**](#concurrency-is-not-parallelism)
- [goroutine ≠ thread](#goroutine--thread)
- [`defer`, `recover`](#defer-recover)
- [**Be careful with `defer` and deadlock**](#be-careful-with-defer-and-deadlock)
- [channel to communicate](#channel-to-communicate)
- [questions](#questions)
	- [#1-1. synchronous, asynchronous channel](#1-1-synchronous-asynchronous-channel)
	- [#1-2. **buffered channel faster** because it’s *non-blocking*?](#1-2-buffered-channel-faster-because-its-non-blocking)
	- [#1-3. **be careful with buffered channel!**](#1-3-be-careful-with-buffered-channel)
	- [#1-4. **non-deterministic receive from buffered channel**](#1-4-non-deterministic-receive-from-buffered-channel)
	- [#2. why is this receiving only one value?](#2-why-is-this-receiving-only-one-value)
	- [#3. **wait for all goroutines to finish with `close`**](#3-wait-for-all-goroutines-to-finish-with-close)
- [**select for channel**: `select` ≠ `switch`](#select-for-channel-select--switch)
- [**receive `nil` from channel**](#receive-nil-from-channel)
- [`sync.Mutex`, race condition](#syncmutex-race-condition)
- [**Share memory by communicating**](#share-memory-by-communicating)
- [memory leak #1](#memory-leak-1)
- [`sync/atomic`](#syncatomic)
- [web server](#example-web-server)
- [`sync.Mutex` is just a value](#syncmutex-is-just-a-value)

[↑ top](#go-concurrency)
<br><br><br><br>
<hr>




#### Reference

- [*Thread*](https://en.wikipedia.org/wiki/Thread_(computing))
- [*Process*](https://en.wikipedia.org/wiki/Process_(computing))
- [*Asynchronous I/O*](https://en.wikipedia.org/wiki/Asynchronous_I/O)
- [*Context switch*](https://en.wikipedia.org/wiki/Context_switch)
- [*Green threads*](https://en.wikipedia.org/wiki/Green_threads)
- [**Concurrency is Not Parallelism**](https://www.youtube.com/watch?v=cN_DpYBzKso) *by Rob Pike* ([Slide](http://talks.golang.org/2012/waza.slide#1))
- [**Go Concurrency Patterns**](https://www.youtube.com/watch?v=f6kdp27TYZs) *by Rob Pike* ([Slide](https://talks.golang.org/2012/concurrency.slide#1))
- [**Effective Go — Concurrency**](https://golang.org/doc/effective_go.html#concurrency)
- [**Advanced Go Concurrency Patterns**](https://www.youtube.com/watch?v=QDDwwePbDtw) *by Sameer Ajmani* ([Slide](https://talks.golang.org/2013/advconc.slide#1))
- [**Go Concurrency Patterns : Pipelines and cancellation**](https://blog.golang.org/pipelines) *by Sameer Ajmani*
- [**Go Memory Model**](https://golang.org/ref/mem)
- [**Five things that make Go fast**](http://dave.cheney.net/2014/06/07/five-things-that-make-go-fast) *by Dave Cheney*
- [**Why is a goroutine’s stack infinite**](http://dave.cheney.net/2013/06/02/why-is-a-goroutines-stack-infinite) *by Dave Cheney*
- [**High performance servers without the event loop**](http://go-talks.appspot.com/github.com/davecheney/presentations/performance-without-the-event-loop.slide#1) *by Dave Cheney*
- [**Goroutines vs OS Threads**](https://groups.google.com/d/msg/golang-nuts/j51G7ieoKh4/wxNaKkFEfvcJ)

[↑ top](#go-concurrency)
<br><br><br><br>
<hr>




#### Concurrency is not parallelism


My YouTube video on Go concurrency:

<a href="https://www.youtube.com/watch?v=jsxshnyyTUY" target="_blank"><img src="http://img.youtube.com/vi/jsxshnyyTUY/0.jpg"></a>

You write any concurrent code, but you **run with a single processor, then your
program is** **_not parallel_** because it is not executing anything in
parallel. But Go code can still be concurrent with a single processor: when
there are multiple processors available, the code runs in parallel
automatically.

<br>

> Go is a concurrent language. Concurrency and parallelism are not the same
> thing. **Concurrency** is the **composition** of **independently** executing
> processes(computations). **Parallelism** is the **simultaneous** execution of
> (possibly related) computations. **Concurrency** *is about dealing with a lot
> of things at once.* **Parallelism** *is about doing a lot of things at once.*
> **Concurrency** is about programming structure. **Parallelism** is about
> **execution**. **Concurrency** provides a way to **structure** a solution to
> solve a problem that may (but *not necessariliy*) be parallelizable.
>
> [**_Rob Pike_**](https://www.youtube.com/watch?v=cN_DpYBzKso)


<br>
**_Independently Executing Procedure + Coordination = Concurrency_**
<br>

**Go’s concurrency is coordination, communication of independently executing
procedures.** Go concurrency model is like communication of UNIX pipelines: `ls
-l | grep key | less`. Go concurrency is more a [type-safe
generalization](https://golang.org/doc/effective_go.html#concurrency) of Unix
pipes. **_goroutine_** is like *ampersand* `&` in a shell command, which **runs
things in the background but does not wait for it to end**, as
[here](http://play.golang.org/p/5t84yLWCG9) and
[here](http://play.golang.org/p/rPP4s5ULSo):

```go
package main
 
import "fmt"
 
func main() {
	// launch goroutine in background
	go func() {
		fmt.Println("Hello, playground")
	}()
	//
	// Does not print anything
	//
	// when main returns
	// the program exits
	// and the goroutine will not be run
	// and gets garbage-collected
}
```
```go
package main
 
import (
	"fmt"
	"time"
)
 
func b() {
	fmt.Println("b is still running")
	fmt.Println("because although a exited but main hasn't exited yet!")
}
 
func a() {
	fmt.Println("a exits")
	go b()
}
 
func main() {
	a()
	time.Sleep(time.Second)
	// a exits
	// b is still running
	// because although a exited but main hasn't exited yet!
 
	go func() {
		fmt.Println("Hello, playground")
	}()
	time.Sleep(time.Second)
	// Hello, playground
}
```

[↑ top](#go-concurrency)
<br><br><br><br>
<hr>




#### goroutine ≠ thread

**_Thread_** is a lightweight process since it executes within the context of one
process. Both threads and processes are independent units of execution.
**Threads** under the **same process** **_run in one shared memory_** space,
while **process** **_run in separate memory_** spaces.

> Each *process* provides the resources needed to execute a program. A process
> has a virtual address space, executable code, open handles to system objects,
> a security context, a unique process identifier, environment variables, a
> priority class, minimum and maximum working set sizes, and at least one
> thread of execution. Each process is started with a single thread, often
> called the primary thread, but can create additional threads from any of its
> threads.
>
> A *thread* is the entity within a process that can be scheduled for
> execution. All threads of a process share its virtual address space and
> system resources. In addition, each thread maintains exception handlers, a
> scheduling priority, thread local storage, a unique thread identifier, and a
> set of structures the system will use to save the thread context until it is
> scheduled. The thread context includes the thread’s set of machine registers,
> the kernel stack, a thread environment block, and a user stack in the address
> space of the thread’s process. Threads can also have their own security
> context, which can be used for impersonating clients.
>
> [**_About Processes and Threads by
> Microsoft_**](https://msdn.microsoft.com/en-us/library/windows/desktop/ms681917%28v=vs.85%29.aspx)

<br>
When you say *8-core machine*, the `core` represents the actual physical
processors. *8-core machine* has 8 independent processing units (*cores* or
*CPU*s). Not to be confused with processor, a `process` is a computer program
instance that is being executed. A `process` can be made up of multiple
`threads` executing instructions concurrently. Again, `core` is an actual
physical `processor`, and `process` and `thread` are independent units of
program execution: `threads` under the same `process` run in a shared memory
space, whereas `processes` run in separate memory spaces. `threads` are more
dependent on an operating system, than a hardware or CPU. Normally one CPU can
handle one `thread` at a time, but one CPU with
[hyper threading](https://en.wikipedia.org/wiki/Hyper-threading)
can handle two `threads` simultaneously.

<br>
> [Threads] are conceptually the same as processes, but share the same memory space.
>
> As threads share address space, they are lighter than processes so are faster
> to create and faster to switch between.
>
> Threads still have an expensive context switch cost, a lot of state must be
> retained.
>
> Goroutines take the idea of threads a step further.
>
> Many goroutines are multiplexed onto a single operating system thread.
>	- Super cheap to create.
>	- Super cheap to switch between as it all happens in user space.
>	- Tens of thousands of goroutines in a single process are the norm,
>	    hundreds of thousands not unexpected.
>
> This results in relatively few operating system threads per Go process, with
> the Go runtime taking care of assigning a runnable Goroutine to a free
> operating system thread.
>
> [**High performance servers without the event
>  loop**](http://go-talks.appspot.com/github.com/davecheney/presentations/performance-without-the-event-loop.slide#1)
> *by Dave Cheney*

<br><br>

**goroutine** is an independently executing function, launched with go
statement. **goroutine** is **NOT a thread**. Think of goroutine as a **very
cheap, lightweight thread**. A program may have **thousands of goroutines**
*but with only one thread*.

> In telecommunications and computer networks, **multiplexing** (sometimes
> contracted to **muxing**) is a method by which multiple analog message
> signals or digital data streams are combined into one signal over a shared
> medium.
>
> [*Multiplexing*](https://en.wikipedia.org/wiki/Multiplexing) *by Wikipedia*

**Go runtime multiplexes goroutines into multiple OS threads**: when one
goroutine blocks such as waiting for I/O, the **thread blocks** too but
**no other goroutine blocks**. When a goroutine blocks on a thread,
Go runtime moves other goroutines to a different, available thread,
so they won't be blocked.

As of [Go 1.4](http://golang.org/doc/go1.4#runtime), the garbage collector has
become precise enough that **goroutine stack now takes only 2048 bytes of
memory**. goroutine has its own call stack that grows and shrinks as required.
It starts small and allocates, frees heap storage automatically. Go allows you
to write high-performance program without much expert knowledge or dealing
with OS threads.

<br>

> Each goroutine starts with a small stack, allocated from the heap. The size
> has fluctuated over time, but in Go 1.5 each goroutine starts with a 2k
> allocation.
>
> Instead of using guard pages, the Go compiler inserts a check as part of every
> function call to test if there is sufficient stack for the function to run.
>
> If there is insufficient space, the runtime will allocate a large stack segment
> on the heap, copy the contents of the current stack to the new segment, free
> the old segment, and the function call restarted.
>
> [*Goroutine stack
> growth*](http://go-talks.appspot.com/github.com/davecheney/presentations/performance-without-the-event-loop.slide#27)
> *by Dave Cheney*


<br>
Note that when the **_main_** function returns, the **program exists**.
goroutines that were running in background get **garbage-collected**, like
[here](http://play.golang.org/p/bODiFAAfTP):

```go
package main
 
import (
	"fmt"
	"time"
)
 
func a() {
	fmt.Println("a() called")
	go func() {
		time.Sleep(10 * time.Second)
		fmt.Println("go func() called")
		// this is not called
		//
		// you can get this printed with channel
	}()
	go b()
}
 
func b() {
	time.Sleep(1 * time.Second)
	fmt.Println("b() called")
}
 
func main() {
	a()
	time.Sleep(5 * time.Second)
	// when main returns all others return as well
}
 
/*
a() called
b() called
*/
```


[↑ top](#go-concurrency)
<br><br><br><br>
<hr>




#### `defer`, `recover`

**`defer`** **_delays the function execution_** **until just before the enclosing
function exit(return)**. The order of execution is:

- **`defer`**: Stack(*Last-In-First-Out*)
- **`goroutine`**: Queue(*First-In-First-Out*)

**LAST** *defer* statement runs **FIRST**, like
[here](http://play.golang.org/p/aDacv_4wn6):

```go
package main
 
import "fmt"
 
func main() {
	defer println("Defer 1")
	defer println("Defer 2")
	defer println("Defer 3")
 
	defer func() {
		fmt.Println("Recover:", recover())
	}()
	panic("Panic!!!")
 
	/*
		Recover: Panic!!!
		Defer 3
		Defer 2
		Defer 1
	*/
 
	// recover stops the panic
	// recover returns the value from panic
	// panic function is to cause a run time error
	// for "cannot happen" situations
	// And stops the program to begin panicking
	// So even if it's recovered
	// the next lines after panic won't be run.
	for {
		fmt.Println("This does not print! Anything below not being run!")
	}
}
```


**FIRST** *goroutine* runs **FIRST**, like
[here](http://play.golang.org/p/JrQzbVKvuR):

```go
package main
 
import "time"
 
func main() {
	// goroutine #01 : Queue
	go println(1)
 
	// goroutine #02
	// Anonymous Function Closure
	// Not function literal
	// So we need parenthesis at the end
	go func() {
		println(2)
	}()
 
	// goroutine #03
	// Anonymous Function Closure with input
	go func(n int) {
		println(n)
	}(3)
 
	// 1
	// 2
	// 3
 
	time.Sleep(time.Nanosecond)
	// main goroutine does not wait(block) for goroutine's return
	// Without this, we just reach the end of main and goroutine does not run
}
```

<br>

Note that **_defer_** still gets **executed_** even when a function `panic`s,
like [here](http://play.golang.org/p/JTnbSooYdK):

```go
package main
 
import (
	"fmt"
	"time"
)
 
func main() {
	go func() {
		defer fmt.Println("Hello, playground")
		panic(1)
	}()
 
	time.Sleep(time.Second)
}
 
/*
Hello, playground
panic: 1
*/
```


<br>

> `panic` is a built-in function that stops the ordinary flow of control and
> begins panicking. **_When the function F calls panic, execution of F stops,
> any deferred functions in F are executed normally, and then F returns to its
> caller._** To the caller, F then behaves like a call to panic. The process
> continues up the stack until all functions in the current goroutine have
> returned, at which point the program crashes. Panics can be initiated by
> invoking panic directly. They can also be caused by runtime errors, such as
> out-of-bounds array accesses.
>
> `recover` is a built-in function that regains control of a panicking
> goroutine. **_recover_** **is only useful inside defer**-red functions.
> During normal execution, a call to recover will return nil and have no other
> effect. If the current goroutine is panicking, a call to recover will capture
> the value given to panic and resume normal execution.
>
> [**_Andrew Gerrand_**](http://blog.golang.org/defer-panic-and-recover)

<br>

Note that when a function `panic`s, the **function execution stops** and it
runs *any* **defer statements inside the function**, and *it* **_returns_**.
So you won't see *"Hello World"*, from this
[code](http://play.golang.org/p/7abMxKTZDH):

```go
package main
 
import "fmt"
 
func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
 
	panic("Panic!")
 
	fmt.Println("Hello, World!")
	// NOT printed
}
 
/*
Panic!
*/
```

When it `panic`s, the **_main_** goroutine (`main` function) exits. That's why
we didn't see *"Hello World!"* in the code above.

<br><br>

Try this [code](http://play.golang.org/p/pGghekREGe) and
[code](http://play.golang.org/p/I76aHCgHON):

```go
package main
 
import (
	"fmt"
	"time"
)
 
func panicAndrecover() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	panic("Panic!")
}
 
func main() {
	panicAndrecover()
	fmt.Println("Hello, World!")
	/*
	   Panic!
	   Hello, World!
	*/
 
	recursiveRecover()
	/*
	   Restarting after error: [ 0 ] Panic
	   Restarting after error: [ 1 ] Panic
	   Restarting after error: [ 2 ] Panic
	   Restarting after error: [ 3 ] Panic
	   Restarting after error: [ 4 ] Panic
	   Too much panic: 5
	*/
}
 
var count int
 
func recursiveRecover() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Restarting after error:", err)
			time.Sleep(time.Second)
			count++
			if count == 5 {
				fmt.Printf("Too much panic: %d", count)
				return
			}
			recursiveRecover()
		}
	}()
	panic(fmt.Sprintf("[ %d ] Panic", count))
}
```

```go
package main

import "fmt"

func main() {
	doPanic()
	// error: 1 and recovered
}

func doRecover() {
	if err := recover(); err != nil {
		fmt.Println("error:", err, "and recovered")
	}
}

func doPanic() {
	defer doRecover()
	panic(1)
}
```

The code prints out *"Hello World!"* because `panic` only exits the function
`panicAndRecover`, not the `main` goroutine(`main` function). And the function
`recursiveRecover` shows an interesting usage to self-recover your program.
Here's another [example](http://play.golang.org/p/Vyrrg1NDQU):

```go
package main
 
import (
	"fmt"
	"log"
	"time"
)
 
func main() {
	keepRunning(5)
}
 
/*
Restarting after error: 2009-11-10 23:00:00 +0000 UTC
Restarting after error: 2009-11-10 23:00:00.001 +0000 UTC
Restarting after error: 2009-11-10 23:00:00.002 +0000 UTC
Restarting after error: 2009-11-10 23:00:00.003 +0000 UTC
Restarting after error: 2009-11-10 23:00:00.004 +0000 UTC
Too much panic: 5
2009/11/10 23:00:00 2009-11-10 23:00:00.004 +0000 UTC
*/
 
var count int
 
func keepRunning(limit int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Restarting after error:", err)
 
			time.Sleep(time.Millisecond)
 
			count++
			if count == limit {
				fmt.Printf("Too much panic: %d\n", count)
				log.Fatal(err)
			}
			keepRunning(limit)
		}
	}()
	run()
}
 
func run() {
	panic(time.Now().String())
}

```

[↑ top](#go-concurrency)
<br><br><br><br>
<hr>








#### **Be careful with `defer` and deadlock**

Again, **`defer`** **_delays the function execution_** **until just
before the enclosing function exit(return)**. That means if the function
does not exit `defer` statement never gets executed:

```go
package main

import (
	"fmt"
	"net/http"
	"sync"
)

type storage struct {
	sync.Mutex
	data string
}

var globalStorage storage

func handler(w http.ResponseWriter, r *http.Request) {
	globalStorage.Lock()
	defer globalStorage.Unlock()

	fmt.Fprintf(w, "Hi %s, I love %s!", globalStorage.data, r.URL.Path[1:])
}

func main() {
	globalStorage.Lock()
	// (X) deadlock!
	// defer globalStorage.Unlock()
	globalStorage.data = "start"
	globalStorage.Unlock()

	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

```

[↑ top](#go-concurrency)
<br><br><br><br>
<hr>











#### channel to communicate

**Go concurrency is about composition of independently executing functions.**
Suppose *multiple goroutines* are running independently at the same time. 
Then how would we **_compose_** and **_coordinate_** them? Go has **channel**:

```go
ch1 := make(chan int)
// same as
ch2 := make(chan int, 0) // unbuffered

ch3 := make(chan int, 1) // make channel with buffer 1
ch3 <- 1 // doesn't block
ch3 <- 2 // blocks until another goroutine receives from the channel
// fatal error: all goroutines are asleep - deadlock!
```

<br>

And try [this](http://play.golang.org/p/xLH5yw0x4P):

```go
package main

import "fmt"

func main() {
	ch := make(chan int, 0) // make channel with buffer 1
	go run(ch)
	fmt.Println(<-ch) // 1
}

func run(ch chan int) {
	ch <- 1
}
```

<br>
**_Channel_** can **_communicate_** and **_signal_** **between goroutines**, as
[here](http://play.golang.org/p/92pWGP9tnU):

```go
package main
 
import "fmt"
 
func main() {
	ch := make(chan struct{})
	go func() {
		fmt.Println("Hello, playground")
		ch <- struct{}{}
	}()
 
	// wait until we receive from channel ch
	<-ch
 
	// Hello, playground
}
```

You can either **_send to_** or **_receive from_** a channel. **A receiver
always blocks until it receives data from a channel. A sender only blocks until
an unbuffered channel receiver has received the value, or buffered channel
receiver has copied the value to the buffer** (when the buffer is full, it
waits until some receiver has retrieved the value). **_Unbuffered channel_**
has a **pending receiver** that would **receive the value as soon as the sender
sends a value**:

<a href="https://www.youtube.com/watch?t=749&v=f6kdp27TYZs" target="_blank"><img
src="http://img.youtube.com/vi/f6kdp27TYZs/0.jpg"></a>


Again, a **_receiver always blocks until it receives data_** from a
**_channel_**. A **_sender only blocks until an unbuffered channel receiver has
received a value,_** or buffered channel receiver has copied the value to the
buffer. **Unbuffered channel has a pending receiver that would receive the
value as soon as the sender sends a value.**

> **_A sender and receiver must both be ready to play their part_** in the
> communication. Otherwise we wait until they are. **It’s a blocking
> operation.** Thus **_channels both communicate and synchronize_** (*in a
> single operation*). **Synchronize** by **sending** on sender's side and
> **receiving** on receiver's side. You **_don’t really need locking_** if you use
> **channel**. You can just use the channel to pass the data back and forth
> between goroutines.
>
> [**_Go Concurrency Patterns by Rob
> Pike_**](https://www.youtube.com/watch?v=f6kdp27TYZs)

<br><br>

Go also has [sync](http://golang.org/pkg/sync) package for low-level
*synchronization*. `sync.WaitGroup` is useful for a collection of goroutines,
as [here](http://play.golang.org/p/rGOt32Ahot):

```go
package main
 
import "sync"
 
func main() {
	ch := make(chan struct{})
	var wg sync.WaitGroup
 
	go func() {
		println(1)
		ch <- struct{}{}
	}()
 
	wg.Add(1)
	go func() {
		println(2)
		wg.Done()
	}()
 
	<-ch
	wg.Wait()
 
	// 1
	// 2
}
```


<br>

We can also use **channel**s to spawn many goroutines and exit the program
after the first receive, as [follows](http://play.golang.org/p/ej0ipwx_r-):

```go
package main
 
import (
	"math/rand"
)
 
func main() {
	ch := make(chan int)
 
	for {
		go func() {
			ch <- rand.Intn(10)
		}()
	}
 
	<-ch
	
	// process took too long
}
```

But this code **would consume all your machine memories**, because `for` loop
runs forever in this code, **_never reaching the channel receivers_**. You have
to set *your own limit*, like [here](http://play.golang.org/p/zdRThCCNtm):

```go
package main
 
import "fmt"
 
func main() {
	ch := make(chan int)
 
	for i := 0; i < 5; i++ {
		go func() {
			ch <- i
		}()
	}
 
	fmt.Println(<-ch) // 5
	fmt.Println(<-ch) // 5
	fmt.Println(<-ch) // 5
	fmt.Println(<-ch) // 5
	fmt.Println(<-ch) // 5
}
```

[↑ top](#go-concurrency)
<br><br><br><br>
<hr>




### questions:
- Senders only block until an unbuffered channel receiver has received the
  value, or buffered channel receiver has copied the value to the buffer (when
  the buffer is full, it waits until some receiver has retrieved the value).
  **_Then is channel synchronous or asynchronous?_**
- Why in the example above, is it **_receiving only ONE value 5_**, not 0, 1,
  2, 3, 4?
- Is there any **_easier way to receive all values from channel?_**


[↑ top](#go-concurrency)
<br><br><br><br>
<hr>




#### #1-1. synchronous, asynchronous channel

**By default, channel is UN-buffered**. And **_unbuffered channel is
synchronous_**. The sender blocks until the receiver has received the value.
The receiver also blocks until there’s a value to receive from the sender.
**Without buffer, every single send will block until another goroutine receives
from the channel.**

<br>

> This allows **goroutines** to **synchronize without explicit locks or
> condition variables.**
>
> [**_Go Tour_**](https://tour.golang.org/concurrency/2)

<br>

**_Buffered channel is asynchronous_**, sending or receiving **does not need to
wait(block)**: it won’t wait for other goroutines to finish. **It only blocks
when all the buffers are full.** *goroutine* waits until some receiver has
retrieved a value and created available buffers. Buffered channels can be
useful when we do not need to synchronize all goroutines completely. The
**capacity(buffer)** of the channel limits the **number of the simultaneous
calls**. Try this [code](http://play.golang.org/p/qHXVeei2th):

```go
package main
 
import "fmt"
 
func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(ch, len(ch), cap(ch))
	// 0x1052d080 2 2
 
	<-ch // 1 is retrieved and discarded
	fmt.Println(ch, len(ch), cap(ch))
	// 0x1052d080 1 2
 
	fmt.Println(<-ch) // 2
	fmt.Println(ch, len(ch), cap(ch))
	// 0x1052d080 0 2
 
	// fmt.Println(<-ch)
	// fatal error: all goroutines are asleep - deadlock!
 
	ch <- 5
	ch <- 10
	fmt.Println(ch, len(ch), cap(ch))
	// 0x1052d080 2 2
}
```

<br>

Buffered channel operates in a non-blocking way. When running 100 million
goroutines with *Intel(R) Core(TM) i7–4910MQ CPU @ 2.90GHz,*
**non-blocking(buffered channel) performs 7 times faster than non-buffered
channel**, as [here](http://play.golang.org/p/_pynER2H1R):

```go
package main
 
import (
	"fmt"
	"log"
	"runtime"
	"time"
)
 
func main() {
	num := 100000000
 
	sendOneTo := func(c chan int) {
		for i := 0; i < num; i++ {
			c <- 1
		}
	}
 
	connect := func(cin, cout chan int) {
		for {
			x := <-cin
			cout <- x
		}
	}
 
	round := func(ch1, ch2, ch3, ch4 chan int) {
		go connect(ch1, ch2)
		go connect(ch2, ch3)
		go connect(ch3, ch4)
		go sendOneTo(ch1)
 
		for i := 0; i < num; i++ {
			_ = <-ch4
		}
	}
 
	startBfCh := time.Now()
	bfCh1 := make(chan int, num)
	bfCh2 := make(chan int, num)
	bfCh3 := make(chan int, num)
	bfCh4 := make(chan int, num)
	round(bfCh1, bfCh2, bfCh3, bfCh4)
	fmt.Println("[Asynchronous, Non-Blocking] Buffered   took", time.Since(startBfCh))
 
	startUnCh := time.Now()
	unCh1 := make(chan int)
	unCh2 := make(chan int)
	unCh3 := make(chan int)
	unCh4 := make(chan int)
	round(unCh1, unCh2, unCh3, unCh4)
	fmt.Println("[Synchronous,  Blocking]      UnBuffered took", time.Since(startUnCh))
}
 
/*
[Asynchronous, Non-Blocking] Buffered   took 32.96282781s     (30 seconds)
[Synchronous,  Blocking]     UnBuffered took 3m17.140920286s  (3 minutes)
*/
 
func init() {
	maxCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.Println("Concurrent execution with", maxCPU, "CPUs.")
}
```

[↑ top](#go-concurrency)
<br><br><br><br>
<hr>




#### #1-2. **buffered channel faster** because it’s *non-blocking*?

Not always. Lack of buffers is, *in most cases*, inconsequential to the
performance, because **unbuffered channel has a pending receiver that would
receive the value as soon as the sender sends a value.** And you also need to
consider the memory overhead.

<br>

> Today, `c = make(chan int, 1 << 31)` is prohibitively expensive.
>
> [*Russ
> Cox*](https://groups.google.com/d/msg/golang-dev/TrOv1E6sIfA/JOEfQTkPLsIJ)

Synchronous channel operation is more deterministic and rigorous because we
know which communication is actually being proceeded, which gives more control
over readers and writers on channels. You need **synchronous(unbuffered)
channel when all communications need to remain in lock-step synchronization.**
**_Asynchronous(buffered) channel is useful where you need more throughput and
responsiveness._**

<br>

> A **buffered channel** can be used like a semaphore, for instance to **limit
> throughput**. In this example, incoming requests are passed to handle, which
> sends a value into the channel, processes the request, and then receives a
> value from the channel to ready the “semaphore” for the next consumer.
> **The capacity of the channel buffer limits the number of simultaneous calls
> to process.**
>
> [*Effective Go*](http://golang.org/doc/effective_go.html#channels)

<br>

```go
// http://golang.org/doc/effective_go.html#channels
var sem = make(chan int, MaxOutstanding)

func handle(r *Request) {
    sem <- 1    // Wait for active queue to drain.
    process(r)  // May take a long time.
    <-sem       // Done; enable next request to run.
}

func Serve(queue chan *Request) {
    for {
        req := <-queue
        go handle(req)  // Don't wait for handle to finish.
    }
}
```

[↑ top](#go-concurrency)
<br><br><br><br>
<hr>




#### #1-3. **be careful with buffered channel!**

> just fixed a bug where we blocked during shutdown sending
> to an error channel;  if **it was buffered we'd have silently**
> **ignored them!** #golang
>
> @davecheney but eventually we'd hit the limit of that buffer on 
> shutdown and it'd freeze;  all the while we would not actually report errs
>
> [*Jason Moiron*](https://twitter.com/jmoiron/status/625084303873998849)
>
>
> @jmoiron the bug sounds like the author didn't consider "what happens
> if the reader of this channel never comes along to pick up this value"
>
> [*Dave Cheney*](https://twitter.com/davecheney/status/625092150376566785)

<br>

![jmoiron_channel_01](img/jmoiron_channel_01.png)
![jmoiron_channel_02](img/jmoiron_channel_02.png)


[↑ top](#go-concurrency)
<br><br><br><br>
<hr>










#### #1-4. **non-deterministic receive from buffered channel**

By default, channel is UNbuffered. And unbuffered channel is synchronous:
- `sender` blocks until the receiver has received the value. 
- `receiver` also blocks until there’s a value to receive from the sender.

Without buffer(unbuffered), every single `send` will block until another goroutine receives from the channel.
Unbuffered channel has a pending receiver that would receive the value as soon as the sender sends a value.

<br>
Buffered channel is asynchronous, sending or receiving does not need to wait(block): it won’t wait for other goroutines to finish.
- `sender` and `receiver` do not block, as long as the buffers are not full.
- You can **send** values to buffered `receiver` channel as long as buffers are not full yet.
- You can **receive** values from buffered `sender` channel as long as buffers are not full yet.

<br>
Therefore, **receiving from a buffered channel** can be non-deterministic
because it does not block whether the values are ready to be received or not.
Try this [code](http://play.golang.org/p/hrVdVWxnCD):

```go
package main

import (
	"fmt"
	"log"
)

func main() {
	bufferedSenderChan := make(chan<- int, 3)
	bufferedReceiverChan := make(<-chan int, 3)

	bufferedSenderChan <- 0
	bufferedSenderChan <- 1
	bufferedSenderChan <- 2

	// defer func() {
	// 	if err := recover(); err != nil {
	// 		fmt.Println(err)
	// 	}
	// }()
	// panic(1)

	// You cannot recover from deadlock!
	// <-bufferedReceiverChan
	// fatal error: all goroutines are asleep - deadlock!

	// 	close(bufferedReceiverChan) // (cannot close receive-only channel)
	// 	fmt.Println(<-bufferedReceiverChan)
	_ = bufferedReceiverChan

	bufferedChan := make(chan int, 3)
	bufferedChan <- 0
	bufferedChan <- 1
	bufferedChan <- 2
	fmt.Println(<-bufferedChan)
	fmt.Println(<-bufferedChan)
	fmt.Println(<-bufferedChan)
	/*
	   0
	   1
	   2
	*/

	fmt.Println()
	for i := 0; i < 10; i++ {
		go func(i int) {
			bufferedChan <- i
		}(i)
	}
	for i := 0; i < 10; i++ {
		fmt.Printf("%v ", <-bufferedChan)
	}
	fmt.Println()
	/*
	   9 0 1 6 7 5 2 3 8 4
	*/

	fmt.Println()
	slice := []float64{23.0, 23, 23, -123.2, 23, 123.2, -2.2, 23.1, -101.2, 17.2}
	sum := 0.0
	for _, elem := range slice {
		sum += elem
	}

	counter1 := NewChannelCounter(0)
	defer counter1.Done()
	defer counter1.Close()

	for _, elem := range slice {
		counter1.Add(elem)
	}
	val1 := counter1.Get()
	if val1 != sum {
		log.Fatalf("NewChannelCounter with No Buffer got wrong. Expected %v but got %v\n", sum, val1)
	}

	counter2 := NewChannelCounter(10)
	defer counter2.Done()
	defer counter2.Close()

	for _, elem := range slice {
		counter2.Add(elem)
	}
	val2 := counter2.Get()
	if val2 != sum {
		log.Fatalf("NewChannelCounter with Buffer got wrong. Expected %v but got %v\n", sum, val2)
	}

	// 2015/08/08 14:03:24 NewChannelCounter with Buffer got wrong. Expected 28.167699999999993 but got 23
}

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
		make(chan float64, buf),
		make(chan float64, buf),
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
```

[↑ top](#go-concurrency)
<br><br><br><br>
<hr>









#### #2. why is this receiving only one value?

```go
package main
 
import "fmt"
 
func main() {
	ch := make(chan int)
 
	for i := 0; i < 5; i++ {
		go func() {
			ch <- i
		}()
	}
 
	fmt.Println(<-ch) // 5
	fmt.Println(<-ch) // 5
	fmt.Println(<-ch) // 5
	fmt.Println(<-ch) // 5
	fmt.Println(<-ch) // 5
}
```


<br>

[Go FAQ](http://golang.org/doc/faq#closures_and_goroutines) explains about
**closures** running as **goroutines**. Try [this](http://play.golang.org/p/8D_tOsXrZW):

```go
package main

import "fmt"

func main() {
	// Deferred function runs
	// in Last In First Out order
	// after the surrounding function returns.
	// NOT AFTER FOR-LOOP
	for i := range []int{0, 1, 2, 3, 4, 5} {
		defer func() {
			fmt.Println("i:", i)
		}()
	}
	fmt.Println()
	/*

	*/

	// variables that are defined ON for-loop
	// should be passed as arguments to the closure
	i1 := 0
	for i := 0; i < 3; i++ {
		i1++
		defer func() {
			fmt.Println("i1:", i, i1)
		}()
	}

	i2 := 0
	for i := 0; i < 3; i++ {
		i2++
		defer func(i, i2 int) {
			fmt.Println("i2:", i, i2)
		}(i, i2)
	}
}

/*
i2: 2 3
i2: 1 2
i2: 0 1
i1: 3 3
i1: 3 3
i1: 3 3
i: 5
i: 5
i: 5
i: 5
i: 5
i: 5
*/

```

**Variables that are defined ON for-loop must be passed as arguments to the
closure.** Again. **_Variables that are defined ON for-loop must be passed as
arguments to the closure._** Try this
[code](http://play.golang.org/p/aoTlXTcRVd):

```go
package main
 
import "fmt"
 
func main() {
	ch1, ch2 := make(chan string), make(chan string)
 
	// variables that are defined ON for-loop
	// should be passed as arguments to the closure
 
	i1 := 0
	for i := 0; i < 3; i++ {
		i1++
		go func() {
			ch1 <- fmt.Sprintf("i1: %d %d", i, i1)
		}()
	}
 
	i2 := 0
	for i := 0; i < 3; i++ {
		i2++
		go func(i, i2 int) {
			ch2 <- fmt.Sprintf("i2: %d %d", i, i2)
		}(i, i2)
	}
 
	for i := 0; i < 3; i++ {
		fmt.Println("ch1:", <-ch1)
	}
 
	for i := 0; i < 3; i++ {
		fmt.Println("ch2:", <-ch2)
	}
 
	/*
		ch1: i1: 3 3
		ch1: i1: 3 3
		ch1: i1: 3 3
		ch2: i2: 0 1
		ch2: i2: 1 2
		ch2: i2: 2 3
	*/
}
```

[↑ top](#go-concurrency)
<br><br><br><br>
<hr>



#### #3. wait for all goroutines to finish with `close`

Here's an **easier way to receive all values from a channel**.
Just use the `for` loop, as [here](http://play.golang.org/p/fpGycVPyGy):

```go
package main
 
import "fmt"
 
func main() {
	ch := make(chan int)
 
	for i := 0; i < 5; i++ {
		go func() {
			ch <- i
		}()
	}
 
	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}
}
```

<br>

`range` can also be used for iterating and receiving from a channel, as
[here](http://play.golang.org/p/-kiyFxUeKx):

```go
package main
 
import "fmt"
 
func main() {
	ch := make(chan int)
 
	for i := 0; i < 5; i++ {
		go func() {
			ch <- i
		}()
	}
 
	for v := range ch {
		fmt.Println(v)
	}
}
 
/*
5
5
5
5
5
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
main.main()
	/tmp/sandbox982202598/main.go:14 +0x1e0
*/
```

This panics with [deadlock](http://en.wikipedia.org/wiki/Deadlock) message
because when we *iterate a channel*, `range` **_ends only after the channel is
closed_**. We MUST **make sure to** **_close the channel_** **after the last
sent value is received by the channel**, as
[here](http://play.golang.org/p/wiK2yjYWcC):

```go
package main
 
import "fmt"
 
func main() {
	ch := make(chan int)
 
	for i := 0; i < 5; i++ {
		go func() {
			ch <- i
		}()
	}
 
	i := 0
	for v := range ch {
		fmt.Println(v)
		i++
		if i == 5 {
			close(ch)
		}
	}
}
 
/*
5
5
5
5
5
*/
```

<br>
And try [this](http://play.golang.org/p/gjPJ1Lrfv1). Note that this code
is not concurrent though:

```go
package main
 
import "fmt"
 
func main() {
	ch := make(chan int)
 
	go func() {
		defer close(ch)
		for i := 0; i < 5; i++ {
			ch <- i
		}
	}()
 
	for v := range ch {
		fmt.Println(v)
	}
	// 0
	// 1
	// 2
	// 3
	// 4
 
	v, ok := <-ch
	fmt.Println(v, ok) // 0 false
	// any value received from closed channel succeeds without blocking
	// , returning the zero value of channel type and false.
}
```

<br>
It should be like [this](http://play.golang.org/p/7o5Hs-0WIS):

```go
package main

import "fmt"

func main() {
	func() {
		ch := make(chan int)
		for i := 0; i < 5; i++ {
			go func(i int) {
				ch <- i
			}(i)
		}
		cn := 0
		for v := range ch {
			fmt.Println(v)
			cn++
			if cn == 5 {
				close(ch)
			}
		}
		// 0
		// 1
		// 2
		// 3
		// 4
		v, ok := <-ch
		fmt.Println(v, ok) // 0 false
		// any value received from closed channel succeeds without blocking
		// , returning the zero value of channel type and false.
	}()
	func() {
		slice := []string{"A", "B", "C", "D", "E"}
		if len(slice) > 0 {
			// this only works when slice length
			// is greater than 0. Otherwise, it will
			// be deadlocking receiving no done message.
			done := make(chan struct{})
			for _, v := range slice {
				go func(v string) {
					fmt.Println("Printing:", v)
					done <- struct{}{}
				}(v)
			}
			cn := 0
			for range done {
				cn++
				if cn == len(slice) {
					close(done)
				}
			}
		}
		/*
			Printing: E
			Printing: A
			Printing: B
			Printing: C
			Printing: D
		*/
	}()
}

```

<br>
Note that **received values from a channel are in order**:

> For channels, the iteration values produced are the **successive** values sent on
> the channel until the channel is closed. If the channel is `nil`, the range
> expression blocks forever.
>
> [Go Spec](https://golang.org/ref/spec#For_statements)


[↑ top](#go-concurrency)
<br><br><br><br>
<hr>





#### **select for channel**: `select` ≠ `switch`
[`select`](https://golang.org/ref/spec#Select_statements) is *like*
[`switch`](https://golang.org/doc/effective_go.html#switch) *for*
**_channels_**. Try this [code](http://play.golang.org/p/Ugbe5aUIQM) with
`switch`:

```go
package main
 
import "fmt"
 
func typeName1(v interface{}) string {
	switch typedValue := v.(type) {
	case int:
		fmt.Println("Value:", typedValue)
		return "int"
	case string:
		fmt.Println("Value:", typedValue)
		return "string"
	default:
		fmt.Println("Value:", typedValue)
		return "unknown"
	}
	panic("unreachable")
}
 
func typeName2(v interface{}) string {
	switch v.(type) {
	case int:
		return "int"
	case string:
		return "string"
	default:
		return "unknown"
	}
	panic("unreachable")
}
 
type Stringer interface {
	String() string
}
 
type fakeString struct {
	content string
}
 
// function used to implement the Stringer interface
func (s *fakeString) String() string {
	return s.content
}
 
func printString(value interface{}) {
	switch str := value.(type) {
	case string:
		fmt.Println(str)
	case Stringer:
		fmt.Println(str.String())
	}
}
 
func main() {
	fmt.Println(typeName1(1))
	fmt.Println(typeName1("Hello"))
	fmt.Println(typeName1(-.1))
	/*
	   Value: 1
	   int
	   Value: Hello
	   string
	   Value: -0.1
	   unknown
	*/
 
	fmt.Println(typeName2(1))       // int
	fmt.Println(typeName2("Hello")) // string
	fmt.Println(typeName2(-.1))     // unknown
 
	s := &fakeString{"Ceci n'est pas un string"}
	printString(s)                // Ceci n'est pas un string
	printString("Hello, Gophers") // Hello, Gophers
}
```

<br>

[**`select`**](https://golang.org/ref/spec#Select_statements) chooses the one that is **firstly ready** to
[*send*](https://golang.org/ref/spec#Send_statements) or
[*receive*](https://golang.org/ref/spec#Receive_operator):


> If one or more of the communications can proceed, a single one that can 
> proceed is chosen via a uniform pseudo-random selection.
>
> Otherwise, if there is a default case, that case is chosen. 
> If there is **no default case**, the "select" statement **blocks until** at 
> least one of the communications can proceed.
>
> [Go Spec](https://golang.org/ref/spec#Select_statements)


<br>

Try [this](http://play.golang.org/p/9OwTUHX7iy):

```go
package main
 
import (
	"fmt"
	"time"
)
 
func send(msg string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
			if i == 5 {
				fmt.Println("Sleeping 2 seconds...")
				time.Sleep(2 * time.Second)
			}
		}
	}()
	return ch
}
 
func main() {
	ch := send("Hello")
	for {
		select {
		case v := <-ch:
			fmt.Println("Received:", v)
		case <-time.After(time.Second):
			fmt.Println("Done!")
			return
		}
	}
}
 
/*
Received: Hello 0
Received: Hello 1
Received: Hello 2
Received: Hello 3
Received: Hello 4
Received: Hello 5
Sleeping 2 seconds...
Done!
*/
```

<br>

Also try this [code](http://play.golang.org/p/1lNjyefPM2%27) from this
[thread](https://groups.google.com/d/msg/golang-nuts/1tjcV80ccq8/W61Z5WjiJKsJ):

```go
package main
 
import (
	"log"
	"time"
)
 
func main() {
	chs := make([]chan struct{}, 100)
 
	// init
	for i := range chs {
		chs[i] = make(chan struct{}, 1)
	}
 
	// close
	for _, ch := range chs {
		close(ch)
	}
 
	// receive
	for _, ch := range chs {
		select {
		case <-ch:
			// https://golang.org/ref/spec#Close
			// After calling close, and after any previously sent values
			// have been received, receive operations will return the zero
			// value for the channel's type without blocking.
			log.Println("Succeed")
 
			// http://golang.org/ref/spec#Select_statements
			// time.After _is_ evaluated each time.
			// https://groups.google.com/d/msg/golang-nuts/1tjcV80ccq8/hcoP9uMNiUcJ
		case <-time.After(time.Millisecond):
			log.Fatalf("Receive Delayed!")
		}
	}
}
 
/*
...
2015/06/27 14:34:48 Succeed
2015/06/27 14:34:48 Succeed
2015/06/27 14:34:48 Succeed
2015/06/27 14:34:48 Succeed
*/
```

<br>
Another example:

```go
package main

import (
	"fmt"
	"net/http"
	"time"
)

var sitesToPing = []string{
	"http://www.google.com",
	"http://www.amazon.com",
	"http://nowebsite.net",
}

func main() {
	respChan, errChan := make(chan string), make(chan error)
	for _, target := range sitesToPing {
		go head(target, respChan, errChan)
	}
	for i := 0; i < len(sitesToPing); i++ {
		select {
		case res := <-respChan:
			fmt.Println(res)
		case err := <-errChan:
			fmt.Println(err)
		case <-time.After(time.Second):
			fmt.Println("Timeout!")
		}
	}
	close(respChan)
	close(errChan)
}

/*
200 / http://www.google.com:OK
405 / http://www.amazon.com:Method Not Allowed
Timeout!
*/

func head(
	target string,
	respChan chan string,
	errChan chan error,
) {
	req, err := http.NewRequest("HEAD", target, nil)
	if err != nil {
		errChan <- fmt.Errorf("0 / %s:None with %v", target, err)
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		errChan <- fmt.Errorf("0 / %s:None with %v", target, err)
		return
	}
	defer resp.Body.Close()
	stCode := resp.StatusCode
	stText := http.StatusText(resp.StatusCode)
	respChan <- fmt.Sprintf("%d / %s:%s", stCode, target, stText)
	return
}
```

[↑ top](#go-concurrency)
<br><br><br><br>
<hr>











#### **receive `nil` from channel**

Try this [code](http://play.golang.org/p/m4P8knWILd). Note that
even if it send `nil` to a channel, it receives:

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	errChan := make(chan error)
	go func() {
		errChan <- nil
	}()
	select {
	case v := <-errChan:
		fmt.Println("even if nil, it still receives", v)
	case <-time.After(time.Second):
		fmt.Println("time-out!")
	}
	// even if nil, it still receives <nil>
}
```

[↑ top](#go-concurrency)
<br><br><br><br>
<hr>












#### `sync.Mutex`, race condition

Let’s say we need synchronization between concurrent tasks. Go recommends
channels for higher-level synchronization, which will be covered shortly. We
can also use `mutex`es. Go package [sync](http://golang.org/pkg/sync/) has
mutexes and they are useful for low-level libraries.
[mutex](http://en.wikipedia.org/wiki/Mutual_exclusion) is **mutual exclusion**,
in order to ensure that no two processes or threads be in the critical section
at the same time. It is important to prevent this kind of [race
conditions](http://en.wikipedia.org/wiki/Race_condition):

> There is no benign race condition.
>
> [**_Dmitry
> Vyukov_**](https://software.intel.com/en-us/blogs/2013/01/06/benign-data-races-what-could-possibly-go-wrong)

**Lock** is a synchronization mechanism to enforce the limits on resource
access when there are many executing threads, therefore preventing the race
condition. Go mutex is a binary
[semaphore](http://en.wikipedia.org/wiki/Semaphore_%28programming%29) (record
of a particular resource’s availability) of either **_locked_** or
**_unlocked_**. This can be used to prevent race conditions.

```go
func (m *Mutex) Lock()
func (m *Mutex) Unlock()
```

Use **Lock** to acquire the mutex, and **Unlock** to release the *Lock*.
Calling **Lock on the same mutex twice** causes
[*deadlock*](http://en.wikipedia.org/wiki/Deadlock), where two or more
competing actions are each waiting for the other to finish and thus neither
ever ends. Code is thread-safe if it manipulates shared data structures in the
way that guarantees safe execution of multiple threads at the same time. Take a
look at the following code, where it has to ensure the mutual exclusion for Go
**map** data structure, which is [**_not
thread-safe_**](https://groups.google.com/d/msg/golang-nuts/3FVAs9dPR8k/Jk9T3s7oIPEJ):

```go
/*
go run -race 31_no_race_surbl_with_mutex.go
*/
package main

import (
	"log"
	"net"
	"net/url"
	"strings"
	"sync"
)

// Data is a set of data in map data structure.
// Every element is unique, and it is unordered.
// It maps its value to frequency.
type Data struct {
	// m maps an element to its frequency
	m map[interface{}]int

	// RWMutex is more expensive
	// https://blogs.oracle.com/roch/entry/beware_of_the_performance_of
	// sync.RWMutex
	//
	// to synchronize access to shared state across multiple goroutines.
	//
	sync.Mutex
}

// NewData returns a new Data.
// Map supports the built-in function "make"
// so we do not have to use "new" and
// "make" does not return pointer.
func NewData() *Data {
	nmap := make(map[interface{}]int)
	return &Data{
		m: nmap,
	}
	// return make(Data)
}

// Init initializes the Data.
func (d *Data) Init() {
	// (X) d = NewData()
	// This only updates its pointer
	// , not the Data itself
	//
	*d = *NewData()
}

// GetSize returns the size of set.
func (d Data) GetSize() int {
	return len(d.m)
}

// IsEmpty returns true if the set is empty.
func (d Data) IsEmpty() bool {
	return d.GetSize() == 0
}

// Insert insert values to the set.
func (d *Data) Insert(items ...interface{}) {
	for _, value := range items {
		d.Lock()
		v, ok := d.m[value]
		d.Unlock()
		if ok {
			d.Lock()
			d.m[value] = v + 1
			d.Unlock()
			continue
		}
		d.Lock()
		d.m[value] = 1
		d.Unlock()
	}
}

func main() {
	d := NewData()
	d.Insert(1, 2, -.9, "A", 0, 2, 2, 2)
	if d.IsEmpty() {
		log.Fatalf("IsEmpty() should return false: %#v", d)
	}
	if d.GetSize() != 5 {
		log.Fatalf("GetSize() should return 5: %#v", d)
	}

	rmap2 := Check(goodSlice...)
	for k, v := range rmap2 {
		if v.IsSpam {
			log.Fatalf("Check | Unexpected %+v %+v but it's ok", k, v)
		}
	}
}

var goodSlice = []string{
	"google.com",
}

// DomainInfo contains domain information from Surbl.org.
type DomainInfo struct {
	IsSpam bool
	Types  []string
}

var nonSpam = DomainInfo{
	IsSpam: false,
	Types:  []string{"none"},
}

var addressMap = map[string]string{
	"2":  "SC: SpamCop web sites",
	"4":  "WS: sa-blacklist web sited",
	"8":  "AB: AbuseButler web sites",
	"16": "PH: Phishing sites",
	"32": "MW: Malware sites",
	"64": "JP: jwSpamSpy + Prolocation sites",
	"68": "WS JP: sa-blacklist web sited jwSpamSpy + Prolocation sites",
}

// Check concurrently checks SURBL spam list.
// http://www.surbl.org/guidelines
// http://www.surbl.org/surbl-analysis
func Check(domains ...string) map[string]DomainInfo {
	final := make(map[string]DomainInfo)
	var wg sync.WaitGroup
	var mutex sync.Mutex
	for _, domain := range domains {
		dom := hosten(domain)
		dmToLook := dom + ".multi.surbl.org"
		wg.Add(1)
		go func() {
			defer wg.Done()
			ads, err := net.LookupHost(dmToLook)
			if err != nil {
				switch err.(type) {
				case net.Error:
					if err.(*net.DNSError).Err == "no such host" {
						mutex.Lock()
						final[dom] = nonSpam
						mutex.Unlock()
					}
				default:
					log.Fatal(err)
				}
			} else {
				stypes := []string{}
				for _, add := range ads {
					tempSlice := strings.Split(add, ".")
					flag := tempSlice[len(tempSlice)-1]
					if val, ok := addressMap[flag]; !ok {
						stypes = append(stypes, "unknown_source")
					} else {
						stypes = append(stypes, val)
					}
				}
				info := DomainInfo{
					IsSpam: true,
					Types:  stypes,
				}
				mutex.Lock()
				final[dom] = info
				mutex.Unlock()
			}
		}()
	}
	wg.Wait()
	return final
}

// hosten returns the host of url.
func hosten(dom string) string {
	dom = strings.TrimSpace(dom)
	var domain string
	if strings.HasPrefix(dom, "http:") ||
		strings.HasPrefix(dom, "https:") {
		dmt, err := url.Parse(dom)
		if err != nil {
			log.Fatal(err)
		}
		domain = dmt.Host
	} else {
		domain = dom
	}
	return domain
}

```

[↑ top](#go-concurrency)
<br><br><br><br>
<hr>




#### Share memory by communicating

<br>

> Do not communicate by sharing memory; instead, **_share memory by
> communicating._**
>
> [**_Go Slogan_**](http://golang.org/doc/effective_go.html#concurrency)

<br>

Go has [race detection tool](http://blog.golang.org/race-detector). And
**following carefully channel-based patterns, we can prevent race conditions
even without locking.**

```
A receiver blocks until it receives data from a channel.
```

**Channel is** **_synchronization_** and **_communication_**. You **don't
really need locking** if you use **channel**. High-level **synchronization** is
**better done via communication of channels**. Then what do we mean by *share
memory by communicating*?

Let's first create some race conditions, with
`go run -race 32_race.go` as [below](http://play.golang.org/p/ROz2Y31Vnb):

```go
/*
go run -race 32_race.go
*/
package main

import "sync"

func updateSliceData(sliceData *[]int, num int, wg *sync.WaitGroup) {
	defer wg.Done()
	*sliceData = append(*sliceData, num)
}

func updateMapData(mapData *map[int]bool, num int, wg *sync.WaitGroup) {
	defer wg.Done()
	(*mapData)[num] = true
}

func main() {
	var wg sync.WaitGroup
	var sliceData = []int{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go updateSliceData(&sliceData, i, &wg)
	}
	wg.Wait()

	var mapData = map[int]bool{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go updateMapData(&mapData, i, &wg)
	}
	wg.Wait()
}

/*
==================
WARNING: DATA RACE
Read by goroutine 5:
  main.updateSliceData()
      /home/ubuntu/race.go:7 +0x5f

Previous write by goroutine 4:
  main.updateSliceData()
      /home/ubuntu/race.go:7 +0x147

Goroutine 5 (running) created at:
  main.main()
      /home/ubuntu/race.go:21 +0xfd

Goroutine 4 (finished) created at:
  main.main()
      /home/ubuntu/race.go:21 +0xfd
==================
==================
WARNING: DATA RACE
Read by goroutine 5:
  runtime.growslice()
      /usr/local/go/src/runtime/slice.go:37 +0x0
  main.updateSliceData()
      /home/ubuntu/race.go:7 +0xcd

Previous write by goroutine 4:
  main.updateSliceData()
      /home/ubuntu/race.go:7 +0x104

Goroutine 5 (running) created at:
  main.main()
      /home/ubuntu/race.go:21 +0xfd

Goroutine 4 (finished) created at:
  main.main()
      /home/ubuntu/race.go:21 +0xfd
==================
==================
WARNING: DATA RACE
Write by goroutine 5:
  runtime.mapassign1()
      /usr/local/go/src/runtime/hashmap.go:383 +0x0
  main.updateMapData()
      /home/ubuntu/race.go:12 +0x94

Previous write by goroutine 4:
  runtime.mapassign1()
      /usr/local/go/src/runtime/hashmap.go:383 +0x0
  main.updateMapData()
      /home/ubuntu/race.go:12 +0x94

Goroutine 5 (running) created at:
  main.main()
      /home/ubuntu/race.go:28 +0x1cf

Goroutine 4 (finished) created at:
  main.main()
      /home/ubuntu/race.go:28 +0x1cf
==================
Found 3 data race(s)
exit status 66

*/

```

This code is creating race conditions. **Go slice and map are [NOT
thread-safe](https://groups.google.com/d/msg/golang-nuts/3FVAs9dPR8k/Jk9T3s7oIPEJ)
data structure**.  They do not prevent you from race-conditions. In the code
above, **_race conditions occur_** when *several goroutines* try to
**_communicate_**—sharing and writing to non thread-safe data structure—**_by sharing
memory_**—running concurrently.

Then what can we do to prevent this? Go has
[*Lock*](http://golang.org/pkg/sync/#Locker):

```go
/*
go run -race 33_no_race_with_mutex.go
*/
package main

import (
	"fmt"
	"sync"
)

func updateSliceDataWithLock(sliceData *[]int, num int, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	mutex.Lock()
	*sliceData = append(*sliceData, num)
	mutex.Unlock()
}

func updateMapDataWithLock(mapData *map[int]bool, num int, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	mutex.Lock()
	(*mapData)[num] = true
	mutex.Unlock()
}

// Mutexes can be created as part of other structures
type sliceData struct {
	sync.Mutex
	s []int
}

func updateSliceDataWithLockStruct(data *sliceData, num int, wg *sync.WaitGroup) {
	defer wg.Done()
	data.Lock()
	data.s = append(data.s, num)
	data.Unlock()
}

// Mutexes can be created as part of other structures
type mapData struct {
	sync.Mutex
	m map[int]bool
}

func updateMapDataWithLockStruct(data *mapData, num int, wg *sync.WaitGroup) {
	defer wg.Done()
	data.Lock()
	data.m[num] = true
	data.Unlock()
}

func main() {
	var (
		wg    sync.WaitGroup
		mutex sync.Mutex
	)

	var ds1 = []int{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go updateSliceDataWithLock(&ds1, i, &wg, &mutex)
	}
	wg.Wait()
	fmt.Println(ds1)

	var dm1 = map[int]bool{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go updateMapDataWithLock(&dm1, i, &wg, &mutex)
	}
	wg.Wait()
	fmt.Println(dm1)

	ds2 := sliceData{}
	ds2.s = []int{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go updateSliceDataWithLockStruct(&ds2, i, &wg)
	}
	wg.Wait()
	fmt.Println(ds2)

	dm2 := mapData{}
	dm2.m = make(map[int]bool)
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go updateMapDataWithLockStruct(&dm2, i, &wg)
	}
	wg.Wait()
	fmt.Println(dm2)
}

```

<br>

But idiomatic Go should use **channels**:

> Concurrent programming in many environments is made difficult by the
> subtleties required to implement correct access to shared variables. 
> **Go encourages a different approach in which shared values are passed around
> on channels** and, in fact, never actively shared by separate threads of
> execution. **Only one goroutine has access to the value at any given time.
> Data races cannot occur, by design.** To encourage this way of thinking we
> have reduced it to a slogan:
>
> Do not communicate by sharing memory; instead, **_share memory by
> communicating._**
>
> [**_Effective Go_**](http://golang.org/doc/effective_go.html#sharing)

<br>

Try this [code](http://play.golang.org/p/jjHd0YyKO7) with
`go run -race 34_no_race_with_channel.go`.
Note that we do **not need to pass pointer of channel**,
because channels, like `map` and `slice`, are **syntactically pointer**,
as explained [here](https://golang.org/doc/faq#references):

```go
/*
go run -race 34_no_race_with_channel.go
*/
package main

// channels were syntactically pointers.
// No need to pass reference.
func sendWithChannel(ch chan int, num int) {
	ch <- num
}

func main() {
	ch1 := make(chan int)
	for i := 0; i < 100; i++ {
		go sendWithChannel(ch1, i)
	}
	cn := 0
	var sliceData = []int{}
	for v := range ch1 {
		sliceData = append(sliceData, v)
		cn++
		if cn == 100 {
			close(ch1)
		}
	}

	ch2 := make(chan int)
	var mapData = map[int]bool{}
	for i := 0; i < 100; i++ {
		go sendWithChannel(ch2, i)
	}
	cn = 0
	for v := range ch2 {
		mapData[v] = true
		cn++
		if cn == 100 {
			close(ch2)
		}
	}
}

```

There is **no race condition** in this code. There is **NO `sync.Mutex`** either.
This is what Go means by:

> Do not communicate by sharing memory; instead, **_share memory by
> communicating._**

<br>
With **channel**, you do not need low-level `sync.Mutex` for synchronization.

<br>
**_Thread_** is a lightweight process since it executes within the context of one
process. Both threads and processes are independent units of execution.
**Threads** under the **same process** **_run in one shared memory_** space,
while **process** **_run in separate memory_** spaces.
Again **multiple threads share the same address space (memory)**, reading
and writing on shared data. That is why, *in multi-threaded programming*,
you need to **synchronize access to memory between threads** (not across processes)
with `Mutex`.

<br>
[*Why goroutines, instead of threads?*](https://golang.org/doc/faq#goroutines) explains:

> Goroutines are part of making concurrency easy to use. The idea,
> is to multiplex independently executing functions(**coroutines**)
> onto a set of threads. When a coroutine blocks, such as by calling
> a blocking system call, the *run-time* automatically **moves other
> coroutines** on the same operating system thread
> **to a different, runnable thread** so they won't be blocked.
> The programmer sees none of this, which is the point. The result, which
> we call goroutines, can be very cheap: they have little overhead beyond
> the memory for the stack, which is just a few kilobytes.
>
> To make the stacks small, Go's run-time uses resizable, bounded stacks.
> A newly minted goroutine is given a few kilobytes, which is almost always
> enough. When it isn't, the run-time grows (and shrinks) the memory for
> storing the stack automatically, allowing many goroutines to live in a
> modest amount of memory. The CPU overhead averages about three cheap
> instructions per function call. It is practical to create hundreds of
> thousands of goroutines in the same address space.
>
> **If goroutines were just threads, system resources would run out at a
> much smaller number**.

<br>
**goroutines** are multiplexed onto multiple OS threads.
When a goroutine blocks on a thread, Go run-time moves other goroutines to a
different, available thread, so they won't be blocked. **goroutine** is cheaper
than **threads**, because **goroutines** are multiplexed onto a small number of
OS threads. A program may run **thousands of goroutines**
*in one thread*. We do not need to allocate one-thread-per-one-goroutine.
We don't need to worry about threads in Go. Go handles synchronization.
One goroutine may be blocked by waiting for I/O, and the thread would
block as well, but **other goroutines would never block** because Go
automatically moves other goroutines to another available thread.
Therefore, Go uses relatively fewer OS threads per Go process.

<br>
<br>
To summarize:
- **goroutines**: non-blocking, light-weight thread.
- **channel**: let the **channel** handle the synchronization for you.

Again, the idea of `Do not communicate by sharing memory; instead,
share memory by communicating.` is to:

- avoid using locking(`sync.Mutex`) if possible, because it's a blocking
  operation and easy to cause deadlocks.
- use **channel** instead, then do not worry about locking.

If you `communicate by sharing memory`, you need to manually
*synchronize access to memory between threads* with locking,
because it shares the same address space. If you `share memory
by communicating`, which means you use **channel** and let the **channel**
handle synchronization, you do not worry about locking and race
conditions.


<br>
Now you can refactor this code above, using **channel** instead of
`sync.Mutex`:

```go
/*
go run -race 31_no_race_surbl_with_mutex.go
*/
package main

import (
	"log"
	"net"
	"net/url"
	"strings"
	"sync"
)

// Data is a set of data in map data structure.
// Every element is unique, and it is unordered.
// It maps its value to frequency.
type Data struct {
	// m maps an element to its frequency
	m map[interface{}]int

	// RWMutex is more expensive
	// https://blogs.oracle.com/roch/entry/beware_of_the_performance_of
	// sync.RWMutex
	//
	// to synchronize access to shared state across multiple goroutines.
	//
	sync.Mutex
}

// NewData returns a new Data.
// Map supports the built-in function "make"
// so we do not have to use "new" and
// "make" does not return pointer.
func NewData() *Data {
	nmap := make(map[interface{}]int)
	return &Data{
		m: nmap,
	}
	// return make(Data)
}

// Init initializes the Data.
func (d *Data) Init() {
	// (X) d = NewData()
	// This only updates its pointer
	// , not the Data itself
	//
	*d = *NewData()
}

// GetSize returns the size of set.
func (d Data) GetSize() int {
	return len(d.m)
}

// IsEmpty returns true if the set is empty.
func (d Data) IsEmpty() bool {
	return d.GetSize() == 0
}

// Insert insert values to the set.
func (d *Data) Insert(items ...interface{}) {
	for _, value := range items {
		d.Lock()
		v, ok := d.m[value]
		d.Unlock()
		if ok {
			d.Lock()
			d.m[value] = v + 1
			d.Unlock()
			continue
		}
		d.Lock()
		d.m[value] = 1
		d.Unlock()
	}
}

func main() {
	d := NewData()
	d.Insert(1, 2, -.9, "A", 0, 2, 2, 2)
	if d.IsEmpty() {
		log.Fatalf("IsEmpty() should return false: %#v", d)
	}
	if d.GetSize() != 5 {
		log.Fatalf("GetSize() should return 5: %#v", d)
	}

	rmap2 := Check(goodSlice...)
	for k, v := range rmap2 {
		if v.IsSpam {
			log.Fatalf("Check | Unexpected %+v %+v but it's ok", k, v)
		}
	}
}

var goodSlice = []string{
	"google.com",
}

// DomainInfo contains domain information from Surbl.org.
type DomainInfo struct {
	IsSpam bool
	Types  []string
}

var nonSpam = DomainInfo{
	IsSpam: false,
	Types:  []string{"none"},
}

var addressMap = map[string]string{
	"2":  "SC: SpamCop web sites",
	"4":  "WS: sa-blacklist web sited",
	"8":  "AB: AbuseButler web sites",
	"16": "PH: Phishing sites",
	"32": "MW: Malware sites",
	"64": "JP: jwSpamSpy + Prolocation sites",
	"68": "WS JP: sa-blacklist web sited jwSpamSpy + Prolocation sites",
}

// Check concurrently checks SURBL spam list.
// http://www.surbl.org/guidelines
// http://www.surbl.org/surbl-analysis
func Check(domains ...string) map[string]DomainInfo {
	final := make(map[string]DomainInfo)
	var wg sync.WaitGroup
	var mutex sync.Mutex
	for _, domain := range domains {
		dom := hosten(domain)
		dmToLook := dom + ".multi.surbl.org"
		wg.Add(1)
		go func() {
			defer wg.Done()
			ads, err := net.LookupHost(dmToLook)
			if err != nil {
				switch err.(type) {
				case net.Error:
					if err.(*net.DNSError).Err == "no such host" {
						mutex.Lock()
						final[dom] = nonSpam
						mutex.Unlock()
					}
				default:
					log.Fatal(err)
				}
			} else {
				stypes := []string{}
				for _, add := range ads {
					tempSlice := strings.Split(add, ".")
					flag := tempSlice[len(tempSlice)-1]
					if val, ok := addressMap[flag]; !ok {
						stypes = append(stypes, "unknown_source")
					} else {
						stypes = append(stypes, val)
					}
				}
				info := DomainInfo{
					IsSpam: true,
					Types:  stypes,
				}
				mutex.Lock()
				final[dom] = info
				mutex.Unlock()
			}
		}()
	}
	wg.Wait()
	return final
}

// hosten returns the host of url.
func hosten(dom string) string {
	dom = strings.TrimSpace(dom)
	var domain string
	if strings.HasPrefix(dom, "http:") ||
		strings.HasPrefix(dom, "https:") {
		dmt, err := url.Parse(dom)
		if err != nil {
			log.Fatal(err)
		}
		domain = dmt.Host
	} else {
		domain = dom
	}
	return domain
}

```

<br>
With **channel**:

```go
/*
go run -race 35_no_race_surbl_with_channel.go
*/
package main

import (
	"fmt"
	"log"
	"net"
	"net/url"
	"strings"
)

// DomainInfo contains domain information from Surbl.org.
type DomainInfo struct {
	Domain string
	IsSpam bool
	Types  []string
}

var addressMap = map[string]string{
	"2":  "SC: SpamCop web sites",
	"4":  "WS: sa-blacklist web sited",
	"8":  "AB: AbuseButler web sites",
	"16": "PH: Phishing sites",
	"32": "MW: Malware sites",
	"64": "JP: jwSpamSpy + Prolocation sites",
	"68": "WS JP: sa-blacklist web sited jwSpamSpy + Prolocation sites",
}

// Check concurrently checks SURBL spam list.
// http://www.surbl.org/guidelines
// http://www.surbl.org/surbl-analysis
func Check(domains ...string) map[string]DomainInfo {
	ch := make(chan DomainInfo)
	for _, domain := range domains {
		go func(domain string) {
			dom := hosten(domain)
			dmToLook := dom + ".multi.surbl.org"
			ads, err := net.LookupHost(dmToLook)
			if err != nil {
				switch err.(type) {
				case net.Error:
					if err.(*net.DNSError).Err == "no such host" {
						nonSpam := DomainInfo{
							Domain: domain,
							IsSpam: false,
							Types:  []string{"none"},
						}
						ch <- nonSpam
					}
				default:
					log.Fatal(err)
				}
			} else {
				stypes := []string{}
				for _, add := range ads {
					tempSlice := strings.Split(add, ".")
					flag := tempSlice[len(tempSlice)-1]
					if val, ok := addressMap[flag]; !ok {
						stypes = append(stypes, "unknown_source")
					} else {
						stypes = append(stypes, val)
					}
				}
				info := DomainInfo{
					Domain: domain,
					IsSpam: true,
					Types:  stypes,
				}
				ch <- info
			}
		}(domain)
	}
	final := make(map[string]DomainInfo)
	checkSize := len(domains)
	cn := 0
	for info := range ch {
		final[info.Domain] = info
		cn++
		if cn == checkSize {
			close(ch)
		}
	}
	return final
}

// hosten returns the host of url.
func hosten(dom string) string {
	dom = strings.TrimSpace(dom)
	var domain string
	if strings.HasPrefix(dom, "http:") ||
		strings.HasPrefix(dom, "https:") {
		dmt, err := url.Parse(dom)
		if err != nil {
			log.Fatal(err)
		}
		domain = dmt.Host
	} else {
		domain = dom
	}
	return domain
}

var goodSlice = []string{
	"google.com", "amazon.com", "yahoo.com", "gmail.com", "walmart.com",
	"stanford.edu", "intel.com", "github.com", "surbl.org", "wikipedia.org",
}

func main() {
	fmt.Println(Check(goodSlice...))
}

```

If you benchmark two versions, you can see that the code with **channel**
is faster than the one with `sync.Mutex`, as
[here](https://github.com/gyuho/surbl/blob/master/benchmark_test.go):

```bash
BenchmarkCheckWithLock        	     100	  73032395 ns/op	  149328 B/op	    1961 allocs/op
BenchmarkCheckWithLock-2      	     100	  73151925 ns/op	  149371 B/op	    1962 allocs/op
BenchmarkCheckWithLock-4      	      50	 124766761 ns/op	  149474 B/op	    1963 allocs/op
BenchmarkCheckWithLock-8      	      50	  22952625 ns/op	  149879 B/op	    1964 allocs/op
BenchmarkCheckWithLock-16     	      50	 126122965 ns/op	  150508 B/op	    1967 allocs/op

BenchmarkCheck                	     100	 184853661 ns/op	  149780 B/op	    1974 allocs/op
BenchmarkCheck-2              	     100	 124283447 ns/op	  149742 B/op	    1974 allocs/op
BenchmarkCheck-4              	     100	 128578550 ns/op	  149758 B/op	    1974 allocs/op
BenchmarkCheck-8              	     100	  74226839 ns/op	  149833 B/op	    1975 allocs/op
BenchmarkCheck-16             	      50	  24317567 ns/op	  149880 B/op	    1975 allocs/op
```

(*But not all the time. It depends on the code. Sometimes channel takes
too much memory and slows down the program.*)

[↑ top](#go-concurrency)
<br><br><br><br>
<hr>










#### memory leak #1

When there is a `defer` statement that never gets run
because the function that contains `defer` is long-running
and never returns or etc:


```go
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func main() {

	fpath := "file.txt"

	f, err := os.OpenFile(fpath, os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		f, err = os.Create(fpath)
		if err != nil {
			panic(err)
		}
	}
	defer f.Close()

	i := 0
	for {
		i++

		// if this is inside a long-running function
		// this never gets run and causes memory leak
		defer func() {
			if _, err := f.WriteString(fmt.Sprintf("LINE %d\n", i)); err != nil {
				panic(err)
			}
		}()
		if i == 100 {
			break
		}
	}

	time.Sleep(time.Second)

	fc, err := toString(fpath)
	fmt.Println(fpath, "contents:", fc)
	// file.txt contents:

	defer func() {
		if err := os.Remove(fpath); err != nil {
			panic(err)
		}
	}()
}

func toString(fpath string) (string, error) {
	file, err := os.Open(fpath)
	if err != nil {
		// NOT retur nil, err
		// []byte can be null but not string
		return "", err
	}
	defer file.Close()

	// func ReadAll(r io.Reader) ([]byte, error)
	tbytes, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(tbytes), nil
}
```

[↑ top](#go-concurrency)
<br><br><br><br>
<hr>











#### `sync/atomic`

In concurrent programming, an operation (or set of operations) is **atomic,
linearizable**, indivisible or un-interruptible if it occurs instantaneously to
the rest of the system. Atomicity is a guarantee of isolation from concurrent
processes. For example, let’s say that we have a web application that has a
shared variable declared globally. Every request spawns its own goroutine, and
if each goroutine tries to manipulate the shared variable, there is a high
probability of race condition. If you just need a [reference
counter](http://en.wikipedia.org/wiki/Reference_counting) in a global scope, Go
[sync/atomic](http://golang.org/pkg/sync/atomic/) would be the simplest way to
ensure the atomicity of a global variable between several goroutines, like
[here](http://play.golang.org/p/97-hvnvssi):

```go
package main
 
import (
	"fmt"
	"log"
	"sync/atomic"
	"time"
)
 
func main() {
	var refCounter int32 = 0
	fmt.Println(atomic.LoadInt32(&refCounter))
	fmt.Println(atomic.AddInt32(&refCounter, 1))
	fmt.Println(atomic.LoadInt32(&refCounter))
	fmt.Println(refCounter)
 
	go func() {
		time.Sleep(10 * time.Second)
		atomic.AddInt32(&refCounter, -1)
	}()
 
	for atomic.LoadInt32(&refCounter) != 0 {
		log.Println("Sleeping 20 seconds")
		time.Sleep(20 * time.Second)
		fmt.Println(refCounter)
	}
	atomic.AddInt32(&refCounter, 1)
	atomic.AddInt32(&refCounter, -1)
}
```

[↑ top](#go-concurrency)
<br><br><br><br>
<hr>




#### web server

Go’s HTTP server spawns a goroutine per request—neither a process nor a
thread:

```go
func (srv *Server) Serve(l net.Listener) error
```

> Serve accepts incoming connections on the Listener `l`, **_creating a new
> service goroutine for each._** The service goroutines read requests and then
> call `srv.Handler` to reply to them.
>
> [http://golang.org/pkg/net/http/#Server.Serve](http://golang.org/pkg/net/http/#Server.Serve)

```go
// Serve accepts incoming connections on the Listener l, creating a
// new service goroutine for each.  The service goroutines read requests and
// then call srv.Handler to reply to them.
func (srv *Server) Serve(l net.Listener) error {
	defer l.Close()
	var tempDelay time.Duration // how long to sleep on accept failure
	for {
		rw, e := l.Accept()
		if e != nil {
			if ne, ok := e.(net.Error); ok && ne.Temporary() {
				if tempDelay == 0 {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay *= 2
				}
				if max := 1 * time.Second; tempDelay > max {
					tempDelay = max
				}
				srv.logf("http: Accept error: %v; retrying in %v", e, tempDelay)
				time.Sleep(tempDelay)
				continue
			}
			return e
		}
		tempDelay = 0
		c, err := srv.newConn(rw)
		if err != nil {
			continue
		}
		c.setState(c.rwc, StateNew) // before Serve can return
		go c.serve()
	}
}
```

And try this:

```go
package main
 
import (
	"net/http"
)
 
func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":3000", nil)
}
 
func foo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "A Go Web Server")
	w.WriteHeader(200)
}
 
/*
curl -i localhost:3000

HTTP/1.1 200 OK
Server: A Go Web Server
Date: Fri, 17 Oct 2014 20:01:38 GMT
Content-Length: 0
Content-Type: text/plain; charset=utf-8
*/
```

[↑ top](#go-concurrency)
<br><br><br><br>
<hr>









#### `sync.Mutex` is just a value

Try [this](http://play.golang.org/p/bmh9CsBnuw):

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

/*

Ian Lance Taylor (https://groups.google.com/d/msg/golang-nuts/7Xi2APcqpM0/UzHJnabiDQAJ):

You are looking at this incorrectly in some way that I don't
understand.  A sync.Mutex is a value with two methods: Lock and
Unlock.  Lock acquires a lock on the mutex.  Unlock releases it.  Only
one goroutine can acquire the lock on the mutex at a time.

That's all there is.  A mutex doesn't have a scope.  It can be a field
of a struct but it doesn't have to be.  A mutex doesn't protect
anything in particular by itself.  You have to write your code to call
Lock, do the protected operations, and then call Unlock.

Your example code looks fine.
*/

func main() {
	var hits struct {
		sync.Mutex
		n int
	}
	hits.Lock()
	hits.n++
	hits.Unlock()
	fmt.Println(hits)
	// {{0 0} 1}

	m := map[string]time.Time{}

	// without this:
	// Found 1 data race(s)
	var mutex sync.Mutex

	done := make(chan struct{})
	for range []int{0, 1} {
		go func() {
			mutex.Lock()
			m[time.Now().String()] = time.Now()
			mutex.Unlock()
			done <- struct{}{}
		}()
	}
	cn := 0
	for range done {
		cn++
		if cn == 2 {
			close(done)
		}
	}
	fmt.Println(m)
	/*
	   map[2015-11-05 20:42:36.516629792 -0800 PST:2015-11-05 20:42:36.516678634 -0800 PST 2015-11-05 20:42:36.516685141 -0800 PST:2015-11-05 20:42:36.516686379 -0800 PST]
	*/
}

```

[↑ top](#go-concurrency)
<br><br><br><br>
<hr>
