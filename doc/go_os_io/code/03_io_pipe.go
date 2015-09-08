package main

import (
	"fmt"
	"io"
)

func main() {
	done := make(chan struct{})
	r, w := io.Pipe()

	go func() {
		data := []byte("Hello World!")
		n, err := w.Write(data)
		if err != nil {
			panic(err)
		}
		if n != len(data) {
			panic(data)
		}
		done <- struct{}{}
	}()

	buf := make([]byte, 64)
	n, err := r.Read(buf)
	if err != nil {
		panic(err)
	}

	fmt.Println("wrote:", n)         // 12
	fmt.Println("buf:", string(buf)) // Hello World!

	<-done

	r.Close()
	w.Close()
}
