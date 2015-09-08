package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"strings"
)

var (
	sourceData = `
<?xml version="1.0" encoding="UTF-8"?>
<list id="3218" typeid="3218">
  <title>All Topics And Genres</title>
  <subcategory name="News">
    <item id="1149" num="1" type="topic">
      <title>Afghanistan</title>
      <additionalInfo>Afghanistan</additionalInfo>
    </item>
    <item id="10001" num="92" type="genre">
      <title>Rock</title>
      <additionalInfo>Rock, pop, and folk music performances and features from NPR news, NPR cultural programs, and NPR Music stations.</additionalInfo>
    </item>
    <item id="1103" num="93" type="topic">
      <title>Studio Sessions</title>
      <additionalInfo>Musicians perform and discuss their work in the studios of NPR and NPR Music station partners. Live music sessions, interviews, and the best new songs in rock, pop, folk, classical, jazz, blues, urban, and world music. Watch video sessions.</additionalInfo>
    </item>
    <item id="10004" num="94" type="genre">
      <title>World</title>
      <additionalInfo>World music and features from NPR news, NPR cultural programs, and NPR Music stations.</additionalInfo>
    </item>
  </subcategory>
</list>
`
)

type data struct {
	Title         string `xml:"title"`
	SubCategories []struct {
		Name  string `xml:"name,attr"`
		Items []struct {
			Title string `xml:"title"`
			Info  string `xml:"additionalInfo"`
		} `xml:"item"`
	} `xml:"subcategory"`
}

func main() {
	func() {
		rs := data{}
		rs.SubCategories = []struct {
			Name  string `xml:"name,attr"`
			Items []struct {
				Title string `xml:"title"`
				Info  string `xml:"additionalInfo"`
			} `xml:"item"`
		}{}
		rs.Title = "google.com"
		buf := new(bytes.Buffer)
		if err := xml.NewEncoder(buf).Encode(rs); err != nil {
			panic(err)
		}
		fmt.Println()
		fmt.Printf("xml.NewEncoder with struct: %+v\n", buf.String())
		// xml.NewEncoder with struct: <data><title>google.com</title></data>
	}()

	func() {
		rs := data{}

		// func NewDecoder(r io.Reader) *Decoder
		// func NewReader(s string) *Reader
		dec := xml.NewDecoder(strings.NewReader(sourceData))
		for {
			// func (d *Decoder) Decode(v interface{}) error
			if err := dec.Decode(&rs); err == io.EOF {
				break
			} else if err != nil {
				panic(err)
			}
		}
		fmt.Println()
		fmt.Printf("xml.NewDecoder with struct: %+v\n", rs)
		// xml.NewDecoder with struct: {Title:All Topics And Genres SubCategories:[{Na
		// ...
	}()
}
