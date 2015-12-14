package demogrpc

import (
	"log"
	"net"
	"sync"
	"time"

	pb "github.com/gyuho/learn/doc/go_network/jsonrpc_vs_grpc/messagepb"
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
	log.Println("GRPC on", port)

	ln, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	// defer ln.Close()

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

func Stress(port, endpoint string, keys, vals [][]byte, numConns, numClients int) {

	go startServerGRPC(port)

	conns := make([]*grpc.ClientConn, numConns)
	for i := range conns {
		conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
		if err != nil {
			panic(err)
		}
		conns[i] = conn
	}
	clients := make([]pb.KVClient, numClients)
	for i := range clients {
		clients[i] = pb.NewKVClient(conns[i%int(numConns)])
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

	tt := time.Since(st)
	size := len(keys)
	pt := tt / time.Duration(size)
	log.Printf("GRPC took %v for %d calls with %d client(s) (%v per each).\n", tt, size, numClients, pt)
}
