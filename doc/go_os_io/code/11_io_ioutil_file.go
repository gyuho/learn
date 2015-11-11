package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	func() {
		fpath := "temp.txt"
		txt := "Hello World!"
		if err := toFile1(txt, fpath); err != nil {
			panic(err)
		}
		if s, err := fromFile(fpath); err != nil {
			panic(err)
		} else {
			fmt.Println(fpath, ":", s)
		}
		os.Remove(fpath)
	}()
	// temp.txt : Hello World!

	func() {
		fpath := "temp.txt"
		txt := "Hello World!"
		if err := toFile2(txt, fpath); err != nil {
			panic(err)
		}
		if s, err := fromFile(fpath); err != nil {
			panic(err)
		} else {
			fmt.Println(fpath, ":", s)
		}
		os.Remove(fpath)
	}()
	// temp.txt : Hello World!

	func() {
		fpath := "temp.txt"
		txt := "Hello World!"
		if err := toFile3(txt, fpath); err != nil {
			panic(err)
		}
		if s, err := fromFile(fpath); err != nil {
			panic(err)
		} else {
			fmt.Println(fpath, ":", s)
		}
		os.Remove(fpath)
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

func fromFile(fpath string) (string, error) {
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
