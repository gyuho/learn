package main

import "fmt"

func typeName1(v interface{}) string {
	switch typedValue := v.(type) {
	case int:
		fmt.Println("Value:", typedValue)
		return "int"
	case string:
		fmt.Println("Value:", typedValue)
		return "string"
	default:
		fmt.Println("Value:", typedValue)
		return "unknown"
	}
	panic("unreachable")
}

func typeName2(v interface{}) string {
	switch v.(type) {
	case int:
		return "int"
	case string:
		return "string"
	default:
		return "unknown"
	}
	panic("unreachable")
}

type Stringer interface {
	String() string
}

type fakeString struct {
	content string
}

// function used to implement the Stringer interface
func (s *fakeString) String() string {
	return s.content
}

func printString(value interface{}) {
	switch str := value.(type) {
	case string:
		fmt.Println(str)
	case Stringer:
		fmt.Println(str.String())
	}
}

func main() {
	fmt.Println(typeName1(1))
	fmt.Println(typeName1("Hello"))
	fmt.Println(typeName1(-.1))
	/*
	   Value: 1
	   int
	   Value: Hello
	   string
	   Value: -0.1
	   unknown
	*/

	fmt.Println(typeName2(1))       // int
	fmt.Println(typeName2("Hello")) // string
	fmt.Println(typeName2(-.1))     // unknown

	s := &fakeString{"Ceci n'est pas un string"}
	printString(s)                // Ceci n'est pas un string
	printString("Hello, Gophers") // Hello, Gophers
}
