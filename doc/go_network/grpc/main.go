package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
)

type transporterServer struct { // satisfy TransporterServer
	received string
}

func (t *transporterServer) Transfer(ctx context.Context, r *Request) (*Response, error) {
	fmt.Printf("transporterServer.Transfer has received Request: %+v\n", r)

	t.received = fmt.Sprintf("Data: %v at %v", r.Data, time.Now())
	// return &Response{Result: t.received}, fmt.Errorf("error from Transfer")
	return &Response{Result: t.received}, nil
}

var (
	port     = ":2378"
	endpoint = "localhost" + port
)

func server() {
	fmt.Println("gRPC server has started at", endpoint)

	var (
		grpcServer = grpc.NewServer()
		sender     = &transporterServer{}
	)
	ln, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	RegisterTransporterServer(grpcServer, sender)

	go func() {
		if err := grpcServer.Serve(ln); err != nil {
			log.Fatal(err)
		}
	}()
}

func client() {
	fmt.Println("gRPC client has started at", endpoint)

	conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	cli := NewTransporterClient(conn)
	resp, err := cli.Transfer(context.Background(), &Request{Data: "hello"})
	if err != nil {
		fmt.Printf("client.Transfer error %v\n", err)
	}
	fmt.Printf("client.Transfer Response: %+v\n", resp)
}

func main() {
	server()

	time.Sleep(3 * time.Second)

	client()
}

/*
gRPC server has started at localhost:2378
gRPC client has started at localhost:2378
transporterServer.Transfer has received Request: data:"hello"
client.Transfer Response: result:"Data: hello at 2016-03-12 02:35:25.977548786 -0800 PST"
*/
