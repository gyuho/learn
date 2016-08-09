package main

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

/*
starting listening from 127.0.0.1:20000.27885.sock
finished writeWithDial
#0: conn.Read: "\x00\x00\x00\x00\x00"
#1: conn.Read: "A\x00\x00\x00\x00"
#2: conn.Read: "AA\x00\x00\x00"
#3: conn.Read: "AAA\x00\x00"
#4: conn.Read: "AAAA\x00"
*/

func main() {
	addr := fmt.Sprintf("127.0.0.1:20000.%d.sock", os.Getpid())
	filterFunc := func(c net.Conn) bool { return strings.Contains(c.LocalAddr().String(), "127.0.0.1") }
	ln, err := NewUnixListenerWithFilter(addr, filterFunc)
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	donec := make(chan struct{})
	go writeWithDial("unix", addr, donec)

	fmt.Println("starting listening from", ln.Addr())
	for i := 0; i < 5; i++ {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		b := make([]byte, 5)
		_, err = conn.Read(b)
		fmt.Printf("#%d: conn.Read: %q\n", i, b)

		err = conn.Close()
		if err != nil {
			panic(err)
		}
	}
	<-donec
}

type unixListener struct{ net.Listener }

type unixListenerWithFilter struct {
	ln         net.Listener
	filterFunc func(net.Conn) bool
}

func NewUnixListenerWithFilter(addr string, filterFunc func(net.Conn) bool) (net.Listener, error) {
	if err := os.RemoveAll(addr); err != nil {
		return nil, err
	}
	l, err := net.Listen("unix", addr)
	if err != nil {
		return nil, err
	}
	return &unixListenerWithFilter{ln: l, filterFunc: filterFunc}, nil
}

func (ul *unixListenerWithFilter) Accept() (net.Conn, error) {
	conn, err := ul.ln.Accept()
	if !ul.filterFunc(conn) {
		conn.(*net.UnixConn).SetDeadline(time.Now())
	}
	return conn, err
}

func (ul *unixListenerWithFilter) Close() error {
	if err := os.RemoveAll(ul.ln.Addr().String()); err != nil {
		return err
	}
	return ul.ln.Close()
}

func (ul *unixListenerWithFilter) Addr() net.Addr {
	return ul.ln.Addr()
}

func writeWithDial(proto, addr string, donec chan struct{}) {
	defer close(donec)

	for i := 0; i < 5; i++ {
		conn, err := net.Dial(proto, addr)
		if err != nil {
			panic(err)
		}

		b := bytes.Repeat([]byte("A"), i)
		_, err = conn.Write(b)
		if err != nil {
			panic(err)
		}

		err = conn.Close()
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("finished writeWithDial")
}
