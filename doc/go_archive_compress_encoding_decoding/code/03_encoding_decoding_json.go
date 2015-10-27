package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type pair struct {
	ProxyID  string `json:"proxyID"`
	Endpoint string `json:"endpoint"`
}

func makePair(proxyID, endpoint string) pair {
	return pair{ProxyID: proxyID, Endpoint: endpoint}
}

func encodePair(proxyID, endpoint string) (string, error) {
	p := pair{ProxyID: proxyID, Endpoint: endpoint}
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(p); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func decodePair(rd io.Reader) (pair, error) {
	p := pair{}
	dec := json.NewDecoder(rd)
	for {
		if err := dec.Decode(&p); err == io.EOF {
			break
		} else if err != nil {
			return pair{}, err
		}
	}
	return p, nil
}

func equalPair(s0, s1 string) bool {
	p0, err0 := decodePair(strings.NewReader(s0))
	if err0 != nil {
		return false
	}
	p1, err1 := decodePair(strings.NewReader(s1))
	if err1 != nil {
		return false
	}
	return p0.ProxyID == p1.ProxyID && p0.Endpoint == p1.Endpoint
}

func main() {
	p0 := makePair("test_id", "http://localhost:8080")
	fmt.Printf("makePair: %+v\n", p0)
	// makePair: {ProxyID:test_id Endpoint:http://localhost:8080}

	s0, err := encodePair("test_id", "http://localhost:8080")
	if err != nil {
		panic(err)
	}
	fmt.Printf("encodePair: %s\n", s0)
	// encodePair: {"proxyID":"test_id","endpoint":"http://localhost:8080"}

	p1, err := decodePair(strings.NewReader(`{
	"proxyID": "test_id",
	"endpoint": "http://localhost:8080"
}
`))
	if err != nil {
		panic(err)
	}
	fmt.Printf("decodePair: %+v\n", p1)
	// decodePair: {ProxyID:test_id Endpoint:http://localhost:8080}

	fmt.Println(equalPair(`{
	"proxyID": "test_id",
	"endpoint": "http://localhost:8080"
}
`, `{
	"endpoint": "http://localhost:8080",
	"proxyID": "test_id"
}
`)) // true
}
