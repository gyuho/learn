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
go run main.go -opt="jsonrpc" -size=300000
2015/12/14 02:52:03 Size chosen: 300000
2015/12/14 02:52:03 Option chosen: jsonrpc
2015/12/14 02:52:12 Done with generating random data...
2015/12/14 02:52:12 JSONRPC on :8080
2015/12/14 03:00:19 JSONRPC took 8m7.104095469s for 300000 requests (1.62368ms per each).

go run main.go -opt="grpc" -size=300000
2015/12/14 03:00:22 Size chosen: 300000
2015/12/14 03:00:22 Option chosen: grpc
2015/12/14 03:00:31 Done with generating random data...
2015/12/14 03:00:31 GRPC on :8080
2015/12/14 03:01:07 GRPC took 36.38581904s for 300000 requests (121.286µs per each).

go run main.go -opt="grpc" -size=300000 -numc=100
2015/12/14 03:38:50 Size chosen: 300000
2015/12/14 03:38:50 Option chosen: grpc
2015/12/14 03:38:59 Done with generating random data...
2015/12/14 03:38:59 GRPC on :8080
2015/12/14 03:39:07 GRPC took 7.907762062s for 300000 requests with 100 client(s) (26.359µs per each).
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
		"Size of keys to put.",
	)
	numClienstsPt := flag.Int(
		"numc",
		1,
		"Size of clients to run.",
	)
	optPt := flag.String(
		"opt",
		"grpc",
		"'grpc' or 'jsonrpc'.",
	)
	flag.Parse()

	size = *sizePt
	numClients = *numClienstsPt
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

func randBytes(n int) []byte {
	const (
		letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		letterIdxBits = 6                    // 6 bits to represent a letter index
		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
		letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	)
	src := rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
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
