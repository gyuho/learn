package main

import (
	"testing"
	"time"

	"golang.org/x/net/context"
)

func TestDatabase(t *testing.T) {
	isInVPC := false
	timeout := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	_, err := getConn(ctx, isInVPC, "xx")
	if err != nil {
		panic(err)
	}
}
