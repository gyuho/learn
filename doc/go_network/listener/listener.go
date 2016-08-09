package main

import (
	"bytes"
	"fmt"
	"net"
)

/*
go run listener.go
starting listening from [::]:8080
#0: conn.Read: "\x00\x00\x00\x00\x00" from remote "[::1]:33500"
#1: conn.Read: "A\x00\x00\x00\x00" from remote "[::1]:33502"
#2: conn.Read: "AA\x00\x00\x00" from remote "[::1]:33504"
#3: conn.Read: "AAA\x00\x00" from remote "[::1]:33506"
#4: conn.Read: "AAAA\x00" from remote "[::1]:33508"
finished writeWithDial
*/

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	donec := make(chan struct{})
	go writeWithDial("tcp", ln.Addr().String(), donec)

	fmt.Println("starting listening from", ln.Addr())
	for i := 0; i < 5; i++ {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		b := make([]byte, 5)
		_, err = conn.Read(b)
		fmt.Printf("#%d: conn.Read: %q from remote %q\n", i, b, conn.RemoteAddr().String())

		err = conn.Close()
		if err != nil {
			panic(err)
		}
	}
	<-donec
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
