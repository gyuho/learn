package main

import (
	"fmt"
	"time"

	"golang.org/x/net/context"
)

type key int

const appStartTSKey key = 0

const userIPKey key = 1

const userAgentKey key = 2

func setContextWithAppStartTS(ctx context.Context, ts string) context.Context {
	return context.WithValue(ctx, appStartTSKey, ts)
}

func setContextWithIP(ctx context.Context, userIP string) context.Context {
	return context.WithValue(ctx, userIPKey, userIP)
}

func setContextWithUserAgent(ctx context.Context, userAgent string) context.Context {
	return context.WithValue(ctx, userAgentKey, userAgent)
}

func getAppStartTSFromContext(ctx context.Context) (string, bool) {
	ts, ok := ctx.Value(appStartTSKey).(string)
	return ts, ok
}

func getIPFromContext(ctx context.Context) (string, bool) {
	userIP, ok := ctx.Value(userIPKey).(string)
	return userIP, ok
}

func getUserAgentFromContext(ctx context.Context) (string, bool) {
	userAgent, ok := ctx.Value(userAgentKey).(string)
	return userAgent, ok
}

func main() {
	func() {
		ctx := context.Background()
		ctx = setContextWithAppStartTS(ctx, time.Now().String())
		ctx = setContextWithIP(ctx, "1.2.3.4")
		ctx = setContextWithUserAgent(ctx, "Linux")
		fmt.Println(ctx)
		fmt.Println(getAppStartTSFromContext(ctx))
		fmt.Println(getIPFromContext(ctx))
		fmt.Println(getUserAgentFromContext(ctx))
		fmt.Println("Done 1:", ctx)
	}()
	/*
	   Done 1: context.Background.WithValue(0, "2015-09-02 22:38:00.640981471 -0700 PDT").WithValue(1, "1.2.3.4").WithValue(2, "Linux")
	*/

	fmt.Println()
	func() {
		// Done channel is closed when the deadline expires(times out)
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
		ctx = setContextWithUserAgent(ctx, "Linux")
		cancel()
		select {
		case <-time.After(200 * time.Millisecond):
			panic("overslept")
		case <-ctx.Done():
			// Done channel is closed when the deadline expires(times out)
			fmt.Println("Done 2:", ctx)
			fmt.Println("Done 2:", ctx.Err()) // prints "context canceled"
		}
	}()
	/*
	   Done 2: context.Background.WithDeadline(2015-09-02 22:38:00.841406346 -0700 PDT [99.990302ms]).WithValue(2, "Linux")
	   Done 2: context canceled
	*/

	fmt.Println()
	func() {
		// Done channel is closed when the deadline expires(times out)
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
		go sendRequest(ctx, time.Second)
		defer cancel()
		select {
		case <-time.After(200 * time.Millisecond):
			panic("overslept")
		case <-ctx.Done():
			// Done channel is closed when the deadline expires(times out)
			fmt.Println("Done 3 timed out:", ctx)
			fmt.Println("Done 3 timed out:", ctx.Err()) // prints "context deadline exceeded"
		}
	}()
	time.Sleep(time.Second)
	/*
	   Started: sendRequest
	   Done 3 timed out: context.Background.WithDeadline(2015-09-02 22:38:00.841432965 -0700 PDT [-105.376µs])
	   Done 3 timed out: context deadline exceeded
	   Done: sendRequest
	*/

	fmt.Println()
	func() {
		// Done channel is closed when the deadline expires(times out)
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
		go sendRequest(ctx, time.Millisecond)
		defer cancel()
		select {
		case <-ctx.Done():
			// Done channel is closed when the deadline expires(times out)
			fmt.Println("Done 4 timed out:", ctx)
			fmt.Println("Done 4 timed out:", ctx.Err()) // prints "context deadline exceeded"
		}
	}()
	/*
		Started: sendRequest
		Done: sendRequest
		Done 4 timed out: context.Background.WithDeadline(2015-09-02 23:41:38.874000873 -0700 PDT [-139.856µs])
		Done 4 timed out: context deadline exceeded
	*/

	fmt.Println()
	func() {
		// Done channel is closed when the deadline expires(times out)
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
		defer cancel()
		sendRequestCorrectWay(ctx, 500*time.Millisecond)
		fmt.Println("Done 5")
	}()
	/*
		Started: sendRequestCorrectWay
		Timed out: sendRequestCorrectWay
		Done 5
	*/

	fmt.Println()
	func() {
		// Done channel is closed when the deadline expires(times out)
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
		defer cancel()
		sendRequestCorrectWay(ctx, time.Millisecond)
		fmt.Println("Done 6")
	}()
	/*
		Started: sendRequestCorrectWay
		Done: sendRequestCorrectWay
		Done 6
	*/

	fmt.Println()
	func() {
		// Done channel is closed when the deadline expires(times out)
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
		cancel()
		sendRequestCorrectWay(ctx, time.Millisecond)
		fmt.Println("Done 7")
	}()
	/*
		Started: sendRequestCorrectWay
		Timed out: sendRequestCorrectWay
		Done 7
	*/
}

func sendRequest(ctx context.Context, duration time.Duration) {
	fmt.Println("Started: sendRequest")
	time.Sleep(duration)
	fmt.Println("Done: sendRequest")
}

func sendRequestCorrectWay(ctx context.Context, duration time.Duration) {
	fmt.Println("Started: sendRequestCorrectWay")
	select {
	case <-time.After(duration):
		fmt.Println("Done: sendRequestCorrectWay")
		return
	case <-ctx.Done():
		fmt.Println("Timed out: sendRequestCorrectWay")
		return
	}
}
