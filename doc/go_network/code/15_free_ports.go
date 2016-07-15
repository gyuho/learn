package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	ports, err := FreePorts(5, func(num int) bool { return num > 35000 && num < 40000 })
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ports) // [39324 39472 37019 36835 38461]
}

// FreePorts returns maximum n tcp ports available to use.
func FreePorts(n int, matchFunc func(int) bool) (ports []int, err error) {
	if n > 50000 {
		return nil, fmt.Errorf("too many ports requested (%d)", n)
	}
	for len(ports) < n {
		var addr *net.TCPAddr
		addr, err = net.ResolveTCPAddr("tcp", "localhost:0")
		if err != nil {
			break
		}
		var l *net.TCPListener
		l, err = net.ListenTCP("tcp", addr)
		if err != nil {
			break
		}
		port := l.Addr().(*net.TCPAddr).Port
		l.Close()

		if matchFunc == nil {
			ports = append(ports, port)
			continue
		}
		if matchFunc(port) {
			ports = append(ports, port)
		}
	}
	return
}
