package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	// to create in the current directory.
	// f, err := ioutil.TempFile(".", "temp_prefix_")

	// creates in  TempFile uses the default directory
	// for temporary files (see os.TempDir)
	f, err := ioutil.TempFile("", "temp_prefix_")
	if err != nil {
		panic(err)
	}
	defer os.Remove(f.Name())
	defer f.Close()
	if err := f.Sync(); err != nil {
		panic(err)
	}
	if _, err := f.Seek(0, 0); err != nil {
		panic(err)
	}

	fmt.Println(os.TempDir())            // /tmp
	fmt.Println(f.Name())                // /tmp/temp_prefix_289175735
	fmt.Println(filepath.Base(f.Name())) // temp_prefix_289175735
	fmt.Println(filepath.Dir(f.Name()))  // /tmp
}
