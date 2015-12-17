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
