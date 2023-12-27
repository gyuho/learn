Originally published at https://blog.gopheracademy.com/advent-2015/etcd-distributed-key-value-store-with-grpc-http2/

<br>

## What is etcd?

<img alt="etcd"
     src="https://raw.githubusercontent.com/coreos/etcd/master/logos/etcd-horizontal-color.png"
     style="float:right; margin-bottom:2em;"/>

etcd is a distributed, consistent key-value store, written in Go. Similar to how Linux distributions typically use `/etc` to store local configuration data, etcd can be thought of as a reliable store for *distributed* configuration data. It is distributed by replicating data to multiple machines, therefore highly available against single point of failures. Using the Raft consensus algorithms, etcd gracefully handles network partitions and machine failures, even leader failures. etcd is being widely used [in production](https://github.com/coreos/etcd/blob/master/Documentation/production-users.md): CoreOS, Kubernetes, vulcand, etc.

## How does etcd work?

etcd clusters are based on a strong leader. A leader is elected by other members in cluster. Once elected, the leader starts processing client requests and replicating them to its followers. All server-to-server communication is done by RPC (Remote Procedure Call).

## What matters to etcd performance?

Latency matters. Data should be delivered as fast as possible. Lower latency means higher throughput for the state machine, consequently consistent data replication. Memory/CPU usage matters. etcd needs to replicate messages with minimum resources. gRPC helps reduce such costs and minimize latency.

## What is gRPC?

gRPC is a framework from Google, to handle remote procedure calls. It uses HTTP/2 to support highly performant, scalable APIs and microservices. HTTP/2 can be better than HTTP/1.x in that:

<ul>
	<li>it is binary rather than textual, therefore more compact and efficient.</li>
	<li>it multiplexes requests over a single TCP connection. This allows many messages to be in flight at the same time and reduces network resource usage.</li>
	<li>it uses header compression to reduce the size of requests and responses.</li>
</ul>

<br>

A `frame` is the basic HTTP/2 protocol unit: HTTP/2 splits requests/responses into binary `frames` before sending them over TCP connections. And a `stream` is a bidirectional flow of `frames`, which share a same stream id. HTTP/2 makes `streams` independent to each other, so that one single HTTP/2 connection can have multiple concurrently open streams, and process `frames` from multiple `streams` asynchronously. And by default gRPC uses protocol buffers to exchange messages. Google developed Protocol Buffers for serializing structured data. Protocol buffers are encoded in binary format, therefore more compact efficient than JSON.

## gRPC vs JSON RPC

gRPC uses HTTP/2 and protocol buffers, so it is expected to read and write faster. To benchmark them, let's set up simple RPC servers to store key/value pairs. Here's how to do it with `jsonrpc` package:

```go
type PutRequest struct {
	Key   []byte
	Value []byte
}

type ResponseHeader struct {
	Exist bool
	Value []byte
}

type PutResponse struct {
	Header *ResponseHeader
}

type KVStoreJSONRPC struct {
	mu    sync.Mutex
	store map[string][]byte
}

func (s *KVStoreJSONRPC) Put(r PutRequest, resp *PutResponse) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	resp.Header = &ResponseHeader{}
	if v, ok := s.store[string(r.Key)]; ok {
		resp.Header.Exist = true
		resp.Header.Value = v
	} else {
		s.store[string(r.Key)] = r.Value
	}
	return nil
}

func startServerJSONRPC(port string) {
	s := new(KVStoreJSONRPC)
	s.store = make(map[string][]byte)

	srv := rpc.NewServer()
	srv.Register(s)
	srv.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)

	ln, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		go srv.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}

func clientJSONRPC(endpoint string, msg PutRequest) {
	conn, err := net.Dial("tcp", endpoint)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := jsonrpc.NewClient(conn)

	resp := &PutResponse{}
	if err := client.Call("KVStoreJSONRPC.Put", msg, resp); err != nil {
		panic(err)
	}
}

func Stress(port, endpoint string, keys, vals [][]byte) {
	go startServerJSONRPC(port)

	st := time.Now()

	for i := range keys {
		msg := PutRequest{
			Key:   keys[i],
			Value: vals[i],
		}
		clientJSONRPC(endpoint, msg)
	}

	tt := time.Since(st)
	size := len(keys)
	pt := tt / time.Duration(size)
	log.Printf("JSONRPC took %v for %d requests with 1 client(s) (%v per each).\n", tt, size, pt)
}

```

Here's how to use gRPC with Go. First define service in `*.proto` format, as below:

```proto
syntax = "proto3";
package messagepb;

import "github.com/gogo/protobuf/gogoproto/gogo.proto";

option (gogoproto.marshaler_all) = true;
option (gogoproto.sizer_all) = true;
option (gogoproto.unmarshaler_all) = true;
option (gogoproto.goproto_getters_all) = false;

service KV {
  // Put puts the given key into the store.
  // A put request increases the revision of the store,
  // and generates one event in the event history.
  rpc Put(PutRequest) returns (PutResponse) {}
}

message PutRequest {
  bytes key = 1;
  bytes value = 2;
}

message ResponseHeader {
  bool exist = 1;
  bytes value = 2;
}

message PutResponse {
  ResponseHeader header = 1;
}

```

And generate Go code using `protoc`:

```bash
go get -v -u github.com/gogo/protobuf/{proto,protoc-gen-gogo,gogoproto,protoc-gen-gofast};

protoc \
	--gofast_out=plugins=grpc:. \
	--proto_path=$GOPATH/src:$GOPATH/src/github.com/gogo/protobuf/protobuf:. \
	*.proto;

```

This generates `*.pb.go` file that can be imported as a package. Let's assume that it is generated under `messagepb` package. Then now implement gRPC server/client as follow:

```go
import (
	"log"
	"net"
	"sync"
	"time"

	pb "YOUR_PATH/messagepb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type KVStoreGRPC struct {
	mu    sync.Mutex
	store map[string][]byte
}

func (s *KVStoreGRPC) Put(ctx context.Context, r *pb.PutRequest) (*pb.PutResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	resp := &pb.PutResponse{}
	resp.Header = &pb.ResponseHeader{}
	if v, ok := s.store[string(r.Key)]; ok {
		resp.Header.Exist = true
		resp.Header.Value = v
	} else {
		s.store[string(r.Key)] = r.Value
	}
	return resp, nil
}

func startServerGRPC(port string) {
	ln, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}

	s := &KVStoreGRPC{}
	s.store = make(map[string][]byte)

	grpcServer := grpc.NewServer()
	pb.RegisterKVServer(grpcServer, s)

	go func() {
		if err := grpcServer.Serve(ln); err != nil {
			panic(err)
		}
	}()
}

func Stress(port, endpoint string, keys, vals [][]byte, connsN, clientsN int) {

	go startServerGRPC(port)

	conns := make([]*grpc.ClientConn, connsN)
	for i := range conns {
		conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
		if err != nil {
			panic(err)
		}
		conns[i] = conn
	}
	clients := make([]pb.KVClient, clientsN)
	for i := range clients {
		clients[i] = pb.NewKVClient(conns[i%int(connsN)])
	}

	requests := make(chan *pb.PutRequest, len(keys))
	done, errChan := make(chan struct{}), make(chan error)

	for i := range clients {
		go func(i int, requests chan *pb.PutRequest) {
			for r := range requests {
				if _, err := clients[i].Put(context.Background(), r); err != nil {
					errChan <- err
					return
				}
			}
			done <- struct{}{}
		}(i, requests)
	}

	st := time.Now()

	for i := range keys {
		r := &pb.PutRequest{
			Key:   keys[i],
			Value: vals[i],
		}
		requests <- r
	}

	close(requests)

	cn := 0
	for cn != len(clients) {
		select {
		case err := <-errChan:
			panic(err)
		case <-done:
			cn++
		}
	}
	close(done)
	close(errChan)

	tt := time.Since(st)
	size := len(keys)
	pt := tt / time.Duration(size)
	log.Printf("GRPC took %v for %d requests with %d client(s) (%v per each).\n", tt, size, clientsN, pt)
}

```

## Benchmark results

Tests send 300,000 requests to key/value stores. One with `jsonrpc`, the other with `gRPC`. Both `jsonrpc` and `gRPC` code use only one TCP connection. And another `gRPC` case with one TCP connection but with multiple clients:

|RPC|# of requests|# of clients|total time|per-request time|
|---|---|---|---|:-:|
|`jsonrpc`|300,000|1|8m7.270s|1.624ms|
|gRPC|300,000|1|36.715s|122.383µs|
|gRPC|300,000|100|7.167s|23.892µs|

<br>

And if compared on memory usage:

|RPC|`jsonrpc`|gRPC|delta|
|---|---------|----|:---:|
|NsPerOp|487271046903|36716116701|-92.46%|
|AllocsPerOp|32747687|25221256|-22.98%|
|AllocedBytesPerOp|3182814152|1795122672|-43.60%|

|RPC|`jsonrpc`|gRPC with 100 clients|delta|
|---|---------|----|:---:|
|NsPerOp|487271046903|7168591678|-98.53%|
|AllocsPerOp|32747687|25230286|-22.96%|
|AllocedBytesPerOp|3182814152|1795831944|-43.58%|

|RPC|gRPC|gRPC with 100 clients|delta|
|---|---------|----|:---:|
|NsPerOp|36716116701|7168591678|-80.48%|
|AllocsPerOp|25221256|25230286|+0.04%|
|AllocedBytesPerOp|1795122672|1795831944|+0.04%|

<br>

As you see, gRPC is much faster and lighter than `jsonrpc`. Not only performant, but also gRPC is easier to reason about concurrency. HTTP/1.x requires multiple TCP connections for concurrent requests, while HTTP/2 can have multiple requests over one single TCP connection and still process them asynchronously. Tests above show that gRPC with multiple clients solely speeds up gRPC by 80%, without opening multiple TCP connections.

## How does etcd use gRPC?

etcd already uses protocol buffers for all structured data and supports gRPC as an experimental feature. etcd team is actively working on gRPC integration to replace all etcd RPCs. Here's how different parts of the etcd codebase leverage gRPC and Protocol Buffers:

```
etcdserver/         // Each etcd member runs as an etcdserver.

	etcdserverpb/   // All message, schema, service are defined
	                // here in Protocol Buffer format.

	api/
		v3rpc/      // Implements etcd v3 RPC system based on
					// gRPC and Protocol Buffers.

main.go             // Start 'etcdserver' with grpc framework.

```

## etcd 2.0 vs 3+

- [etcd v2.0](https://github.com/coreos/etcd/tree/release-2.0) uses Protocol Buffer but does not support gRPC.
- [etcd v3+](https://github.com/coreos/etcd) uses gRPC with Protocol Buffer for v3 API.

Here's benchmark test set-up:

- Both use only one TCP connection.
- 100 Put requests to 3-member etcd cluster.
- Each PUT request contains 100 and 750 random byte slices.
- etcd 3+ also uses a single TCP connection but multiple clients.

etcd 2 client can be implemented as below:

```go
import (
	"github.com/coreos/go-etcd/etcd"
)

var (
	connsN   = 1
	clientsN = 100

	stressN    = 100
	stressKeyN = 100
	stressValN = 750
)

func put() {
	client := etcd.NewClient(machines)
	for i := 0; i < stressN; i++ {
		if _, err := client.Set(keys[i], vals[i], 0); err != nil {
			log.Fatal(err)
		}
	}
}

```

etcd 3+ client can be as below:

```go
import (
	"github.com/coreos/etcd/Godeps/_workspace/src/golang.org/x/net/context"
	"github.com/coreos/etcd/Godeps/_workspace/src/google.golang.org/grpc"
	"github.com/coreos/etcd/etcdserver/etcdserverpb"
)

var (
	connsN   = 1
	clientsN = 100

	stressN    = 100
	stressKeyN = 100
	stressValN = 750
)

func put() {
	conns := make([]*grpc.ClientConn, connsN)
	for i := range conns {
		conns[i] = mustCreateConn(endpoint)
	}
	clients := make([]etcdserverpb.KVClient, clientsN)
	for i := range clients {
		clients[i] = etcdserverpb.NewKVClient(conns[i%int(connsN)])
	}

	requests := make(chan *etcdserverpb.PutRequest, stressN)
	done, errChan := make(chan struct{}), make(chan error)

	for i := range clients {
		go func(i int, requests <-chan *etcdserverpb.PutRequest) {
			for r := range requests {
				if _, err := clients[i].Put(context.Background(), r); err != nil {
					errChan <- err
					return
				}
			}
			done <- struct{}{}
		}(i, requests)
	}

	for i := 0; i < stressN; i++ {
		r := &etcdserverpb.PutRequest{
			Key:   keys[i],
			Value: vals[i],
		}
		requests <- r
	}

	close(requests)

	cn := 0
	for cn != len(clients) {
		select {
		case err := <-errChan:
			return err
		case <-done:
			cn++
		}
	}
}

```

Here's performance benchmark result:

- Each PUT request of etcd 2.0 took about 292.30919ms.
- Each PUT request of etcd 3+ took about 6.780163ms.

|RPC|etcd 2.0|etcd 3+|delta|
|---|---------|----|:---:|
|NsPerOp|24004505139|0|-100.00%|
|AllocsPerOp|15182|0|-100.00%|
|AllocedBytesPerOp|2601720|0|-100.00%|

Without consuming any extra TCP connections, etcd 3+ processes client requests concurrently, therefore much faster. This test might bias towards gRPC, because you could get similar performance using multiple TCP connections with HTTP/1.x. But then you have to worry about the file descriptor limits (in Linux), and multiple TCP connections can cause network congestion.

## Interested?

etcd is still in active development, while widely used in production. If interested, please try it out:

- [github.com/coreos/etcd](https://github.com/coreos/etcd/releases/tag/v2.2.2)
- [gRPC](http://www.grpc.io)
- [HTTP/2](https://http2.github.io)
- [Protocol Buffers](https://developers.google.com/protocol-buffers)

