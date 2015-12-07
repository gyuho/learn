package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"strings"
	"time"
)

var (
	port     = ":8080"
	endpoint = "localhost" + port
	msg      = MessageType{
		Contents: "    hello    world!   ",
	}

	callSize = 15000
)

type MessageType struct {
	Contents string
}

func main() {

	go startServerJSONRPC(port)
	sj := time.Now()
	for i := 0; i < callSize; i++ {
		clientJSONRPC(endpoint, msg)
	}
	fmt.Printf("clientJSONRPC took %v for %d calls.\n", time.Since(sj), callSize)

}

func (r *MessageType) MyMethod(msg MessageType, resp *MessageType) error {
	resp.Contents = strings.Join(strings.Fields(strings.TrimSpace(msg.Contents)), " ")
	return nil
}

func startServerJSONRPC(port string) {
	log.Println("RPC on", port)

	srv := rpc.NewServer()
	s := new(MessageType)
	srv.Register(s)
	srv.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)

	// Listen function creates servers,
	// listening for incoming connections.
	ln, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	for {
		// Listen for an incoming connection.
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		go srv.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}

func clientJSONRPC(endpoint string, msg MessageType) {
	conn, err := net.Dial("tcp", endpoint)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := jsonrpc.NewClient(conn)

	resp := &MessageType{}

	if err := client.Call("MessageType.MyMethod", msg, resp); err != nil {
		panic(err)
	}

	if resp.Contents != "hello world!" {
		log.Fatalf("rpc failed with %v", resp.Contents)
	}
}
