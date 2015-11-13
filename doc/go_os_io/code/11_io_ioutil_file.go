package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"syscall"
	"unsafe"
)

func main() {
	func() {
		fpath := "temp.txt"
		txt := "Hello World!"
		if err := toFileWriteString(txt, fpath); err != nil {
			panic(err)
		}
		defer os.Remove(fpath)
		if s, err := fromFileOpenReadAll(fpath); err != nil {
			panic(err)
		} else {
			fmt.Println("fromFileOpenReadAll:", s)
		}
	}()
	// fromFileOpenReadAll: Hello World!

	func() {
		fpath := "temp.txt"
		txt := "Hello World!"
		if err := toFileIO(txt, fpath); err != nil {
			panic(err)
		}
		defer os.Remove(fpath)
		if s, err := fromFileOpenReadAll(fpath); err != nil {
			panic(err)
		} else {
			fmt.Println("fromFileOpenReadAll:", s)
		}
	}()
	// fromFileOpenReadAll: Hello World!

	func() {
		fpath := "temp.txt"
		txt := "Hello World!"
		if err := toFileWrite(txt, fpath); err != nil {
			panic(err)
		}
		defer os.Remove(fpath)
		if s, err := fromFileOpenReadAll(fpath); err != nil {
			panic(err)
		} else {
			fmt.Println("fromFileOpenReadAll:", s)
		}
	}()
	// fromFileOpenReadAll: Hello World!

	func() {
		fpath := "toFileWriteSyscall.txt"
		txt := "Hello World!"
		if err := toFileWriteSyscall(txt, fpath); err != nil {
			panic(err)
		}
		defer os.Remove(fpath)
		if s, err := fromFileOpenReadAll(fpath); err != nil {
			panic(err)
		} else {
			fmt.Println("toFileWriteSyscall and fromFileOpenReadAll:", s)
		}
	}()
	// toFileWriteSyscall and fromFileOpenReadAll: Hello World!

	func() {
		fpath := "temp.txt"
		txt := "Hello World!"
		if err := toFileWriteString(txt, fpath); err != nil {
			panic(err)
		}
		defer os.Remove(fpath)
		if s, err := fromFileOpenFileReadAll(fpath); err != nil {
			panic(err)
		} else {
			fmt.Println("fromFileOpenFileReadAll:", s)
		}
	}()
	// fromFileOpenFileReadAll: Hello World!

	func() {
		fpath := "temp.txt"
		txt := "Hello World!"
		if err := toFileWriteString(txt, fpath); err != nil {
			panic(err)
		}
		defer os.Remove(fpath)
		if s, err := fromFileOpenFileReadFull(fpath, len(txt)); err != nil {
			panic(err)
		} else {
			fmt.Println("fromFileOpenFileReadFull:", s)
		}
	}()
	// fromFileOpenFileReadFull: Hello World!

	func() {
		fpath := "temp.txt"
		txt := strings.Repeat("Hello World!", 10000)
		if err := toFileWriteString(txt, fpath); err != nil {
			panic(err)
		}
		defer os.Remove(fpath)
		isSupported := isDirectIOSupported(fpath)
		fmt.Println("isDirectIOSupported:", isSupported)
		if isSupported {
			if s, err := fromFileDirectIO(fpath); err != nil {
				panic(err)
			} else {
				fmt.Println("fromFileDirectIO:", s[:10], "...")
			}
		}
	}()
	// isDirectIOSupported: true
	// fromFileDirectIO: Hello Worl ...
}

func toFileWriteString(txt, fpath string) error {
	f, err := os.OpenFile(fpath, os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		// OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
		f, err = os.Create(fpath)
		if err != nil {
			return err
		}
	}
	defer f.Close()
	if _, err := f.WriteString(txt); err != nil {
		return err
	}
	return nil
}

func toFileIO(txt, fpath string) error {
	f, err := os.OpenFile(fpath, os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		// OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
		f, err = os.Create(fpath)
		if err != nil {
			return err
		}
	}
	defer f.Close()
	if _, err := io.WriteString(f, txt); err != nil {
		return err
	}
	return nil
}

