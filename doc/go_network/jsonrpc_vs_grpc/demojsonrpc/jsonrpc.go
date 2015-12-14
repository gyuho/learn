package demojsonrpc

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"sync"
	"time"
)

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

/*
Register publishes in the server the set of methods of the receiver value that satisfy the following conditions:

- exported method of exported type
- two arguments, both of exported type
- the second argument is a pointer
- one return value, of type error
*/
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
	log.Println("JSONRPC on", port)

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
