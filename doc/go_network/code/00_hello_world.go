package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = ":8080"

func main() {
	mainRouter := http.NewServeMux()
	mainRouter.HandleFunc("/", handler)

	log.Println("Serving http://localhost" + port)
	if err := http.ListenAndServe(port, mainRouter); err != nil {
		panic(err)
	}
}

/*
curl http://localhost:8080
Hello World!
*/

func handler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		fmt.Fprintln(w, "Hello World!")
	default:
		http.Error(w, "Method Not Allowed", 405)
	}
}
