package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-yaml/yaml"
)

var sourceData = `
_enabled: 1
_tag: AAAA
product_type:
- part_id: 6496293
- part_id: 6493671
configurations:
  my_name:
  - product_type:
    - part_id: 6496293
    - part_id: 6493671
    - part_id: 6495763
    - part_id: 6495685
    - part_id: 6495582
    - part_id: 6495660
    experiment_id: SurvJun1Cat1A0
    price_opt:
      number: 0
    price: 3
  - product_type:
    - part_id: 6496302
    experiment_id: SurvJun4Cat3A4
    price_opt:
      number: 4
      price_id: 1
    price: 30
model: model32
location: USA
platform:
  desktop: {}
  mobile: {}
`

type subData struct {
	PartID string `part_id`
}

type data struct {
	ProductType []subData `product_type`
}

func main() {
	func() {
		rs := make(map[string]interface{})
		rs["Go"] = "Google"
		d, err := yaml.Marshal(rs)
		if err != nil {
			panic(err)
		}
		fmt.Println()
		fmt.Printf("yaml.Marshal with map: %+v\n", string(d))
		// yaml.Marshal with map: Go: Google
	}()

	func() {
		rs := make(map[interface{}]interface{})
		if err := yaml.Unmarshal([]byte(sourceData), &rs); err != nil {
			panic(err)
		}
		fmt.Println()
		fmt.Printf("yaml.Unmarshal with map: %+v\n", rs)
		// yaml.Unmarshal with map: map[platform:map[desktop:map[] mobile:map[]
		// ...

		for k := range rs {
			switch t := k.(type) {
			case string:
				fmt.Println(t)
			}
		}
	}()

	func() {
		rs := struct {
			Name     string
			Number   int `num`
			Category struct {
				Location string
			} `Category`
		}{}
		rs.Name = "Google"
		rs.Category = struct {
			Location string
		}{
			Location: "USA",
		}
		d, err := yaml.Marshal(rs)
		if err != nil {
			panic(err)
		}
		fmt.Println()
		fmt.Printf("yaml.Marshal with struct:\n%+v\n", string(d))
		// yaml.Marshal with struct:
		/*
		   name: Google
		   num: 0
		   Category:
		     location: USA
		*/
	}()

	func() {
		rs := data{}
		if err := yaml.Unmarshal([]byte(sourceData), &rs); err != nil {
			panic(err)
		}
		fmt.Println()
		fmt.Printf("yaml.Unmarshal with struct: %+v\n", rs)
		// yaml.Unmarshal with struct: {ProductType:[{PartID:6496293} {PartID:6493671}]}
	}()

	fmt.Println()
	fmt.Println("isJSON:", isJSON([]byte(sourceData)))
	// false
}

func isJSON(data []byte) bool {
	cmap := make(map[interface{}]interface{})
	dec := json.NewDecoder(bytes.NewReader(data))
	for {
		if err := dec.Decode(&cmap); err == io.EOF {
			break
		} else if err != nil {
			return false
		}
	}
	return true
}
