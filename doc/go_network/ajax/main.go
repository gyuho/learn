package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const port = ":8080"

func main() {
	mainRouter := http.NewServeMux()
	mainRouter.Handle("/", http.FileServer(http.Dir("./static")))
	mainRouter.HandleFunc("/metrics", handler)

	log.Println("Serving http://localhost" + port)
	if err := http.ListenAndServe(port, mainRouter); err != nil {
		panic(err)
	}
}

func handler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		resp := struct {
			RefreshInMillisecond int
			Success              bool
			Message              string
		}{
			int(time.Second.Seconds()) * 1000,
			true,
			fmt.Sprintf("%d", time.Now().Nanosecond()),
		}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			panic(err)
		}

	default:
		http.Error(w, "Method Not Allowed", 405)
	}
}
