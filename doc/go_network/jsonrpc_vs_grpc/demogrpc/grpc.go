package demogrpc

import (
	"fmt"
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
	fmt.Println("Got", len(s.store))
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

func Run(port, endpoint string, keys, vals [][]byte, totalConns, totalClients int) {

	go startServerGRPC(port)

	conns := make([]*grpc.ClientConn, totalConns)
	for i := range conns {
		conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
		if err != nil {
			panic(err)
		}
		conns[i] = conn
	}

	clients := make([]pb.KVClient, totalClients)
	for i := range clients {
		clients[i] = pb.NewKVClient(conns[i%int(totalConns)])
	}

	// 	conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	defer conn.Close()

	requests := make(chan *pb.PutRequest, len(keys))

	var wg sync.WaitGroup

	for i := range clients {
		wg.Add(1)
		go func(i int, requests chan *pb.PutRequest) {
			defer wg.Done()
			for r := range requests {
				if _, err := clients[i].Put(context.Background(), r); err != nil {
					panic(err)
				}
			}
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

	wg.Wait()

	fmt.Printf("clientGRPC took %v for %d calls.\n", time.Since(st), len(keys))
}
