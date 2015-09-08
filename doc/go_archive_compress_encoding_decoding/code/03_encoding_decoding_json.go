package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

var (
	sourceURL  = "http://httpbin.org/headers"
	sourceData = `
{
  "headers": {
    "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8", 
    "Accept-Encoding": "gzip, deflate, sdch", 
    "Accept-Language": "en-US,en;q=0.8,ko;q=0.6", 
    "Cookie": "_ga=GA1.2.630704613.1440642077", 
    "Dnt": "1", 
    "Host": "httpbin.org", 
    "Referer": "http://httpbin.org/", 
    "Upgrade-Insecure-Requests": "1", 
    "User-Agent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/44.0.2403.157 Safari/537.36"
  }
}
`
)

type data struct {
	Headers struct {
		Accept    string `json:"Accept"`
		Host      string `json:"Host"`
		UserAgent string `json:"User-Agent"`
	} `json:"headers"`
}

func main() {
	func() {
		rs := make(map[string]interface{})
		rs["Go"] = "Google"
		buf := new(bytes.Buffer)
		if err := json.NewEncoder(buf).Encode(rs); err != nil {
			panic(err)
		}
		fmt.Println()
		fmt.Printf("json.NewEncoder with map: %+v\n", buf.String())
		// json.NewEncoder with map: {"Go":"Google"}
	}()

	func() {
		// func Get(url string) (resp *Response, err error)
		// type Response struct {
		// 		Body io.ReadCloser
		resp, err := http.Get(sourceURL)
		if resp == nil {
			return
		}
		defer resp.Body.Close()
		if err != nil {
			panic(err)
		}
		rs := make(map[string]interface{})

		// resp implements Write
		// resp.Body implements Read

		// http://jmoiron.net/blog/crossing-streams-a-love-letter-to-ioreader/
		// func ReadAll(r io.Reader) ([]byte, error)
		// body, err := ioutil.ReadAll(resp.Body)

		// func NewDecoder(r io.Reader) *Decoder
		dec := json.NewDecoder(resp.Body)
		for {
			// func (d *Decoder) Decode(v interface{}) error
			if err := dec.Decode(&rs); err == io.EOF {
				break
			} else if err != nil {
				panic(err)
			}
		}
		fmt.Println()
		fmt.Printf("json.NewDecoder with map: %+v\n", rs)
		// json.NewDecoder with map: map[headers:map[Accept-Encoding:gzip Host:httpbin.org User-Agent:Go-http-client/1.1]]
	}()

	func() {
		rs := data{}
		rs.Headers = struct {
			Accept    string `json:"Accept"`
			Host      string `json:"Host"`
			UserAgent string `json:"User-Agent"`
		}{
			"",
			"",
			"",
		}
		rs.Headers.Host = "google.com"
		buf := new(bytes.Buffer)
		if err := json.NewEncoder(buf).Encode(rs); err != nil {
			panic(err)
		}
		fmt.Println()
		fmt.Printf("json.NewEncoder with struct: %+v\n", buf.String())
		// json.NewEncoder with struct: {"headers":{"Accept":"","Host":"google.com","User-Agent":""}}
	}()

	func() {
		rs := data{}

		// func NewDecoder(r io.Reader) *Decoder
		// func NewReader(s string) *Reader
		dec := json.NewDecoder(strings.NewReader(sourceData))
		for {
			// func (d *Decoder) Decode(v interface{}) error
			if err := dec.Decode(&rs); err == io.EOF {
				break
			} else if err != nil {
				panic(err)
			}
		}
		fmt.Println()
		fmt.Printf("json.NewDecoder with struct: %+v\n", rs)
		// json.NewDecoder with struct: {Headers:{Accept:text/html,application/
		// ...
	}()
}
