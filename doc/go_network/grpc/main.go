package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
)

type Data struct {
	Operation string
}

func (d *Data) Send(ctx context.Context, r *Request) (*Response, error) {
	d.Operation = fmt.Sprintf("Operation: %v | ID: %v | %v", r.Operation, r.Id, time.Now())
	return &Response{
		Result: d.Operation,
	}, nil
}

var (
	port     = ":2378"
	endpoint = "localhost" + port
	gs       = grpc.NewServer()
	data     = &Data{}
)

func main() {
	log.Println("net.Listen:", endpoint)
	ln, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	RegisterSenderServer(gs, data)
	go func() {
		if err := gs.Serve(ln); err != nil {
			log.Fatal(err)
		}
	}()

	log.Println("grpc.Dial:", endpoint)
	conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	cli := NewSenderClient(conn)
	req := &Request{Operation: Request_Start, Id: "hello"}
	log.Println("cli.Send:", req)
	resp, err := cli.Send(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()
	fmt.Printf("resp: %+v\n", resp)
	fmt.Printf("data: %+v\n", data)
}
