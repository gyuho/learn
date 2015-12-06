package main

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	stdlog "log"
)

var (
	port   = ":8080"
	logger = stdlog.New(os.Stdout, "[TEST] ", stdlog.Ldate|stdlog.Ltime)
)

func wrapFunc(fn func(w http.ResponseWriter, req *http.Request)) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		fn(w, req)
		logger.Printf("wrapFunc: %s %s | Took %s", req.Method, req.URL.Path, time.Since(start))
	}
}

func wrapHandlerFunc(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, req)
		logger.Printf("wrapHandlerFunc: %s %s | Took %s", req.Method, req.URL.Path, time.Since(start))
	}
}

func main() {
	go sendRequest(port, "/")
	go sendRequest(port, "/json")
	go sendRequest(port, "/gob")

	mainRouter := http.NewServeMux()
	mainRouter.HandleFunc("/", wrapFunc(handler))
	mainRouter.HandleFunc("/json", wrapHandlerFunc(handlerJSON))
	mainRouter.HandleFunc("/gob", wrapHandlerFunc(handlerGOB))

	stdlog.Println("Serving http://localhost" + port)
	if err := http.ListenAndServe(port, mainRouter); err != nil {
		panic(err)
	}
}

/*
XXXX/XX/XX 17:17:15 Serving http://localhost:8080
[TEST] XXXX/XX/XX 17:17:18 wrapHandlerFunc: GET /json | Took 59.349µs
[TEST] XXXX/XX/XX 17:17:18 wrapFunc: GET / | Took 7.359µs
response for / = Hello World!
[TEST] XXXX/XX/XX 17:17:18 wrapHandlerFunc: GET /gob | Took 132.61µs
response for /json = {Go 1000 XXXX-XX-XX 17:17:18}
response for /gob = {Go 1000 XXXX-XX-XX 17:17:18}
*/

type Data struct {
	Name  string
	Value float64
	TS    string
}

func handler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		fmt.Fprintf(w, "Hello World!")
	default:
		http.Error(w, "Method Not Allowed", 405)
	}
}

func handlerJSON(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		data := Data{}
		data.Name = "Go"
		data.Value = 1000
		data.TS = time.Now().String()[:19]
		if err := json.NewEncoder(w).Encode(data); err != nil {
			panic(err)
		}

	default:
		http.Error(w, "Method Not Allowed", 405)
	}
}

func handlerGOB(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		data := Data{}
		data.Name = "Go"
		data.Value = 1000
		data.TS = time.Now().String()[:19]
		if err := gob.NewEncoder(w).Encode(data); err != nil {
			panic(err)
		}

	default:
		http.Error(w, "Method Not Allowed", 405)
	}
}

func sendRequest(port, endPoint string) {
	time.Sleep(3 * time.Second)

	client := http.DefaultClient
	req, err := http.NewRequest("GET", "http://localhost"+port+endPoint, nil)
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	switch endPoint {
	case "/":
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println("response for", endPoint, "=", string(b))

	case "/json":
		data := Data{}
		for {
			if err := json.NewDecoder(resp.Body).Decode(&data); err == io.EOF {
				break
			} else if err != nil {
				panic(err)
			}
		}
		fmt.Println("response for", endPoint, "=", data)

	case "/gob":
		data := Data{}
		for {
			if err := gob.NewDecoder(resp.Body).Decode(&data); err == io.EOF {
				break
			} else if err != nil {
				panic(err)
			}
		}
		fmt.Println("response for", endPoint, "=", data)
	}
}
