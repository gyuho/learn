package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

const port = ":8000"

func main() {
	for i := 0; i < 100; i++ {
		go func() {
			// req := http.Request{}
			resp, err := http.Get("http://localhost" + port)
			if err != nil {
				fmt.Println("err1:", err)
				return
			} else {
				fmt.Println("success")
			}

			// b, err := ioutil.ReadAll(resp.Body)
			// if err != nil {
			// 	fmt.Println("err2:", err)
			// 	return
			// }
			// fmt.Println(string(b))

			_, err = io.Copy(ioutil.Discard, resp.Body)
			if err != nil {
				fmt.Println("err2:", err)
				return
			}
		}()
	}
	mainRouter := http.NewServeMux()
	mainRouter.HandleFunc("/", handler)

	log.Println("Serving http://localhost" + port)
	if err := http.ListenAndServe(port, mainRouter); err != nil {
		panic(err)
	}
}

var cnt int
var used = make(map[string]struct{})

func handler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		if _, ok := used[req.RemoteAddr]; !ok {
			used[req.RemoteAddr] = struct{}{}
		} else {
			fmt.Println("duplicate:", req.RemoteAddr)
		}
		fmt.Println("hello", cnt, req.RemoteAddr)
		fmt.Fprintln(w, "Hello World!", cnt)
		cnt++
	default:
		http.Error(w, "Method Not Allowed", 405)
	}
}
