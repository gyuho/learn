package count

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"sync"
	"testing"
	"time"
)

var (
	logger = log.New(os.Stdout, "[TEST] ", log.Ldate|log.Ltime)
)

func httpLog(isDebug bool, fn func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if isDebug {
			start := time.Now()
			fn(w, r)
			logger.Printf("%s %s   |  Took %s", r.Method, r.URL.Path, time.Since(start))
		} else {
			fn(w, r)
		}
	}
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintln(w, "Hello, World!")
	default:
		http.Error(w, "Method Not Allowed", 405)
	}
}

func RunHelloWorldHandler(isDebug bool) string {
	mainRouter := http.NewServeMux()
	mainRouter.HandleFunc("/", httpLog(isDebug, helloWorldHandler))
	ts := httptest.NewServer(mainRouter)
	defer ts.Close()
	res, err := http.Get(ts.URL)
	if err != nil {
		panic(err)
	}
	greeting, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		panic(err)
	}
	return string(greeting)
}

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
