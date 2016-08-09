package main

import (
	"fmt"
	"net"
	"os"
	"reflect"
)

const basePort = 20000

func main() {
	addr := fmt.Sprintf("127.0.0.1:%d.%d.sock", basePort, os.Getpid())
	fmt.Println("starting listening", addr)

	ln, err := NewUnixListener(addr)
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	donec := make(chan struct{})
	go func() {
		defer close(donec)

		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Println("after ln.Accept:", reflect.TypeOf(conn), conn.RemoteAddr())

		b := make([]byte, 15)
		_, err = conn.Read(b)
		fmt.Printf("after conn.Read: %q\n", b)
		// this will be "hello" if make([]byte, 5)

		err = conn.Close()
		if err != nil {
			panic(err)
		}
	}()

	fmt.Println("starting dialing", addr)
	conn, err := net.Dial("unix", addr)
	if err != nil {
		panic(err)
	}
	fmt.Println("after net.Dial:", reflect.TypeOf(conn), conn.RemoteAddr())

	_, err = conn.Write([]byte("hello world!"))
	if err != nil {
		panic(err)
	}
	fmt.Println("after conn.Write")

	err = conn.Close()
	if err != nil {
		panic(err)
	}

	<-donec
	fmt.Println("Closed", ln.Addr())
}

/*
starting listening 127.0.0.1:20000.13386.sock
starting dialing 127.0.0.1:20000.13386.sock
after net.Dial: *net.UnixConn 127.0.0.1:20000.13386.sock
after conn.Write
after ln.Accept: *net.UnixConn @
after conn.Read: "hello world!\x00\x00\x00"
Closed 127.0.0.1:20000.13386.sock
*/

type unixListener struct{ net.Listener }

func NewUnixListener(addr string) (net.Listener, error) {
	if err := os.RemoveAll(addr); err != nil {
		return nil, err
	}
	l, err := net.Listen("unix", addr)
	if err != nil {
		return nil, err
	}
	return &unixListener{l}, nil
}

func (ul *unixListener) Close() error {
	if err := os.RemoveAll(ul.Addr().String()); err != nil {
		return err
	}
	return ul.Listener.Close()
}
