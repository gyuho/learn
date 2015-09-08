package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

var (
	atomicVar *int32
	countVar  *int32

	timeChan = time.NewTimer(time.Minute).C
	tickChan = time.NewTicker(time.Second).C
)

func readAtomic() {
	for {
		select {
		case <-timeChan:
			fmt.Println("done")
			return
		case <-tickChan:
			fmt.Println("atomicVar:", atomic.LoadInt32(atomicVar))
			fmt.Println("countVar:", atomic.LoadInt32(countVar))
		}
	}
}

func addAtomic() {
	fmt.Println("addAtomic")
	for {
		if atomic.LoadInt32(atomicVar) == 0 {
			atomic.StoreInt32(atomicVar, 1)
			atomic.AddInt32(countVar, 1)
			atomic.StoreInt32(atomicVar, 0)
			break
		}
		time.Sleep(time.Second)
	}
}

func main() {
	var num1 int32
	atomicVar = &num1
	atomic.StoreInt32(atomicVar, 0)

	var num2 int32
	countVar = &num2
	atomic.StoreInt32(countVar, 0)

	go readAtomic()

	http.HandleFunc("/", root)
	http.HandleFunc("/add", add)

	http.ListenAndServe(":3333", nil)
}

func root(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func add(w http.ResponseWriter, r *http.Request) {
	go addAtomic()
	w.Write([]byte("ADD"))
}

/*
localhost:3333
localhost:3333/add
sudo kill $(sudo netstat -tlpn | perl -ne 'my @a = split /[ \/]+/; print "$a[6]\n" if m/:3333 /gio')
atomicVar: 0
countVar: 3
atomicVar: 0
countVar: 3
atomicVar: 0
countVar: 3
addAtomic
atomicVar: 0
countVar: 4
atomicVar: 0
countVar: 4
addAtomic
atomicVar: 0
countVar: 5
atomicVar: 0
countVar: 5
*/
