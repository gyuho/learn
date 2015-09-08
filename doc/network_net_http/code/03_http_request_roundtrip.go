package main

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

func main() {
	st1, err := head("http://google.com")
	if err != nil {
		panic(err)
	}
	fmt.Println(st1)

	st2, err := headRoundTrip("http://google.com")
	if err != nil {
		panic(err)
	}
	fmt.Println(st2)

	st3, err := get("http://httpbin.org/redirect/3")
	if err != nil {
		panic(err)
	}
	fmt.Println(st3)
}

/*
`head` Took: 210.14204ms [ 200 | OK | http://google.com ]
200
`headRoundTrip` Took: 83.776081ms [ 301 | Moved Permanently | http://google.com ]
301
Found 1 redirects.
Found 2 redirects.
Found 3 redirects.
`get` Took: 1.056694074s [ 200 | OK | http://httpbin.org/redirect/3 ]
200
*/

func head(target string) (int, error) {
	now := time.Now()
	req, err := http.NewRequest("HEAD", target, nil)
	if err != nil {
		return -1, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		resp.Body.Close()
		return 0, err
	}
	resp.Body.Close()
	stCode := resp.StatusCode
	stText := http.StatusText(resp.StatusCode)
	fmt.Printf("`head` Took: %v [ %d | %s | %s ]\n", time.Since(now), stCode, stText, target)
	return stCode, nil
}

func headRoundTrip(target string) (int, error) {
	now := time.Now()
	req, err := http.NewRequest("HEAD", target, nil)
	if err != nil {
		return -1, err
	}
	client := newClient(30*time.Second, 30*time.Second, 30*time.Second, 5, true)
	resp, err := client.Transport.RoundTrip(req)
	if err != nil {
		resp.Body.Close()
		return 0, err
	}
	resp.Body.Close()
	stCode := resp.StatusCode
	stText := http.StatusText(resp.StatusCode)
	fmt.Printf("`headRoundTrip` Took: %v [ %d | %s | %s ]\n", time.Since(now), stCode, stText, target)
	return stCode, nil
}

func newClient(
	dialTimeout time.Duration,
	responseHeaderTimeout time.Duration,
	responseTimeout time.Duration,
	redirectLimit int,
	disableAlive bool,
) *http.Client {

	// http://golang.org/src/pkg/net/http/transport.go
	dialfunc := func(network, addr string) (net.Conn, error) {
		cn, err := net.DialTimeout(network, addr, dialTimeout)
		if err != nil {
			return nil, err
		}
		return cn, err
	}

	// This will still be type http.RoundTripper
	// If we want to update Transport, we need to type-assert like:
	// client.Transport.(*http.Transport)
	transport := &http.Transport{
		ResponseHeaderTimeout: responseHeaderTimeout,
		Dial:              dialfunc,
		DisableKeepAlives: disableAlive,
	}

	// http://golang.org/src/pkg/net/http/client.go
	client := &http.Client{
		Transport: transport,
		Timeout:   responseTimeout,
		Jar:       nil, // TODO(gyuho): Add Cookie
	}

	// Without this, for redirects, all HTTP headers get reset by default
	// https://code.google.com/p/go/issues/detail?id=4800&q=request%20header
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		if len(via) > redirectLimit {
			return fmt.Errorf("%d consecutive requests(redirects)", len(via))
		}
		if len(via) == 0 {
			fmt.Println("No redirect")
			return nil
		}
		fmt.Printf("Found %d redirects.\n", len(via))
		// mutate the subsequent redirect requests with the first Header
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return nil
	}
	return client
}

func get(target string) (int, error) {
	now := time.Now()
	req, err := http.NewRequest("GET", target, nil)
	if err != nil {
		return -1, err
	}
	client := newClient(30*time.Second, 30*time.Second, 30*time.Second, 5, true)
	resp, err := client.Do(req)
	if err != nil {
		resp.Body.Close()
		return 0, err
	}
	resp.Body.Close()
	stCode := resp.StatusCode
	stText := http.StatusText(resp.StatusCode)
	fmt.Printf("`get` Took: %v [ %d | %s | %s ]\n", time.Since(now), stCode, stText, target)
	return stCode, nil
}
