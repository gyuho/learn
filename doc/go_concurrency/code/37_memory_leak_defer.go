package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func main() {

	fpath := "file.txt"

	f, err := os.OpenFile(fpath, os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		f, err = os.Create(fpath)
		if err != nil {
			panic(err)
		}
	}
	defer f.Close()

	i := 0
	for {
		i++

		// if this is inside a long-running function
		// this never gets run and causes memory leak
		defer func() {
			if _, err := f.WriteString(fmt.Sprintf("LINE %d\n", i)); err != nil {
				panic(err)
			}
		}()
		if i == 100 {
			break
		}
	}

	time.Sleep(time.Second)

	fc, err := toString(fpath)
	fmt.Println(fpath, "contents:", fc)
	// file.txt contents:

	defer func() {
		if err := os.Remove(fpath); err != nil {
			panic(err)
		}
	}()
}

func toString(fpath string) (string, error) {
	file, err := os.Open(fpath)
	if err != nil {
		// NOT retur nil, err
		// []byte can be null but not string
		return "", err
	}
	defer file.Close()

	// func ReadAll(r io.Reader) ([]byte, error)
	tbytes, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(tbytes), nil
}
