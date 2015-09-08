package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var x uint32 = 0x0A0B0C0D
	switch *(*byte)(unsafe.Pointer(&x)) {
	case 0x0A:
		fmt.Println("Big-Endian")
	case 0x0D:
		fmt.Println("Little-Endian")
	}
}
