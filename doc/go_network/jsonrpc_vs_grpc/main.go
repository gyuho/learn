package main

import (
	"flag"
	"log"
	"math/rand"
	"time"

	"github.com/gyuho/learn/doc/go_network/jsonrpc_vs_grpc/demogrpc"
	"github.com/gyuho/learn/doc/go_network/jsonrpc_vs_grpc/demojsonrpc"
)

/*
go run main.go -opt="grpc"
go run main.go -opt="jsonrpc"

go run main.go -opt="grpc"
2015/12/14 00:29:05 Size chosen: 100000
2015/12/14 00:29:05 Option chosen: grpc
2015/12/14 00:29:08 Done with generating random data...
2015/12/14 00:29:08 GRPC on :8080
2015/12/14 00:29:20 GRPC Took 11.96749331s for 100000 calls (119.674µs per each).

go run main.go -opt="jsonrpc"
2015/12/14 00:29:32 Size chosen: 100000
2015/12/14 00:29:32 Option chosen: jsonrpc
2015/12/14 00:29:35 Done with generating random data...
2015/12/14 00:29:35 JSONRPC on :8080
2015/12/14 00:32:13 JSONRPC Took 2m38.121612447s for 100000 calls (1.581216ms per each).
*/

func main() {
	switch opt {

	case "grpc":
		demogrpc.Stress(port, endpoint, keys, vals, numConns, numClients)

	case "jsonrpc":
		demojsonrpc.Stress(port, endpoint, keys, vals)

	}
}

var (
	port     = ":8080"
	endpoint = "localhost" + port

	numConns   = 1
	numClients = 1
	// numClients = 100

	size = 100000
	opt  = "grpc"

	keys = make([][]byte, size)
	vals = make([][]byte, size)
)

func init() {
	sizePt := flag.Int(
		"size",
		100000,
		"Size of keys to put",
	)
	optPt := flag.String(
		"opt",
		"grpc",
		"'grpc' or 'jsonrpc'",
	)
	flag.Parse()

	size = *sizePt
	opt = *optPt
	if opt != "grpc" && opt != "jsonrpc" {
		log.Fatalf("%s is unknown\n", opt)
	}
	log.Println("Size chosen:", size)
	log.Println("Option chosen:", opt)

	keys = make([][]byte, size)
	vals = make([][]byte, size)
	for i := range keys {
		keys[i] = randBytes(100)
		vals[i] = randBytes(100)
	}
	log.Println("Done with generating random data...")
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
