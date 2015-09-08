package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Args struct {
	A, B int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func startServer(port string) {
	srv := rpc.NewServer()
	arith := new(Arith)
	srv.Register(arith)
	srv.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)
	// Listen function creates servers,
	// listening for incoming connections.
	ln, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	defer ln.Close()
	fmt.Println("Listening on", port)
	for {
		// Listen for an incoming connection.
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		go srv.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}

func main() {
	const port = ":5000"
	go startServer(port)

	conn, err := net.Dial("tcp", "localhost"+port)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	args := &Args{5, 10}
	var reply int

	client := jsonrpc.NewClient(conn)
	if err := client.Call("Arith.Multiply", args, &reply); err != nil {
		panic(err)
	}
	fmt.Println("reply:", reply)
	// reply: 50
}
