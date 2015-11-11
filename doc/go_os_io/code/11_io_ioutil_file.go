package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"syscall"
)

func main() {
	func() {
		fpath := "temp.txt"
		txt := "Hello World!"
		if err := toFile1(txt, fpath); err != nil {
			panic(err)
		}
		defer os.Remove(fpath)
		if s, err := fromFile1(fpath); err != nil {
			panic(err)
		} else {
			fmt.Println(fpath, ":", s)
		}
	}()
	// temp.txt : Hello World!

	func() {
		fpath := "temp.txt"
		txt := "Hello World!"
		if err := toFile2(txt, fpath); err != nil {
			panic(err)
		}
		defer os.Remove(fpath)
		if s, err := fromFile1(fpath); err != nil {
			panic(err)
		} else {
			fmt.Println(fpath, ":", s)
		}
	}()
	// temp.txt : Hello World!

	func() {
		fpath := "temp.txt"
		txt := "Hello World!"
		if err := toFile3(txt, fpath); err != nil {
			panic(err)
		}
		defer os.Remove(fpath)
		if s, err := fromFile1(fpath); err != nil {
			panic(err)
		} else {
			fmt.Println(fpath, ":", s)
		}
	}()
	// temp.txt : Hello World!

	func() {
		fpath := "temp.txt"
		txt := "Hello World!"
		if err := toFile1(txt, fpath); err != nil {
			panic(err)
		}
		defer os.Remove(fpath)
		if s, err := fromFile2(fpath); err != nil {
			panic(err)
		} else {
			fmt.Println(fpath, ":", s)
		}
	}()
	// temp.txt : Hello World!

	func() {
		fpath := "temp.txt"
		txt := "Hello World!"
		if err := toFile1(txt, fpath); err != nil {
			panic(err)
		}
		defer os.Remove(fpath)
		isSupported := isDirectIOSupported()
		fmt.Println("isDirectIOSupported:", isSupported)
		if isSupported {
			if s, err := fromFile3(fpath); err != nil {
				panic(err)
			} else {
				fmt.Println(fpath, ":", s)
			}
		}
	}()
	// temp.txt : Hello World!
}

func toFile1(txt, fpath string) error {
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

func toFile2(txt, fpath string) error {
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

func toFile3(txt, fpath string) error {
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

func fromFile1(fpath string) (string, error) {
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

func fromFile2(fpath string) (string, error) {
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

// panic: read temp.txt: invalid argument
func fromFile3(fpath string) (string, error) {
	f, err := os.OpenFile(fpath, os.O_RDONLY|syscall.O_DIRECT, 0777)
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

func isDirectIOSupported() bool {
	path := os.TempDir()
	defer os.RemoveAll(path)
	f, err := os.OpenFile(path, syscall.O_DIRECT, 0)
	defer f.Close()
	return err == nil
}
