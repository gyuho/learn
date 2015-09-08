package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	global int

	timeChan = time.NewTimer(5 * time.Second).C
	tickChan = time.NewTicker(time.Second).C
)

func readGlobal() {
	for {
		select {
		case <-timeChan:
			fmt.Println("5 seconds passed: done")
			return
		case <-tickChan:
			fmt.Println("global:", global)
		}
	}
}

func httpen(dom string) string {
	dom = strings.TrimSpace(dom)
	if !strings.HasPrefix(dom, "http://") {
		dom = "http://" + dom
	}
	return dom
}

func ping(i int) {
	time.Sleep(2 * time.Second)
	global += i
	req, err := http.NewRequest("HEAD", "http://localhost:5500", nil)
	if err != nil {
		log.Fatal(err)
	}
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	stCode := resp.StatusCode
	stText := http.StatusText(resp.StatusCode)
	log.Printf("[ %d | %s ]", stCode, stText)
}

func main() {
	go readGlobal()
	for i := 0; i < 5; i++ {
		go ping(i)
	}
	http.HandleFunc("/", root)
	http.ListenAndServe(":5500", nil)
}

func root(w http.ResponseWriter, r *http.Request) {
	global++
	w.Write([]byte("OK"))
}

/*
sudo kill $(sudo netstat -tlpn | perl -ne 'my @a = split /[ \/]+/; print "$a[6]\n" if m/:5500 /gio')
global: 0
global: 0
2015/07/05 17:28:23 [ 200 | OK ]
2015/07/05 17:28:23 [ 200 | OK ]
2015/07/05 17:28:23 [ 200 | OK ]
2015/07/05 17:28:23 [ 200 | OK ]
2015/07/05 17:28:23 [ 200 | OK ]
global: 15
global: 15
global: 15
5 seconds passed: done
*/
