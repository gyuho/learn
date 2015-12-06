package main

import (
	"errors"
	"fmt"
	"net"
	"net/http"
)

// Below, copied from
// https://github.com/camlistore/camlistore/blob/master/pkg/netutil/netutil.go

// localhostLookup looks for a loopback IP by resolving localhost.
func localhostLookup() net.IP {
	// IPv6, it returns 0:0:0:0:0:0:0:1 or ::1
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

// localhost returns the first address found when
// doing a lookup of "localhost". If not successful,
// it looks for an ip on the loopback interfaces.
func localhost() (net.IP, error) {
	if ip := localhostLookup(); ip != nil {
		return ip, nil
	}
	if ip := loopbackIP(); ip != nil {
		return ip, nil
	}
	return nil, errors.New("No loopback ip found.")
}

func main() {
	lp, err := localhost()
	if err != nil {
		panic(err)
	}
	if lp.To4() != nil {
		fmt.Println(lp.String(), "is IPv4!")
	} else if lp.To16() != nil {
		fmt.Println(lp.String(), "is IPv6!")
	}
	// ::1 is IPv6!

	fmt.Println("loopbackIP:", loopbackIP().String())
	// loopbackIP: 127.0.0.1

	mainRouter := http.NewServeMux()
	mainRouter.HandleFunc("/", handler)

	const port = ":8080"
	fmt.Println("Serving http://localhost" + port)
	if err := http.ListenAndServe(port, mainRouter); err != nil {
		panic(err)
	}
}

/*
$ curl http://localhost:8080
r.RemoteAddr: [::1]:52310

$ curl -L http://localhost:8080
r.RemoteAddr: [::1]:52316
*/

// RemoteAddr is for the remote client's IP and port
// (original requester or last proxy address).
// Whenever a client sends a request, its OS open an ephemeral port
// (https://en.wikipedia.org/wiki/Ephemeral_port) to send that request,
// and then the IP and port used for this request is shown in RemoteAddr.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "r.RemoteAddr: %s\n", r.RemoteAddr)
}
