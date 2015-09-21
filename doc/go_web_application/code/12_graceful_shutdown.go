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
