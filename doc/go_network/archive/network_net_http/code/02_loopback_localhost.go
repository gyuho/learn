package main

import (
	"errors"
	"fmt"
	"net"
	"net/http"
)

// RemoteAddr is for the remote client's IP and port
// (original requester or last proxy address).
// Whenever a client sends a request, its OS open an ephemeral port
// (https://en.wikipedia.org/wiki/Ephemeral_port) to send that request,
// and then the IP and port used for this request is shown in RemoteAddr.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "r.RemoteAddr: %s\n", r.RemoteAddr)
	/*
	   $ curl -L http://localhost:5000
	   r.RemoteAddr: 127.0.0.1:57065

	   $ curl -L http://localhost:5000
	   r.RemoteAddr: 127.0.0.1:57066

	   $ curl -L http://localhost:5000
	   r.RemoteAddr: 127.0.0.1:57067
	*/
}

const port = ":5000"

func handlerResolve(w http.ResponseWriter, r *http.Request) {
	lp, err := Localhost()
	if err != nil {
		panic(err)
	}
	ips := lp.String()
	fmt.Println("ip:", ips) // 127.0.0.1

	netaddr, err := net.ResolveIPAddr("ip4", ips)
	if err != nil {
		panic(err)
	}
	fmt.Println("netaddr:", netaddr.String()) // 127.0.0.1

	conn, err := net.ListenIP("ip4:icmp", netaddr)
	if err != nil {
		panic(err)
	}

	buf := make([]byte, 1024)
	num, _, _ := conn.ReadFrom(buf)
	// build and ping localhost with sudo
	fmt.Printf("ReadPacket: %X\n", buf[:num])
	// ReadPacket: 0800FD6729...
}

func runServer() {
	mainRouter := http.NewServeMux()
	mainRouter.HandleFunc("/", handler)
	mainRouter.HandleFunc("/resolve", handlerResolve)
	fmt.Println("Serving http://localhost" + port)
	if err := http.ListenAndServe(port, mainRouter); err != nil {
		panic(err)
	}
}

func main() {
	runServer()
}

// Below, copied from
// https://github.com/camlistore/camlistore/blob/master/pkg/netutil/netutil.go

// Localhost returns the first address found when
// doing a lookup of "localhost". If not successful,
// it looks for an ip on the loopback interfaces.
func Localhost() (net.IP, error) {
	if ip := localhostLookup(); ip != nil {
		return ip, nil
	}
	if ip := loopbackIP(); ip != nil {
		return ip, nil
	}
	return nil, errors.New("No loopback ip found.")
}

// localhostLookup looks for a loopback IP by resolving localhost.
func localhostLookup() net.IP {
	if ips, err := net.LookupIP("localhost"); err == nil && len(ips) > 0 {
		return ips[0]
	}
	return nil
}

// loopbackIP returns the first loopback IP address sniffing network
// interfaces or nil if none is found.
func loopbackIP() net.IP {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil
	}
	for _, inf := range interfaces {
		const flagUpLoopback = net.FlagUp | net.FlagLoopback
		if inf.Flags&flagUpLoopback == flagUpLoopback {
			addrs, _ := inf.Addrs()
			for _, addr := range addrs {
				ip, _, err := net.ParseCIDR(addr.String())
				if err == nil && ip.IsLoopback() {
					return ip
				}
			}
		}
	}
	return nil
}
