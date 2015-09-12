package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"

	"golang.org/x/net/context"
)

/*
sudo kill $(sudo netstat -tlpn | perl -ne 'my @a = split /[ \/]+/; print "$a[6]\n" if m/:5000/gio');
*/

type ContextHandler interface {
	ServeHTTPContext(context.Context, http.ResponseWriter, *http.Request) error
}

// ContextHandlerFunc wraps func(context.Context, ResponseWriter, *Request)
type ContextHandlerFunc func(context.Context, http.ResponseWriter, *http.Request) error

func (f ContextHandlerFunc) ServeHTTPContext(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	return f(ctx, w, req)
}

// ContextAdapter satisfies:
//	1. interface 'ContextHandler'
//	2. interface 'http.Handler'
type ContextAdapter struct {
	ctx     context.Context
	handler ContextHandler // to wrap func(context.Context, ResponseWriter, *Request) error
}

func (ca *ContextAdapter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if err := ca.handler.ServeHTTPContext(ca.ctx, w, req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
	}
}

func main() {
	rootContext := context.Background()
	mainRouter := http.NewServeMux()

	// values are not shared between handlers!
	// context package is for passing request-scoped values,
	mainRouter.Handle("/", &ContextAdapter{
		ctx:     rootContext,
		handler: ContextHandlerFunc(handlerRoot),
	})
	mainRouter.Handle("/root", &ContextAdapter{
		ctx:     rootContext,
		handler: withTS(ContextHandlerFunc(handlerRoot)),
	})
	mainRouter.Handle("/set", &ContextAdapter{
		ctx:     rootContext,
		handler: ContextHandlerFunc(handlerSet),
	})
	mainRouter.Handle("/send1", &ContextAdapter{
		ctx:     rootContext,
		handler: ContextHandlerFunc(handlerSend1),
	})
	mainRouter.Handle("/send2", &ContextAdapter{
		ctx:     rootContext,
		handler: ContextHandlerFunc(handlerSend2),
	})

	port := ":5000"
	fmt.Println("Serving http://localhost" + port)
	if err := http.ListenAndServe(port, mainRouter); err != nil {
		panic(err)
	}
}

/*
sendRequest1 started
sendRequest1 done!

sendRequest2 started
sendRequest2 timed out!
context deadline exceeded
*/

func handlerRoot(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	switch req.Method {
	case "GET":
		fmt.Println("handlerRoot")
		ts, _ := getAppStartTSFromContext(ctx)
		ip, _ := getIPFromContext(ctx)
		ua, _ := getUserAgentFromContext(ctx)
		fmt.Fprintf(w, "Root = AppStartTS: %s / IP: %v / UserAgent: %s", ts, ip, ua)
		return nil

	default:
		http.Error(w, "Method Not Allowed", 405)
		return fmt.Errorf("Method Not Allowed:", req.Method)
	}
}

func handlerSet(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	switch req.Method {
	case "GET":
		fmt.Println("handlerSet")
		ts := time.Now().String()
		ip, err := getIP(req)
		if err != nil {
			return err
		}
		ua := req.UserAgent()
		ctx = setContextWithAppStartTS(ctx, ts)
		ctx = setContextWithIP(ctx, ip)
		ctx = setContextWithUserAgent(ctx, ua)
		fmt.Fprintf(w, "Set = AppStartTS: %s / IP: %v / UserAgent: %s", ts, ip, ua)
		return nil

	default:
		http.Error(w, "Method Not Allowed", 405)
		return fmt.Errorf("Method Not Allowed:", req.Method)
	}
}

func withTS(h ContextHandler) ContextHandler {
	return ContextHandlerFunc(func(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
		ctx = setContextWithAppStartTS(ctx, time.Now().String())
		return h.ServeHTTPContext(ctx, w, req)
	})
}

func handlerSend1(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	switch req.Method {
	case "GET":
		fmt.Fprintf(w, "handlerSend1")
		// Done channel is closed when the deadline expires(times out)
		ctx, cancel := context.WithTimeout(ctx, time.Second)
		defer cancel()
		sendRequest(ctx, "sendRequest1", time.Millisecond)
		return nil

	default:
		http.Error(w, "Method Not Allowed", 405)
		return fmt.Errorf("Method Not Allowed:", req.Method)
	}
}

func handlerSend2(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	switch req.Method {
	case "GET":
		fmt.Fprintf(w, "handlerSend2")
		// Done channel is closed when the deadline expires(times out)
		ctx, cancel := context.WithTimeout(ctx, time.Second)
		defer cancel()
		sendRequest(ctx, "sendRequest2", 2*time.Second)
		return nil

	default:
		http.Error(w, "Method Not Allowed", 405)
		return fmt.Errorf("Method Not Allowed:", req.Method)
	}
}

func sendRequest(ctx context.Context, msg string, duration time.Duration) {
	fmt.Println(msg, "started")
	select {
	case <-time.After(duration):
		fmt.Println(msg, "done!")
		return
	case <-ctx.Done():
		// Done channel is closed when the deadline expires(times out)
		fmt.Println(msg, "timed out!")
		fmt.Println(ctx.Err())
		return
	}
}

// The key type is unexported to prevent collisions with context keys defined in
// other packages.
type key int

const appStartTSKey key = 0

const userIPKey key = 1

const userAgentKey key = 2

func setContextWithAppStartTS(ctx context.Context, ts string) context.Context {
	return context.WithValue(ctx, appStartTSKey, ts)
}

func setContextWithIP(ctx context.Context, userIP net.IP) context.Context {
	return context.WithValue(ctx, userIPKey, userIP)
}

func setContextWithUserAgent(ctx context.Context, userAgent string) context.Context {
	return context.WithValue(ctx, userAgentKey, userAgent)
}

func getAppStartTSFromContext(ctx context.Context) (string, bool) {
	// ctx.Value returns nil if ctx has no value for the key;
	// the string type assertion returns ok=false for nil.
	ts, ok := ctx.Value(appStartTSKey).(string)
	return ts, ok
}

func getIPFromContext(ctx context.Context) (net.IP, bool) {
	userIP, ok := ctx.Value(userIPKey).(net.IP)
	return userIP, ok
}

func getUserAgentFromContext(ctx context.Context) (string, bool) {
	userAgent, ok := ctx.Value(userAgentKey).(string)
	return userAgent, ok
}

// getIP extracts the user IP address from req, if present.
// https://blog.golang.org/context/userip/userip.go
func getIP(req *http.Request) (net.IP, error) {
	ip, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		return nil, fmt.Errorf("userip: %q is not IP:port", req.RemoteAddr)
	}
	userIP := net.ParseIP(ip)
	if userIP == nil {
		return nil, fmt.Errorf("userip: %q is not IP:port", req.RemoteAddr)
	}
	return userIP, nil
}
