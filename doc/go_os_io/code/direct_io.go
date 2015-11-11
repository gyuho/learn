package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	fmt.Println(isDirectIOSupported())
}

func isDirectIOSupported() bool {
	path := os.TempDir()
	defer os.RemoveAll(path)
	f, err := os.OpenFile(path, syscall.O_DIRECT, 0)
	defer f.Close()
	return err == nil
}
