package main

import (
	"math/rand"
	"time"

	"github.com/gyuho/learn/doc/go_network/jsonrpc_vs_grpc/demogrpc"
)

var (
	port     = ":5000"
	endpoint = "localhost" + port

	callSize = 10000
	keys     = make([][]byte, callSize)
	vals     = make([][]byte, callSize)
)

func init() {
	for i := range keys {
		keys[i] = randBytes(100)
		vals[i] = randBytes(100)
	}
}

func main() {
	demogrpc.Run(port, endpoint, keys, vals, 1, 1)
	// demojsonrpc.Run(port, endpoint, keys, vals)
}

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func randBytes(n int) []byte {
	src := rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return b
}
