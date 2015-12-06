package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

/*
sudo kill $(sudo netstat -tlpn | perl -ne 'my @a = split /[ \/]+/; print "$a[6]\n" if m/:3000/gio');
*/

// global variable shared by all concurrent requests
var color string

// This is not a good practice because the global variable
// is being affected by race conditions with concurrent web requests.

func main() {
	for i := 0; i < 100; i++ {
		go sendRequestRed()
		go sendRequestBlue()
	}
	http.HandleFunc("/", handler)
	http.HandleFunc("/red", red)
	http.HandleFunc("/blue", blue)
	fmt.Println("Listening to http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		fmt.Fprintf(w, "Hello World! Global color is now %s", color)
	default:
		http.Error(w, "Method Not Allowed", 405)
	}
}

func red(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		color = "red"
		fmt.Fprintf(w, "set red")
	default:
		http.Error(w, "Method Not Allowed", 405)
	}
}

func blue(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		color = "blue"
		fmt.Fprintf(w, "set blue")
	default:
		http.Error(w, "Method Not Allowed", 405)
	}
}

func sendRequestRed() {
	time.Sleep(3 * time.Second)
	client := http.DefaultClient
	req, err := http.NewRequest("GET", "http://localhost:3000/red", nil)
	if err != nil {
		log.Panic(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("response is", string(b))
}

func sendRequestBlue() {
	time.Sleep(3 * time.Second)
	client := http.DefaultClient
	req, err := http.NewRequest("GET", "http://localhost:3000/blue", nil)
	if err != nil {
		log.Panic(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("response is", string(b))
}
