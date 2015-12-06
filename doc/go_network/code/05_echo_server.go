package main

import (
	"fmt"
	"net"
)

func startServer(port string) {
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
		go handleRequests(conn)
	}
}

// Handles incoming requests.
func handleRequests(conn net.Conn) {
	fmt.Printf("Received from %s → %s\n", conn.RemoteAddr(), conn.LocalAddr())
	buf := make([]byte, 5) // read max 5 characters
	if _, err := conn.Read(buf); err != nil {
		panic(err)
	}
	conn.Write([]byte("received message: " + string(buf) + "\n"))
	conn.Close()
}

func main() {
	const port = ":8080"
	startServer(port)
}

/*
$ echo "Hello server" | nc localhost 8080

Received from 127.0.0.1:58405 → 127.0.0.1:8080
Received from 127.0.0.1:58406 → 127.0.0.1:8080
Received from 127.0.0.1:58407 → 127.0.0.1:8080
Received from 127.0.0.1:58408 → 127.0.0.1:8080
Received from 127.0.0.1:58409 → 127.0.0.1:8080
...
*/