func toFileWrite(txt, fpath string) error {
	f, err := os.OpenFile(fpath, os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		// OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
		f, err = os.Create(fpath)
		if err != nil {
			return err
		}
	}
	defer f.Close()
	if _, err := f.Write([]byte(txt)); err != nil {
		return err
	}
	return nil
}

func toFileWriteSyscall(txt, fpath string) error {
	f, err := os.OpenFile(fpath, os.O_RDWR|syscall.MAP_POPULATE, 0777)
	if err != nil {
		// OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
		f, err = os.Create(fpath)
		if err != nil {
			return err
		}
	}
	defer f.Close()
	if _, err := f.Write([]byte(txt)); err != nil {
		return err
	}
	return nil
}

func fromFileOpenReadAll(fpath string) (string, error) {
	f, err := os.Open(fpath)
	if err != nil {
		// NOT retur nil, err
		// []byte can be null but not string
		return "", err
	}
	defer f.Close()
	// func ReadAll(r io.Reader) ([]byte, error)
	tbytes, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	return string(tbytes), nil
}

func fromFileOpenFileReadAll(fpath string) (string, error) {
	f, err := os.OpenFile(fpath, os.O_RDONLY, 0777)
	if err != nil {
		// NOT retur nil, err
		// []byte can be null but not string
		return "", err
	}
	defer f.Close()
	// func ReadAll(r io.Reader) ([]byte, error)
	tbytes, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	return string(tbytes), nil
}

func fromFileOpenFileReadFull(fpath string, length int) (string, error) {
	f, err := os.OpenFile(fpath, os.O_RDONLY, 0777)
	if err != nil {
		// NOT retur nil, err
		// []byte can be null but not string
		return "", err
	}
	defer f.Close()
	buf := make([]byte, length)
	if _, err := io.ReadFull(f, buf); err != nil {
		return "", err
	}
	return string(buf), nil
}

func isDirectIOSupported(fpath string) bool {
	f, err := os.OpenFile(fpath, syscall.O_DIRECT, 0)
	defer f.Close()
	return err == nil
}

func fromFileDirectIO(fpath string) (string, error) {
	f, err := os.OpenFile(fpath, os.O_RDONLY|syscall.O_DIRECT, 0777)
	if err != nil {
		return "", err
	}
	defer f.Close()
	block := AlignedBlock(BlockSize)
	if _, err := io.ReadFull(f, block); err != nil {
		return "", err
	}
	return string(block), nil
}

/*****************************************************/

// Copied from https://github.com/ncw/directio

// alignment returns alignment of the block in memory
// with reference to AlignSize
//
// Can't check alignment of a zero sized block as &block[0] is invalid
func alignment(block []byte, AlignSize int) int {
	return int(uintptr(unsafe.Pointer(&block[0])) & uintptr(AlignSize-1))
}

// AlignedBlock returns []byte of size BlockSize aligned to a multiple
// of AlignSize in memory (must be power of two)
func AlignedBlock(BlockSize int) []byte {
	block := make([]byte, BlockSize+AlignSize)
	if AlignSize == 0 {
		return block
	}
	a := alignment(block, AlignSize)
	offset := 0
	if a != 0 {
		offset = AlignSize - a
	}
	block = block[offset : offset+BlockSize]
	// Can't check alignment of a zero sized block
	if BlockSize != 0 {
		a = alignment(block, AlignSize)
		if a != 0 {
			log.Fatal("Failed to align block")
		}
	}
	return block
}

const (
	// Size to align the buffer to
	AlignSize = 4096

	// Minimum block size
	BlockSize = 4096
)

// OpenFile is a modified version of os.OpenFile which sets O_DIRECT
func OpenFile(name string, flag int, perm os.FileMode) (file *os.File, err error) {
	return os.OpenFile(name, syscall.O_DIRECT|flag, perm)
}

/*****************************************************/
