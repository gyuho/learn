[*back to contents*](https://github.com/gyuho/learn#contents)
<br>

# Go: security, rate limit, graceful shutdown

- [Reference](#reference)
- [authentication](#authentication)
- [rate limit](#rate-limit)
- [graceful shutdown](#graceful-shutdown)

[↑ top](#go-security-rate-limit-graceful-shutdown)
<br><br><br><br>
<hr>



#### Reference

- [The OAuth 2.0 Authorization Framework](http://tools.ietf.org/html/rfc6749)
- [Using OAuth 2.0 to Access Google APIs](https://developers.google.com/identity/protocols/OAuth2)
- [Google+ Platform for Web: Sign In Users](https://developers.google.com/+/web/signin/)
- [golang.org/x/oauth2](http://godoc.org/golang.org/x/oauth2)
- [github.com/markbates/goth](https://github.com/markbates/goth)
- [github.com/bradrydzewski/go.auth](https://github.com/bradrydzewski/go.auth)
- [github.com/gyuho/ratelimit](https://github.com/gyuho/ratelimit)
- [Graceful Shutdown, Linger Options, and Socket Closure](https://msdn.microsoft.com/en-us/library/windows/desktop/ms738547(v=vs.85).aspx)
- [github.com/tylerb/graceful](https://github.com/tylerb/graceful)

[↑ top](#go-security-rate-limit-graceful-shutdown)
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
		if user.Email() != "gyuho.cs@gmail.com" {
			fmt.Println("only gyuho.cs@gmail.com can access")
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

[↑ top](#go-security-rate-limit-graceful-shutdown)
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



[↑ top](#go-security-rate-limit-graceful-shutdown)
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

[↑ top](#go-security-rate-limit-graceful-shutdown)
<br><br><br><br>
<hr>
