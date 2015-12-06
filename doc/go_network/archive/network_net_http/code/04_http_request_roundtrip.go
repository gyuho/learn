package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/context"
)

func main() {
	Do("HEAD", "google.com", time.Second)
	Do("HEAD", "http://httpbin.org/redirect/3", time.Second)
	DoWithRedirectCheck("HEAD", "http://httpbin.org/redirect/3", time.Second)
	RoundTrip("HEAD", "google.com", time.Second)
	RoundTrip("HEAD", "http://httpbin.org/redirect/3", time.Second)
}

/*
2015/10/22 15:04:25 [SUCCESS] HEAD => Took: 615.66134ms [ 200 | OK | http://google.com ]                                                                                                                                      │DE12015/10/22 15:04:26 [SUCCESS] HEAD => Took: 379.768387ms [ 200 | OK | http://httpbin.org/redirect/3 ]                                                                                                                         │kIF2015/10/22 15:04:26 http://httpbin.org/redirect/3 => redirect [1]                                                                                                                                                             │s
2015/10/22 15:04:26 http://httpbin.org/redirect/3 => redirect [2]                                                                                                                                                             │On
2015/10/22 15:04:26 http://httpbin.org/redirect/3 => redirect [3]                                                                                                                                                             │bra
2015/10/22 15:04:26 [SUCCESS] HEAD => Took: 586.503507ms [ 200 | OK | http://httpbin.org/redirect/3 ]                                                                                                                         │nch
2015/10/22 15:04:26 [SUCCESS] HEAD => Took: 146.107978ms [ 301 | Moved Permanently | http://google.com ]                                                                                                                      │ ma
2015/10/22 15:04:27 [SUCCESS] HEAD => Took: 153.807455ms [ 302 | Found | http://httpbin.org/redirect/3 ]
*/

func httpen(dom string) string {
	dom = strings.TrimSpace(dom)
	if !strings.HasPrefix(dom, "http://") {
		dom = "http://" + dom
	}
	return dom
}

// Do sends HTTP requests.
func Do(requestType, target string, timeout time.Duration) {
	run(requestType, target, timeout, do)
}

// DoWithRedirectCheck sends HTTP requests and checks the redirects.
func DoWithRedirectCheck(requestType, target string, timeout time.Duration) {
	run(requestType, target, timeout, doWithRedirectCheck)
}

// RoundTrip sends HTTP requests with RoundTripper.
func RoundTrip(requestType, target string, timeout time.Duration) {
	run(requestType, target, timeout, roundTrip)
}

func run(requestType, target string, timeout time.Duration, f func(string, string) (int, error)) {
	target = httpen(target)
	qt, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	now := time.Now()
	done := make(chan struct{})
	var (
		statusCode int
		errMsg     error
	)
	go func() {
		statusCode, errMsg = f(requestType, target)
		done <- struct{}{}
	}()
	select {
	case <-done:
		if errMsg != nil {
			log.Printf(
				"[ERROR (%v)] %s => Took: %v [ %d | %s | %s ]",
				errMsg,
				requestType,
				time.Since(now),
				statusCode,
				http.StatusText(statusCode),
				target,
			)
		} else {
			log.Printf(
				"[SUCCESS] %s => Took: %v [ %d | %s | %s ]",
				requestType,
				time.Since(now),
				statusCode,
				http.StatusText(statusCode),
				target,
			)
		}
	case <-qt.Done():
		log.Printf(
			"[ERROR - timed out (%v)] %s => Took: %v [ %d | %s | %s ]",
			qt.Err(),
			requestType,
			time.Since(now),
			statusCode,
			http.StatusText(statusCode),
			target,
		)
	}
}

func do(requestType, target string) (int, error) {
	req, err := http.NewRequest(requestType, target, nil)
	if err != nil {
		return -1, err
	}
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return -1, err
	}
	resp.Body.Close()
	return resp.StatusCode, nil
}

func doWithRedirectCheck(requestType, target string) (int, error) {
	req, err := http.NewRequest(requestType, target, nil)
	if err != nil {
		return -1, err
	}
	client := http.DefaultClient

	// Without this, for redirects, all HTTP headers get reset by default
	// https://code.google.com/p/go/issues/detail?id=4800&q=request%20header
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		log.Printf("%s => redirect [%d]", target, len(via))
		// mutate the subsequent redirect requests with the first Header
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return nil
	}

	resp, err := client.Do(req)
	if err != nil {
		return -1, err
	}
	resp.Body.Close()
	return resp.StatusCode, nil
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
		Jar:       nil,
	}

	// Without this, for redirects, all HTTP headers get reset by default
	// https://code.google.com/p/go/issues/detail?id=4800&q=request%20header
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		if len(via) > redirectLimit {
			return fmt.Errorf("%d consecutive requests(redirects)", len(via))
		}
		if len(via) == 0 {
			// no redirect
			return nil
		}
		// mutate the subsequent redirect requests with the first Header
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return nil
	}

	return client
}

// roundTrip executes a single HTTP transaction.
func roundTrip(requestType, target string) (int, error) {
	req, err := http.NewRequest(requestType, target, nil)
	if err != nil {
		return -1, err
	}
	client := newClient(30*time.Second, 30*time.Second, 30*time.Second, 5, true)
	resp, err := client.Transport.RoundTrip(req)
	if err != nil {
		return -1, err
	}
	resp.Body.Close()
	return resp.StatusCode, nil
}
