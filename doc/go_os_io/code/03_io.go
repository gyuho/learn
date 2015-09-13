package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// And os.File implements Read and Write method
// therefore satisfies io.Reader and io.Writer method

func main() {
	fpath := "testdata/sample.json"

	file, err := os.Open(fpath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	tbytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	jsonStream := string(tbytes)
	decodeString(jsonStream)
	// map[Go:Gopher Hello:World]

	decodeFile(file)
	// map[]

	// need to open again
	file2, err := os.Open(fpath)
	if err != nil {
		panic(err)
	}
	decodeFile(file2)
	// map[Go:Gopher Hello:World]
}

func decodeFile(file *os.File) {
	rmap := map[string]string{}
	dec := json.NewDecoder(file)
	for {
		if err := dec.Decode(&rmap); err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
	fmt.Printf("%+v\n", rmap)
}

func decodeString(jsonStream string) {
	rmap := map[string]string{}
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		if err := dec.Decode(&rmap); err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
	fmt.Printf("%+v\n", rmap)
}
