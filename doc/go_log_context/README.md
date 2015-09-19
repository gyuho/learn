[*back to contents*](https://github.com/gyuho/learn#contents)
<br>

# Go: web server with template, log, context

- [Reference](#reference)
- [`log`](#log)
- [`logrus`](#logrus)
- [`net/context`](#netcontext)

[↑ top](#go-log-context)
<br><br><br><br>
<hr>










#### Reference

- [Using The Log Package In Go](http://www.goinggo.net/2013/11/using-log-package-in-go.html)
- [Sirupsen/logrus](https://github.com/Sirupsen/logrus)
- [`net/context`](https://godoc.org/golang.org/x/net/context)
- [Go Concurrency Patterns: Context](https://blog.golang.org/context)
- [Go's net/context and http.Handler](https://joeshaw.org/net-context-and-http-handler/)

[↑ top](#go-log-context)
<br><br><br><br>
<hr>










#### `log`

Here's how I use [`log`](http://golang.org/pkg/log/) package.

```go
// go run 00_log.go 1>>stdout.log 2>>stderr.log;
package main

import (
	"fmt"
	"log"
)

func main() {
	log.Println("First log message!")
	ss := []int{1, 2, 3}
	fmt.Println(ss)
	fmt.Println(ss[5])
}

```

<br>
And if you run this program with
`go run 00_log.go 1>>stdout.log 2>>stderr.log;`,
you will see output of `log` package pipes into `stderr`
and output of `fmt` package pipes to `stdout`.

[↑ top](#go-log-context)
<br><br><br><br>
<hr>







#### `logrus`

If you need to store logs in files, you can use
[Sirupsen/logrus](https://github.com/Sirupsen/logrus):

```go
package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
)

func openToAppend(fpath string) (*os.File, error) {
	f, err := os.OpenFile(fpath, os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		f, err = os.Create(fpath)
		if err != nil {
			return f, err
		}
	}
	return f, nil
}

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(new(log.JSONFormatter))

	// https://godoc.org/github.com/Sirupsen/logrus#Level
	// log.SetLevel(log.PanicLevel)
	// log.SetLevel(log.FatalLevel)
	// log.SetLevel(log.ErrorLevel)
	// log.SetLevel(log.WarnLevel)
	// log.SetLevel(log.InfoLevel)
	log.SetLevel(log.DebugLevel)
}

func main() {
	lf, err := openToAppend("my.log")
	if err != nil {
		panic(err)
	}
	defer lf.Close()
	// Output to stderr instead of stdout, could also be a file.
	log.SetOutput(lf)

	log.Println("Hello World!") // Println belongs to Infoln
	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	log.Panic("panic")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 100,
	}).Panic("The ice breaks!")
}

```

```
# my.log
{"level":"warning","msg":"The group's number increased tremendously!","number":122,"omg":true,"time":"2015-09-01T23:23:44-07:00"}
{"level":"fatal","msg":"The ice breaks!","number":100,"omg":true,"time":"2015-09-01T23:23:44-07:00"}
{"animal":"walrus","level":"info","msg":"A group of walrus emerges from the ocean","size":10,"time":"2015-09-01T23:25:48-07:00"}
{"level":"warning","msg":"The group's number increased tremendously!","number":122,"omg":true,"time":"2015-09-01T23:25:48-07:00"}
{"level":"panic","msg":"The ice breaks!","number":100,"omg":true,"time":"2015-09-01T23:25:48-07:00"}
{"animal":"walrus","level":"info","msg":"A group of walrus emerges from the ocean","size":10,"time":"2015-09-01T23:27:00-07:00"}
{"level":"warning","msg":"The group's number increased tremendously!","number":122,"omg":true,"time":"2015-09-01T23:27:00-07:00"}
{"level":"panic","msg":"panic","time":"2015-09-01T23:27:00-07:00"}
{"level":"info","msg":"Hello World!","time":"2015-09-01T23:40:25-07:00"}
{"animal":"walrus","level":"info","msg":"A group of walrus emerges from the ocean","size":10,"time":"2015-09-01T23:40:25-07:00"}
{"level":"warning","msg":"The group's number increased tremendously!","number":122,"omg":true,"time":"2015-09-01T23:40:25-07:00"}
{"level":"panic","msg":"panic","time":"2015-09-01T23:40:25-07:00"}

```

[↑ top](#go-log-context)
<br><br><br><br>
<hr>










#### `net/context`

> In Go servers, **each incoming request** is handled in its **own goroutine**.
> Request handlers often start additional goroutines to access backends such
> as databases and RPC services. The set of goroutines working on a request
> typically needs access to request-specific values such as the identity of the
> end user, authorization tokens, and the request's deadline. When a request is
> canceled or times out, all the goroutines working on that request should exit
> quickly so the system can reclaim any resources they are using.
>
> At Google, we developed a `context` package that makes it easy to pass
> **request-scoped values**, cancelation signals, and deadlines across API
> boundaries to all the **goroutines involved in handling a request**.
>
> [*Go Concurrency Patterns: Context*](https://blog.golang.org/context) *by
> Sameer Ajmani*

<br>
> **Incoming requests** to a server should **create a Context**,
> and **outgoing calls** to servers should **accept a Context**.
> The chain of function calls between must **propagate the Context**,
> optionally replacing it  with a modified copy created using WithDeadline,
> WithTimeout, WithCancel, or WithValue.
>
> [package `context`](https://godoc.org/golang.org/x/net/context)

<br>
And here's an simple example of `context` package:

```go
package main

import (
	"fmt"
	"time"

	"golang.org/x/net/context"
)

type key int

const appStartTSKey key = 0

const userIPKey key = 1

const userAgentKey key = 2

func setContextWithAppStartTS(ctx context.Context, ts string) context.Context {
	return context.WithValue(ctx, appStartTSKey, ts)
}

func setContextWithIP(ctx context.Context, userIP string) context.Context {
	return context.WithValue(ctx, userIPKey, userIP)
}

func setContextWithUserAgent(ctx context.Context, userAgent string) context.Context {
	return context.WithValue(ctx, userAgentKey, userAgent)
}

func getAppStartTSFromContext(ctx context.Context) (string, bool) {
	ts, ok := ctx.Value(appStartTSKey).(string)
	return ts, ok
}

func getIPFromContext(ctx context.Context) (string, bool) {
	userIP, ok := ctx.Value(userIPKey).(string)
	return userIP, ok
}

func getUserAgentFromContext(ctx context.Context) (string, bool) {
	userAgent, ok := ctx.Value(userAgentKey).(string)
	return userAgent, ok
}

func main() {
	func() {
		ctx := context.Background()
		ctx = setContextWithAppStartTS(ctx, time.Now().String())
		ctx = setContextWithIP(ctx, "1.2.3.4")
		ctx = setContextWithUserAgent(ctx, "Linux")
		fmt.Println(ctx)
		fmt.Println(getAppStartTSFromContext(ctx))
		fmt.Println(getIPFromContext(ctx))
		fmt.Println(getUserAgentFromContext(ctx))
		fmt.Println("Done 1:", ctx)
	}()
	/*
	   Done 1: context.Background.WithValue(0, "2015-09-02 22:38:00.640981471 -0700 PDT").WithValue(1, "1.2.3.4").WithValue(2, "Linux")
	*/

	fmt.Println()
	func() {
		timeout := 100 * time.Millisecond
		processingTime := time.Nanosecond
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		cancel()
		send(ctx, processingTime)
		fmt.Println("Done 2")
	}()
	/*
		send Timeout: context canceled
		Done 2
	*/

	fmt.Println()
	func() {
		timeout := 100 * time.Millisecond
		processingTime := time.Nanosecond
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		send(ctx, processingTime)
		fmt.Println("Done 3")
	}()
	/*
		send Done!
		Done 3
	*/

	fmt.Println()
	func() {
		timeout := 100 * time.Millisecond
		processingTime := time.Minute
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		send(ctx, processingTime)
		fmt.Println("Done 4")
	}()
	/*
		send Timeout: context deadline exceeded
		Done 4
	*/
}

func send(ctx context.Context, processingTime time.Duration) {
	done := make(chan struct{})
	go func() {
		time.Sleep(processingTime)
		done <- struct{}{}
	}()
	select {
	case <-done:
		fmt.Println("send Done!")
		return
	case <-ctx.Done():
		// Done channel is closed when the deadline expires(times out)
		// or canceled.
		fmt.Println("send Timeout:", ctx.Err())
		return
	}
}

```


<br>
And here's my approach to incorporate `net/context` package with Go standard
[`http`](http://golang.org/pkg/net/http) package *as smoothly as possible*,
and to share contexts, states, information between handlers and other services,
optionally with timeouts, and so on.

Key idea is to understand the [`http.Handler`](http://golang.org/pkg/net/http/#Handler)
interface:

```go
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
```

Any **type** that *implements* `ServeHTTP(ResponseWriter, *Request)` satisfies
`http.Handler` **interface**. For example:

1. type [`http.HandlerFunc`](http://golang.org/pkg/net/http/#HandlerFunc) satisfies `http.Handler` interface,
   with [`ServeHTTP`](http://golang.org/pkg/net/http/#HandlerFunc.ServeHTTP).
2. type [`http.ServeMux`](http://golang.org/pkg/net/http/#ServeMux) satisfies `http.Handler` interface,
   with [`ServeHTTP`](http://golang.org/pkg/net/http/#ServeMux.ServeHTTP).

```go
// The HandlerFunc type is an adapter to allow the use of
// ordinary functions as HTTP handlers.  If f is a function
// with the appropriate signature, HandlerFunc(f) is a
// Handler object that calls f.
type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}

type ServeMux struct {
	mu    sync.RWMutex
	m     map[string]muxEntry
	hosts bool // whether any patterns contain hostnames
}

// ServeHTTP dispatches the request to the handler whose
// pattern most closely matches the request URL.
func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request) {
	if r.RequestURI == "*" {
		if r.ProtoAtLeast(1, 1) {
			w.Header().Set("Connection", "close")
		}
		w.WriteHeader(StatusBadRequest)
		return
	}
	h, _ := mux.Handler(r)
	h.ServeHTTP(w, r)
}

```

<br>
So that they can be used as below:

```
func Handle(pattern string, handler Handler)
func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
func ListenAndServe(addr string, handler Handler) error

func (mux *ServeMux) Handle(pattern string, handler Handler)
func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request))
```

<br>
Then I need to come up with a type that implements `ServeHTTP`
just like `http.HandlerFunc` and `http.ServeMux`, and then satisfy the
interface `http.Handler` so that we can just use the same patterns.
This is a type of function that I want to use globally:

```go
func handler(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
```

But the function signature does not match with `http.HandlerFunc`:

```go
type HandlerFunc func(ResponseWriter, *Request)
```

<br>
Then I need to wrap `context.Context` with methods:

```go
func (ctx context.Context) handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
```

This is exactly what does to satisfy the `http.Handler` interface:

```go
type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}
```

<br>
So now we know that we need to define a type that wraps `context.Context`,
and the type must implement `ServeHTTP`:

```go
type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}


type ContextAdapter struct{
	ctx context.Context
}

func (ca *ContextAdapter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// ContextAdapter should contain some kind of function...
	// so that we can execute func(ResponseWriter, *Request)
	//                     or func(context.Context, ResponseWriter, *Request)
}
```

This tells that I need to create our own interface to define
serving methods. And I can:

```go
type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}


type ContextHandler interface {
	ServeHTTPContext(context.Context, http.ResponseWriter, *http.Request) error
}

// ContextHandlerFunc wraps func(context.Context, ResponseWriter, *Request)
type ContextHandlerFunc func(context.Context, http.ResponseWriter, *http.Request) error

func (f ContextHandlerFunc) ServeHTTPContext(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	return f(ctx, r, req)
}

type ContextAdapter struct{
	ctx     context.Context
	handler ContextHandler // to wrap func(context.Context, ResponseWriter, *Request)
}

func (ca *ContextAdapter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ca.handler.ServeHTTPContext(ca.ctx, r, req)
}
```

<br>
Now I have created an interface `ContextHandler` with a method `ContextHandlerFunc`
in order to wrap `context.Context` in handlers. And with this interface, I have
`ContextAdapter` type that embeds `ContextHandler` interface and implements
`ServeHTTP` so that it can satisfy the interface `http.Handler`. 
Now everything seems ready to use handler with `context.Context`:

```
func handler(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	fmt.Fprintf(w, "Hello World")
	return nil
}

func Handle(pattern string, handler Handler)
func ListenAndServe(addr string, handler Handler) error

func (mux *ServeMux) Handle(pattern string, handler Handler)
```

And here's an example:

```go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/net/context"
)

type ContextHandler interface {
	ServeHTTPContext(context.Context, http.ResponseWriter, *http.Request) error
}

// ContextHandlerFunc wraps func(context.Context, ResponseWriter, *Request)
type ContextHandlerFunc func(context.Context, http.ResponseWriter, *http.Request) error

func (f ContextHandlerFunc) ServeHTTPContext(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	return f(ctx, w, req)
}

// ContextAdapter satisfies:
//	1. interface 'ContextHandler'
//	2. interface 'http.Handler'
type ContextAdapter struct {
	ctx     context.Context
	handler ContextHandler // to wrap func(context.Context, ResponseWriter, *Request) error
}

func (ca *ContextAdapter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if err := ca.handler.ServeHTTPContext(ca.ctx, w, req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
	}
}

// The key type is unexported to prevent collisions with context keys defined in
// other packages.
type key int

const appStartTSKey key = 0

func withTS(h ContextHandler) ContextHandler {
	return ContextHandlerFunc(func(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
		ctx = context.WithValue(ctx, appStartTSKey, time.Now().String())
		return h.ServeHTTPContext(ctx, w, req)
	})
}

func handlerRoot(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	switch req.Method {
	case "GET":
		ts := ctx.Value(appStartTSKey).(string)
		fmt.Fprintf(w, "TS: %s", ts)
		return nil

	default:
		http.Error(w, "Method Not Allowed", 405)
		return fmt.Errorf("Method Not Allowed:", req.Method)
	}
}

func main() {
	rootContext := context.Background()

	mainRouter := http.NewServeMux()

	// values are not shared between handlers!
	// context package is for passing request-scoped values,
	mainRouter.Handle("/", &ContextAdapter{
		ctx:     rootContext,
		handler: withTS(ContextHandlerFunc(handlerRoot)),
	})

	port := ":5000"
	fmt.Println("Serving http://localhost" + port)
	if err := http.ListenAndServe(port, mainRouter); err != nil {
		panic(err)
	}
}

```

<br>
Note that `context` package is for passing request-scoped values.
I was trying to use it to pass values between handlers, but `context`
is not the way of doing it. **Idiomatic Go** only **initialize** 
the `context` within the boundary of an incoming request, and optionally
pass it to the outgoing request from there. If you want to have timeouts
in your downstream jobs, `context` is a great way of implementing it.

[↑ top](#go-log-context)
<br><br><br><br>
<hr>
