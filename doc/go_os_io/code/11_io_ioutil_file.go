package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"syscall"

	"github.com/ncw/directio"
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

func fromFileDirectIO(fpath string) (string, error) {
	f, err := os.OpenFile(fpath, os.O_RDONLY|syscall.O_DIRECT, 0777)
	if err != nil {
		return "", err
	}
	defer f.Close()
	block := directio.AlignedBlock(directio.BlockSize)
	if _, err := io.ReadFull(f, block); err != nil {
		return "", err
	}
	return string(block), nil
}

func isDirectIOSupported(fpath string) bool {
	f, err := os.OpenFile(fpath, syscall.O_DIRECT, 0)
	defer f.Close()
	return err == nil
}
