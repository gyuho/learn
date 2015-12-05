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
