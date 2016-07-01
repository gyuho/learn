package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

const (
	// privateFileMode grants owner to read/write a file.
	privateFileMode = 0600

	// privateDirMode grants owner to make/remove files inside the directory.
	privateDirMode = 0700
)

func openToRead(fpath string) (*os.File, error) {
	f, err := os.OpenFile(fpath, os.O_RDONLY, privateFileMode)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func openToOverwrite(fpath string) (*os.File, error) {
	f, err := os.OpenFile(fpath, os.O_RDWR|os.O_TRUNC|os.O_CREATE, privateFileMode)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func openToAppend(fpath string) (*os.File, error) {
	f, err := os.OpenFile(fpath, os.O_RDWR|os.O_APPEND|os.O_CREATE, privateFileMode)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func main() {
	fpath := "./testdata/sample.txt"

	func() {
		f, err := openToRead(fpath)
		if err != nil {
			panic(err)
		}
		defer func() {
			fmt.Println("Closing", f.Name())
			f.Close()
		}()
		if f.Name() != fpath {
			panic(f.Name())
		}
		tbytes, err := ioutil.ReadAll(f)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(tbytes))
	}()
	/*
	   Hello World!
	   Closing ./testdata/sample.txt
	*/

	fmt.Println()
	fmt.Println()

	func() {
		fpath := "test.txt"
		for range []int{0, 1} {
			f, err := openToOverwrite(fpath)
			if err != nil {
				panic(err)
			}
			if _, err := f.WriteString("Hello World!"); err != nil {
				panic(err)
			}
			f.Close()
		}
		f, err := openToRead(fpath)
		defer f.Close()
		if err != nil {
			panic(err)
		}
		tbytes, err := ioutil.ReadAll(f)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(tbytes))
		if err := os.Remove(fpath); err != nil {
			panic(err)
		}
	}()
	// Hello World!

	fmt.Println()
	fmt.Println()

	func() {
		fpath := "test.txt"
		for _, k := range []int{0, 1} {
			f, err := openToAppend(fpath)
			if err != nil {
				panic(err)
			}
			if _, err := f.WriteString(fmt.Sprintf("Hello World! %d\n", k)); err != nil {
				panic(err)
			}
			f.Close()
		}
		f, err := openToRead(fpath)
		defer f.Close()
		if err != nil {
			panic(err)
		}
		tbytes, err := ioutil.ReadAll(f)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(tbytes))
		if err := os.Remove(fpath); err != nil {
			panic(err)
		}
	}()
	/*
	   Hello World! 0
	   Hello World! 1
	*/
}
