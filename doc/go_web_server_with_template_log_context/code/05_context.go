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
		defer cancel()
		sendRequestWithWaitTime(ctx, 500*time.Millisecond)
		fmt.Println("Done 3")
	}()
	/*
		Started: sendRequestWithWaitTime
		Timed out: sendRequestWithWaitTime
		context deadline exceeded
		Done 3
	*/

	fmt.Println()
	func() {
		// Done channel is closed when the deadline expires(times out)
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
		defer cancel()
		sendRequestWithWaitTime(ctx, time.Millisecond)
		fmt.Println("Done 4")
	}()
	/*
		Started: sendRequestWithWaitTime
		wait is Done: sendRequestWithWaitTime
		Done 4
	*/

	fmt.Println()
	func() {
		// Done channel is closed when the deadline expires(times out)
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
		cancel()
		sendRequestWithWaitTime(ctx, time.Millisecond)
		fmt.Println("Done 5")
	}()
	/*
		Started: sendRequestWithWaitTime
		Timed out: sendRequestWithWaitTime
		context canceled
		Done 5
	*/
}

func sendRequestWithWaitTime(ctx context.Context, wait time.Duration) {
	fmt.Println("Started: sendRequestWithWaitTime")
	select {
	case <-time.After(wait):
		fmt.Println("wait is Done: sendRequestWithWaitTime")
		return
	case <-ctx.Done():
		// Done channel is closed when the deadline expires(times out)
		fmt.Println("Timed out: sendRequestWithWaitTime")
		fmt.Println(ctx.Err())
		return
	}
}
