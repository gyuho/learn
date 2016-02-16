package main

import "fmt"

type MyInterface interface {
	Hey() string
}

type MyType struct {
	Name string
}

func (t MyType) Hey() string {
	return t.Name
}

func main() {
	// interface evaluated runtime
	ex := MyType{Name: "Hello"}
	m := make(map[string]MyInterface)

	m["id"] = ex
	fmt.Println(m)
	// map[id:{Hello}]
}
