package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
)

func main() {
	buf := bytes.NewBuffer(nil)
	sizeN := uint64(3)
	err := binary.Write(buf, binary.BigEndian, sizeN)
	if err != nil {
		log.Fatal(err)
	}
	n, err := buf.Write(bytes.Repeat([]byte("a"), 5))
	fmt.Println(n, err)                                 // 5 <nil>
	fmt.Printf("%q (%s)\n", buf.String(), buf.String()) // "\x00\x00\x00\x00\x00\x00\x00\x03aaaaa" (aaaaa)

	var num uint64
	err = binary.Read(buf, binary.BigEndian, &num)
	fmt.Println(err, num) // <nil> 3

	src := make([]byte, int(num))
	_, err = io.ReadFull(buf, src)
	fmt.Println(err, string(src)) // <nil> aaa
}
