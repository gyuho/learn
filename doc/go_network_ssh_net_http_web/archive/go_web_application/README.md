[*back to contents*](https://github.com/gyuho/learn#contents)
<br>

# Go: web application

- [Reference](#reference)
- [simple web server](#simple-web-server)
- [`log`](#log)
- [`logrus`](#logrus)
- [`net/context`](#netcontext)
- [`text/template`](#texttemplate)
- [serve image](#serve-image)
- [concurrent web requests](#concurrent-web-requests)
- [authentication](#authentication)
- [rate limit](#rate-limit)
- [graceful shutdown](#graceful-shutdown)
- [**_Web application with Javascript frontend_**](#web-application-with-javascript-frontend)
- [proxy](#proxy)
- [container: docker, rkt](#container-docker-rkt)

[↑ top](#go-web-application)
<br><br><br><br>
<hr>










#### Reference

<br>
**`log`, `context`**:
- [**Sample web application source code by gyuho**](./app)
- [Sirupsen/logrus](https://github.com/Sirupsen/logrus)
- [`net/context`](https://godoc.org/golang.org/x/net/context)
- [Go Concurrency Patterns: Context](https://blog.golang.org/context)
- [**Go's net/context and http.Handler**](https://joeshaw.org/net-context-and-http-handler/)

<br>
**security, rate limit**:
- [The OAuth 2.0 Authorization Framework](http://tools.ietf.org/html/rfc6749)
- [Using OAuth 2.0 to Access Google APIs](https://developers.google.com/identity/protocols/OAuth2)
- [Google+ Platform for Web: Sign In Users](https://developers.google.com/+/web/signin/)
- [golang.org/x/oauth2](http://godoc.org/golang.org/x/oauth2)
- [github.com/markbates/goth](https://github.com/markbates/goth)
- [github.com/bradrydzewski/go.auth](https://github.com/bradrydzewski/go.auth)
- [github.com/gyuho/ratelimit](https://github.com/gyuho/ratelimit)
- [Graceful Shutdown, Linger Options, and Socket Closure](https://msdn.microsoft.com/en-us/library/windows/desktop/ms738547(v=vs.85).aspx)
- [github.com/tylerb/graceful](https://github.com/tylerb/graceful)

<br>
**web application, deploy**:
- [Javascript](https://developer.mozilla.org/en-US/docs/Web/JavaScript)
- [Javascript Wikipedia](https://en.wikipedia.org/wiki/JavaScript)
- [jQuery Wikipedia](https://en.wikipedia.org/wiki/JQuery)
- [Proxy server](https://en.wikipedia.org/wiki/Proxy_server)
- [Reverse proxy](https://en.wikipedia.org/wiki/Reverse_proxy)
- [Nginx](https://en.wikipedia.org/wiki/Nginx)
- [Docker](https://www.docker.com/)
- [Docker Jumpstart, by Andrew Odewahn](https://github.com/odewahn/docker-jumpstart/)
- [Best practices for writing Dockerfiles](https://docs.docker.com/articles/dockerfile_best-practices/)
- [CoreOS](https://coreos.com/)
- [App Container Specification](https://github.com/appc/spec)
- [Amazon Web Service (AWS)](https://aws.amazon.com/)
- [Google Cloud Platform (GCP)](https://cloud.google.com/)
- [**How DNS works**](https://howdns.works/)
- [How the Domain Name System (DNS) Works](http://www.verisign.com/en_US/domain-names/online/how-dns-works/index.xhtml?inc=www.verisigninc.com)

[↑ top](#go-web-application)
<br><br><br><br>
<hr>









#### simple web server

```go
package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr)
	fmt.Println(r.Form)
	fmt.Fprintf(w, "Hello World!")
	// io.WriteString(w, "Hello World!")
}

func no(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "Sorry! Not found!")
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/no", no)
	http.ListenAndServe(":8080", nil)
}

/*
curl http://localhost:8080;
Hello World!
curl http://localhost:8080;
Hello World!

127.0.0.1:42476
map[]
127.0.0.1:42477
map[]
127.0.0.1:42478
map[]
127.0.0.1:42479
map[]
127.0.0.1:42480
map[]
127.0.0.1:42481
map[]
127.0.0.1:42482
map[]
127.0.0.1:42483
map[]
127.0.0.1:42484
map[]
...
*/

```

[↑ top](#go-web-application)
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

[↑ top](#go-web-application)
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

[↑ top](#go-web-application)
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

[↑ top](#go-web-application)
<br><br><br><br>
<hr>











#### `text/template`

[Here](http://play.golang.org/p/V5fh24NbSf)'s an example of
[text/template](http://golang.org/pkg/text/template/) package:

```go
package main
 
import (
	"bytes"
	"fmt"
	"log"
	"text/template"
)
 
func main() {
	tagName := "{{.BranchName}}_{{.Type}}"
	tagStruct := struct {
		BranchName string
		Type       string
	}{
		"gyuho",
		"prod",
	}
	buf := new(bytes.Buffer)
	if err := template.Must(template.New("tmpl").Parse(tagName)).Execute(buf, tagStruct); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.String())
	// gyuho_prod
 
	queryStruct := struct {
		SchemaName string
		TableName  string
		Slice      []map[string]string
		LastIndex  int
	}{
		"my",
		"table",
		[]map[string]string{
			map[string]string{"key": "VARCHAR(100) PRIMARY KEY NOT NULL"},
			map[string]string{"value1": "INTEGER"},
			map[string]string{"value2": "INTEGER"},
		},
		2,
	}
	var queryTmpl = `CREATE TABLE IF NOT EXISTS {{.SchemaName}}.{{.TableName}}  ({{$lastIndex := .LastIndex}}
{{range $index, $valueMap := .Slice}}{{if ne $lastIndex $index}}{{range $key, $value := $valueMap}}	{{$key}} {{$value}},{{end}}
{{else}}{{range $key, $value := $valueMap}}	{{$key}} {{$value}}{{end}}
{{end}}{{end}});`
	tb := new(bytes.Buffer)
	if err := template.Must(template.New("tmpl").Parse(queryTmpl)).Execute(tb, queryStruct); err != nil {
		log.Fatal(err)
	}
	fmt.Println(tb.String())
	/*
	   CREATE TABLE IF NOT EXISTS my.table  (
	   	key VARCHAR(100) PRIMARY KEY NOT NULL,
	   	value1 INTEGER,
	   	value2 INTEGER
	   );
	*/
}
```

[↑ top](#go-web-application)
<br><br><br><br>
<hr>











#### serve image

```go
package main

import (
	"net/http"
	"path"
)

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":3000", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	fp := path.Join(".", "gopherbw.png")
	http.ServeFile(w, r, fp)
}

```

[↑ top](#go-web-application)
<br><br><br><br>
<hr>










#### concurrent web requests

<br>
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
> boundaries to all the goroutines involved in handling a request.
>
> [Go Concurrency Patterns: Context](https://blog.golang.org/context) *by
> Sameer Ajmani*
<br>


```go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

/*
sudo kill $(sudo netstat -tlpn | perl -ne 'my @a = split /[ \/]+/; print "$a[6]\n" if m/:3000/gio');
*/

// global variable shared by all concurrent requests
var color string

// This is not a good practice because the global variable
// is being affected by race conditions with concurrent web requests.

func main() {
	for i := 0; i < 100; i++ {
		go sendRequestRed()
		go sendRequestBlue()
	}
	http.HandleFunc("/", handler)
	http.HandleFunc("/red", red)
	http.HandleFunc("/blue", blue)
	fmt.Println("Listening to http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		fmt.Fprintf(w, "Hello World! Global color is now %s", color)
	default:
		http.Error(w, "Method Not Allowed", 405)
	}
}

func red(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		color = "red"
		fmt.Fprintf(w, "set red")
	default:
		http.Error(w, "Method Not Allowed", 405)
	}
}

func blue(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		color = "blue"
		fmt.Fprintf(w, "set blue")
	default:
		http.Error(w, "Method Not Allowed", 405)
	}
}

func sendRequestRed() {
	time.Sleep(3 * time.Second)
	client := http.DefaultClient
	req, err := http.NewRequest("GET", "http://localhost:3000/red", nil)
	if err != nil {
		log.Panic(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("response is", string(b))
}

func sendRequestBlue() {
	time.Sleep(3 * time.Second)
	client := http.DefaultClient
	req, err := http.NewRequest("GET", "http://localhost:3000/blue", nil)
	if err != nil {
		log.Panic(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("response is", string(b))
}

```


[↑ top](#go-web-application)
<br><br><br><br>
<hr>











#### authentication

[The OAuth 2.0 Authorization Framework](http://tools.ietf.org/html/rfc6749) has
very clear explanation about `OAuth 2.0`. And I will just show you how to use it
based on [github.com/markbates/goth](https://github.com/markbates/goth),
which is based on [golang.org/x/oauth2](http://godoc.org/golang.org/x/oauth2).

For Google sign-in, just follow
[Google+ Platform for Web: Sign In Users](https://developers.google.com/+/web/signin/):

![auth_00](img/auth_00.png)
![auth_01](img/auth_01.png)
![auth_02](img/auth_02.png)
![auth_03](img/auth_03.png)

Then you will be given:
- client ID
- client secret

```go
package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/gplus"
)

/*
sudo kill $(sudo netstat -tlpn | perl -ne 'my @a = split /[ \/]+/; print "$a[6]\n" if m/:8080/gio')
*/

const (
	port              = ":8080"
	gPlusClientID     = "YOUR_KEY"
	gPlusClientSecret = "YOUR_SECRET"

	// sessionName is the key used to access the session store.
	sessionName = "_gothic_session"

	// appKey should be replaced by applications using gothic.
	appKey = "XDZZYmriq8pJ5k8OKqdDuUFym2e7Im5O1MzdyapfotOnrqQ7ZEdTN9AA7K6aPieC"
)

var sessionStore sessions.Store

func init() {
	if sessionStore == nil {
		sessionStore = sessions.NewCookieStore([]byte(appKey))
	}
}

func main() {
	goth.UseProviders(gplus.New(gPlusClientID, gPlusClientSecret, "http://localhost:8080/auth/gplus/callback"))

	mainRouter := http.NewServeMux()
	mainRouter.HandleFunc("/", logInHandler)
	mainRouter.HandleFunc("/auth/gplus", beginAuthHandler)
	mainRouter.HandleFunc("/auth/gplus/callback", callbackHandler)
	mainRouter.HandleFunc("/hello", helloHandler)
	fmt.Println("Serving http://localhost" + port)
	if err := http.ListenAndServe(port, mainRouter); err != nil {
		panic(err)
	}
}

func logInHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		fmt.Fprintf(w, `<p><a href="/auth/gplus">Log in with Google</a></p>`)
	default:
		http.Error(w, "Method Not Allowed", 405)
	}
}

func beginAuthHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		url, err := getAuthURL(w, req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, err)
			return
		}
		http.Redirect(w, req, url, http.StatusTemporaryRedirect)

	default:
		http.Error(w, "Method Not Allowed", 405)
	}
}

func callbackHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		user, err := completeUserAuth(w, req)
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}
		fmt.Println("user:", user)
		http.Redirect(w, req, "http://localhost:8080/hello", http.StatusSeeOther)

	default:
		http.Error(w, "Method Not Allowed", 405)
	}
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		fmt.Fprintf(w, "Hello World!")
	default:
		http.Error(w, "Method Not Allowed", 405)
	}
}

func getAuthURL(w http.ResponseWriter, req *http.Request) (string, error) {
	provider, err := goth.GetProvider("gplus")
	if err != nil {
		return "", err
	}
	sess, err := provider.BeginAuth("state")
	if err != nil {
		return "", err
	}
	url, err := sess.GetAuthURL()
	if err != nil {
		return "", err
	}
	session, _ := sessionStore.Get(req, sessionName)
	session.Values[sessionName] = sess.Marshal()
	if err := session.Save(req, w); err != nil {
		return "", err
	}
	return url, err
}

func completeUserAuth(w http.ResponseWriter, req *http.Request) (goth.User, error) {
	provider, err := goth.GetProvider("gplus")
	if err != nil {
		return goth.User{}, err
	}
	session, _ := sessionStore.Get(req, sessionName)
	if session.Values[sessionName] == nil {
		return goth.User{}, errors.New("could not find a matching session for this request")
	}
	sess, err := provider.UnmarshalSession(session.Values[sessionName].(string))
	if err != nil {
		return goth.User{}, err
	}
	if _, err := sess.Authorize(provider, req.URL.Query()); err != nil {
		return goth.User{}, err
	}
	return provider.FetchUser(sess)
}

```

<br>
There's another example with [github.com/bradrydzewski/go.auth](https://github.com/bradrydzewski/go.auth):

![auth_04](img/auth_04.png)

```go
package main

import (
	"fmt"
	"net/http"

	auth "github.com/bradrydzewski/go.auth"
)

const (
	port               = ":8080"
	isProd             = false
	googleClientID     = "YOUR_KEY"
	googleClientSecret = "YOUR_SECRET"
)

func main() {
	googleRedirect := "http://yourrealwebsite.com/auth/login"
	if !isProd {
		googleRedirect = fmt.Sprintf("http://localhost%s/auth/login", port)
	}
	auth.Config.CookieSecret = []byte("YOUR_COOKIE_SECRET")
	auth.Config.LoginSuccessRedirect = "/secret"
	auth.Config.CookieSecure = false

	mainRouter := http.NewServeMux()
	mainRouter.Handle("/auth/login", auth.Google(googleClientID, googleClientSecret, googleRedirect))
	mainRouter.HandleFunc("/", logInHandler)
	mainRouter.HandleFunc("/secret", auth.SecureFunc(secreteHandler))
	mainRouter.HandleFunc("/private", auth.SecureFunc(secreteHandler))
	mainRouter.HandleFunc("/auth/logout", logoutHandler)

	fmt.Println("Serving http://localhost" + port)
	if err := http.ListenAndServe(port, mainRouter); err != nil {
		panic(err)
	}
}

func logInHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		fmt.Fprintf(w, `<a href="/auth/login"><img class="mainpage" src="https://developers.google.com/+/images/branding/sign-in-buttons/Red-signin_Long_base_44dp.png" style="width:250px;height:50px"></a>`)
	default:
		http.Error(w, "Method Not Allowed", 405)
	}
}

func secreteHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		user, err := auth.GetUserCookie(req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, err)
			return
		}
		if user.Email() != "gyuhox@gmail.com" {
			fmt.Println("only gyuhox@gmail.com can access")
			auth.DeleteUserCookie(w, req)
			http.Redirect(w, req, "http://google.com", http.StatusSeeOther)
		}
		fmt.Fprintf(w, `<a href="/auth/logout">logout</a><br>authorized user: %+v`, user)

	default:
		http.Error(w, "Method Not Allowed", 405)
	}
}

func logoutHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		auth.DeleteUserCookie(w, req)
		http.Redirect(w, req, "/", http.StatusSeeOther)
	default:
		http.Error(w, "Method Not Allowed", 405)
	}
}

```

[↑ top](#go-web-application)
<br><br><br><br>
<hr>








#### rate limit

In production, you want to use
[redis-backed rate limiter](https://github.com/etcinit/speedbump),
but let's implement a simple one:

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	func() {
		tick := time.NewTicker(time.Second / 2)
		// defer tick.Stop()

		cnt := 0
		now := time.Now()
		for t := range tick.C {
			fmt.Println("took:", t.Sub(now))
			fmt.Println("took:", time.Since(now))
			now = time.Now()
			fmt.Println()

			cnt++
			if cnt == 5 {
				tick.Stop()
				break
			}
		}
		/*
		   took: 499.814697ms
		   took: 499.881208ms

		   took: 500.001328ms
		   took: 500.102576ms

		   took: 499.859404ms
		   took: 499.918472ms

		   ...
		*/
	}()

	func() {
		rate := time.Second / 10
		tick := time.NewTicker(rate)
		defer tick.Stop()

		burstNum := 5
		throttle := make(chan time.Time, burstNum)
		go func() {
			for t := range tick.C {
				throttle <- t
			}
		}()
		now := time.Now()
		for range []int{0, 0, 0, 0, 0, 0, 0, 0, 0} {
			// rate limit
			<-throttle
			fmt.Println("took:", time.Since(now))
			now = time.Now()
		}
		/*
		   took: 100.175925ms
		   took: 100.00098ms
		   took: 100.001815ms
		   took: 99.967231ms
		   took: 99.926131ms
		   took: 99.954405ms
		   took: 99.98205ms
		   took: 99.934787ms
		   took: 99.957386ms
		*/
	}()
}

```

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	q := NewQueue(10, time.Second/10)
	tick := time.NewTicker(time.Second / 8)
	now := time.Now()
	for t := range tick.C {
		fmt.Println("took:", time.Since(now))
		now = time.Now()

		isExceeded := q.Push(t)
		if isExceeded {
			fmt.Println(t, "has exceeded the rate limit", q.rate)
			tick.Stop()
			break
		}
	}
	/*
	   took: 125.18325ms
	   took: 124.804029ms
	   took: 124.965582ms
	   took: 124.955137ms
	   took: 124.960788ms
	   took: 124.965522ms
	   took: 124.982369ms
	   took: 124.959646ms
	   took: 124.944575ms
	   took: 124.927333ms
	   2015-09-03 14:45:27.868522371 -0700 PDT has exceeded the rate limit 100ms
	*/
}

// timeSlice stores a slice of time.Time
// in a thread-safe way.
type timeSlice struct {
	times []time.Time

	// RWMutex is more expensive
	// https://blogs.oracle.com/roch/entry/beware_of_the_performance_of
	// sync.RWMutex
	//
	// to synchronize access to shared state across multiple goroutines.
	//
	sync.Mutex
}

func newTimeSlice() *timeSlice {
	tslice := timeSlice{}
	sl := make([]time.Time, 0)
	tslice.times = sl
	return &tslice
}

func (t *timeSlice) push(ts time.Time) {
	t.Lock()
	t.times = append(t.times, ts)
	t.Unlock()
}

func (t *timeSlice) length() int {
	t.Lock()
	d := len(t.times)
	t.Unlock()
	return d
}

func (t *timeSlice) pop() {
	if t.length() != 0 {
		t.Lock()
		t.times = t.times[1:len(t.times):len(t.times)]
		t.Unlock()
	}
}

func (t *timeSlice) first() (time.Time, bool) {
	if t.length() == 0 {
		return time.Time{}, false
	}
	t.Lock()
	v := t.times[0]
	t.Unlock()
	return v, true
}

// Queue contains the slice of timestamps
// and other rate limiter configurations.
type Queue struct {
	slice *timeSlice

	// burstSize is like a buffer.
	// If burstSize is 5, it allows rate exceeding
	// for the fist 5 elements.
	burstSize int
	rate      time.Duration
}

// NewQueue returns a new Queue.
func NewQueue(burstSize int, rate time.Duration) *Queue {
	tslice := newTimeSlice()
	q := Queue{}
	q.slice = tslice
	q.burstSize = burstSize
	q.rate = rate
	return &q
}

// Push appends the timestamp to the Queue.
// It return true if rate has exceeded.
// We need a pointer of Queue, where it defines
// timeSlice with pointer as well. To append to slice
// and update struct members, we need pointer types.
func (q *Queue) Push(ts time.Time) bool {
	if q.slice.length() == q.burstSize {
		q.slice.pop()
	}
	q.slice.push(ts)
	if q.slice.length() < q.burstSize {
		return false
	}
	ft, ok := q.slice.first()
	if !ok {
		return false
	}
	diff := ft.Sub(ts)
	return q.rate > diff
}

func (q *Queue) String() string {
	return fmt.Sprintf("times: %+v / burstSize: %d / rate: %v", q.slice.times, q.burstSize, q.rate)
}

```

<br>
And you can maintain your own global map to limit rates per IP:

```go
// ip address to Queue.
var ipToQueue = map[string]Queue{}
```



[↑ top](#go-web-application)
<br><br><br><br>
<hr>









#### graceful shutdown

What is *graceful shutdown*? First, here's **NOT** graceful shutdown
of web server:

```go
package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	stdlog "log"
)

/*
sudo kill $(sudo netstat -tlpn | perl -ne 'my @a = split /[ \/]+/; print "$a[6]\n" if m/:5000/gio')
*/

func main() {
	const port = ":5000"
	mainRouter := http.NewServeMux()
	mainRouter.HandleFunc("/", wrapHandlerFunc(handler))
	fmt.Println("Serving http://localhost" + port)
	if err := http.ListenAndServe(port, mainRouter); err != nil {
		panic(err)
	}
}

var logger = stdlog.New(os.Stdout, "[TEST] ", stdlog.Ldate|stdlog.Ltime)

func wrapHandlerFunc(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, req)
		logger.Printf("%s %s   |  Took %s", req.Method, req.URL.Path, time.Since(start))
	}
}

func handler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		fmt.Fprintf(w, "Hello World!")
	default:
		http.Error(w, "Method Not Allowed", 405)
	}
}

```

There are many ways to kill/stop this web server (*in Ubuntu*):

- Just close the running terminal process.
- Run `sudo kill $(sudo netstat -tlpn | perl -ne 'my @a = split /[ \/]+/; print "$a[6]\n" if m/:5000/gio')`
- <kbd>Ctrl</kbd> + <kbd>c</kbd> to kill the process with `SIGINT`, but it may not work.
- <kbd>Ctrl</kbd> + <kbd>z</kbd> to suspend the process with `SIGSTOP`, but it is not graceful. The port that was being used is not freed. You will get `panic: listen tcp :5000: bind: address already in use` message in your second run.

And they are not graceful shutdown in that:

- It looses all active connections and states.
- There is no mechanism to allow other processes to listen to that port immediately.

<br>
Here's an example of graceful shutdown with
[graceful](https://github.com/tylerb/graceful):

```go
package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	stdlog "log"

	graceful "gopkg.in/tylerb/graceful.v1"
)

/*
sudo kill $(sudo netstat -tlpn | perl -ne 'my @a = split /[ \/]+/; print "$a[6]\n" if m/:5001/gio')
*/

func main() {
	const port = ":5001"
	mainRouter := http.NewServeMux()
	mainRouter.HandleFunc("/", wrapHandlerFunc(handler))
	fmt.Println("Serving http://localhost" + port)
	// if err := http.ListenAndServe(port, mainRouter); err != nil {
	// 	panic(err)
	// }
	graceful.Run(port, 10*time.Second, mainRouter)
}

var logger = stdlog.New(os.Stdout, "[TEST] ", stdlog.Ldate|stdlog.Ltime)

func wrapHandlerFunc(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, req)
		logger.Printf("%s %s   |  Took %s", req.Method, req.URL.Path, time.Since(start))
	}
}

func handler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		fmt.Fprintf(w, "Hello World!")
	default:
		http.Error(w, "Method Not Allowed", 405)
	}
}

/*
ubuntu@ubuntu:~$ ps aux | grep 05_
ubuntu   16709  0.0  0.0  29884  5152 pts/11   Sl+  15:12   0:00 ./05_graceful_shutdown
ubuntu   17063  0.0  0.0  17164  2244 pts/14   S+   15:13   0:00 grep --color=auto 05_

ubuntu@ubuntu:~$ kill -SIGINT 16709;
*/

```

[↑ top](#go-web-application)
<br><br><br><br>
<hr>











#### **_Web application with Javascript frontend_**

Output of [**sample web application source code by gyuho**](./app):

<br>
![main](img/app0.png)

<img src="img/app01.png"
alt="app01"
width="150" height="120" />
<br>

[↑ top](#go-web-application)
<br><br><br><br>
<hr>








#### proxy

**`HTTP` allows intermediate networks**. `HTTP` proxy sits between client and
server.

> In computer networks, a **proxy server** is a server (a computer system or an
> application) that acts as an **intermediary** for **requests from clients**
> seeking resources from other servers. A client connects to the proxy server,
> requesting some service, such as a file, connection, web page, or other
> resource available from a different server and the proxy server **evaluates
> the request** as a way to simplify and control its complexity.
>
> A reverse proxy is usually an Internet-facing proxy used as a front-end to
> control and **protect access to a server** on a private network. A reverse
> proxy commonly also performs tasks such as load-balancing, authentication,
> decryption or caching.
>
> A **reverse proxy** (or surrogate) is a proxy server that *appears to clients* to
> be an *ordinary server*. **Requests are forwarded to** one or more **proxy servers**
> which handle the request. The **response from the proxy server** is returned *as
> if* it came directly from the **original server**, leaving the client no knowledge
> of the origin servers.
>
> [*Proxy server*](https://en.wikipedia.org/wiki/Proxy_server) *by Wikipedia*

<br>
> A reverse proxy taking requests from the Internet and forwarding them to
> servers in an internal network. Those making requests to the proxy may not be
> aware of the internal network.
> 
> Reverse proxies can hide the existence and characteristics of an origin
> server or servers.
>
> [*Reverse proxy*](https://en.wikipedia.org/wiki/Reverse_proxy) *by Wikipedia*

<br>
We can use [Nginx](http://wiki.nginx.org/Main) as an HTTP server, reverse proxy
along with Go web servers:

![proxy_reverse](img/proxy_reverse.png)

<br>
Then why do we bother to run another web server, or reverse proxy while we can
do pretty much everything in Go?

```go
package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}

```

<br>
It's because popular web proxies like `Nginx` provides useful features
out-of-the box. So it's not to reinvent the wheels while we can just add
another module to `Nginx` configuration. `Nginx` provides:

- Rate limiting.
- Access, error logs.
- Serve static files with `try_files`.
- Auth, compression support.
- Serve cached contents while the application is down.

For more, please visit [Nginx wiki](http://wiki.nginx.org/Main).

[↑ top](#go-web-application)
<br><br><br><br>
<hr>










#### container: docker, rkt

These are great introductory articles about Docker, rkt, containers:

- https://www.docker.com/whatisdocker
- https://coreos.com/using-coreos/containers/
- https://coreos.com/rkt/docs/latest/

And this is just how I understand them.

<br>
**Docker** is like a big binary file for an application.
Just like Unix, you can compile an application into one file, and run it as a
process. **Docker** is an application container, not a system container.

Then what is *container*? Containers *contain* the *application* and all its
*dependencies*, but they share the kernel with other containers. And containers
are running as a separate process under its host operating system. (It's
recommended that you [run only one process per
container](https://docs.docker.com/articles/dockerfile_best-practices/).)
It's usually faster, lighter than VMs. Pre-existing techonologies include:

- [*`cgroups`*](https://en.wikipedia.org/wiki/Cgroups)
- [*`LXC`*](https://en.wikipedia.org/wiki/LXC)
- [*`namespace`*](http://man7.org/linux/man-pages/man7/namespaces.7.html)
- [*`AUFS`*](https://en.wikipedia.org/wiki/Aufs)

<br>
> The most powerful feature of containers is the ability to run any Linux
> userland that's compatible with the latest kernel.
>
> [*Container Overview*](https://coreos.com/using-coreos/containers/) *by
> CoreOS*

<br>
Software engineering becomes frustrating when you have to deal with
inconsistent development, production environments. This is the core problem
that containerization solves, by **documenting** those processes in
`Dockerfile` or `App Container Specification`. It helps define and maintain
homogeneous *dev/test/prod* environments in a reproduceable way. There are
many other reasons why you would use
[containers](http://kubernetes.io/v1.0/docs/whatisk8s.html):

- Application centric development with higher level of abstraction running an
  application on an OS using logical resources, not having to worry about
  running an OS on virtual hardware.
- Separation of development and operation, build and deployment.
- Container image immutability provides continous development, quick and easy
  way to roll-back.
- Good for distributed, independent micro-services.
- Consistent development, testing, production environments.
- Portable with any Cloud and OS.
- Resource isolation and utilization.

<br>
Then let's write actual `Dockerfile` and `App Container Specification`
to deploy my
[web application](./app):

<br>
**`Docker`**:

```
FROM ubuntu:14.04
RUN apt-get update
RUN apt-get install -y nginx
ADD ./nginx.conf /etc/nginx/sites-available/default
RUN service nginx restart

# automatically copies the package source,
# fetches the application dependencies
# builds the program, and configures it to run on startup.
FROM golang:onbuild
EXPOSE 8080

```

```sh
#!/bin/bash
sudo docker build -t app .;
sudo docker run --publish 8080:8080 --name test --rm app;

```


<br>
<br>
**`App Container Specification`**:

**TODO: THIS IS NOT WORKING... (https://github.com/appc/spec/issues/480)**


```json
{
	"acKind": "ImageManifest",
	"acVersion": "0.6.1",
	"name": "gyuho/app",
	"labels": [
		{
			"name": "version",
			"value": "0.0.1"
		},
		{
			"name": "arch",
			"value": "amd64"
		},
		{
			"name": "os",
			"value": "linux"
		}
	],
	"app": {
		"user": "root",
		"group": "root",
		"exec": [
			"/usr/bin/app"
		],
		"mountPoints": [
			{
				"name": "static",
				"path": "/usr/bin/static"
			},
			{
				"name": "templates",
				"path": "/usr/bin/templates"
			}
		],
		"ports": [
			{
				"name": "web-server",
				"protocol": "tcp",
				"port": 8080
			}
		]
	},
	"annotations": [
		{
			"name": "authors",
			"value": "Gyu-Ho Lee <gyuhox@gmail.com>"
		}
	]
}

```

```sh
#!/bin/bash
# https://github.com/coreos/rkt/blob/master/Documentation/getting-started-guide.md

CGO_ENABLED=0 GOOS=linux go build -o app -a -installsuffix cgo .;
file app;
ldd app;
sudo ./actool --debug validate manifest.json;

mkdir -p image/rootfs/usr/bin;

sudo cp manifest.json image/manifest;

sudo cp app image/rootfs/usr/bin;
sudo cp -rf static/ image/rootfs/usr/bin;
sudo cp -rf templates/ image/rootfs/usr/bin;

sudo ./actool build --overwrite image/ app-0.0.1-linux-amd64.aci;
sudo ./actool --debug validate app-0.0.1-linux-amd64.aci;

sudo ./rkt metadata-service >/dev/null 2>&1 & # run in background

sudo ./rkt --insecure-skip-verify run \
app-0.0.1-linux-amd64.aci \
--volume static,kind=host,source=/usr/bin/static \
--volume templates,kind=host,source=/usr/bin/templates \
-- \
;

```

[↑ top](#go-web-application)
<br><br><br><br>
<hr>
