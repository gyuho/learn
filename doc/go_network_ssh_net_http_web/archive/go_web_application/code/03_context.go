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
		timeout := 100 * time.Millisecond
		processingTime := time.Nanosecond
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		cancel()
		send(ctx, processingTime)
		fmt.Println("Done 2")
	}()
	/*
		send Timeout: context canceled
		Done 2
	*/

	fmt.Println()
	func() {
		timeout := 100 * time.Millisecond
		processingTime := time.Nanosecond
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		send(ctx, processingTime)
		fmt.Println("Done 3")
	}()
	/*
		send Done!
		Done 3
	*/

	fmt.Println()
	func() {
		timeout := 100 * time.Millisecond
		processingTime := time.Minute
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		send(ctx, processingTime)
		fmt.Println("Done 4")
	}()
	/*
		send Timeout: context deadline exceeded
		Done 4
	*/
}

func send(ctx context.Context, processingTime time.Duration) {
	done := make(chan struct{})
	go func() {
		time.Sleep(processingTime)
		done <- struct{}{}
	}()
	select {
	case <-done:
		fmt.Println("send Done!")
		return
	case <-ctx.Done():
		// Done channel is closed when the deadline expires(times out)
		// or canceled.
		fmt.Println("send Timeout:", ctx.Err())
		return
	}
}
