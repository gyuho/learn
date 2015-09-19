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
