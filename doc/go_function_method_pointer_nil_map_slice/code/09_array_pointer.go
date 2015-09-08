package main

import "fmt"

var array = [3]string{"A", "B", "C"}

func changeArray1(m [3]string) { m[0] = "X" }

func changeArray1pt(m *[3]string) { (*m)[0] = "Y" }

type arrayType [3]string

func (m arrayType) changeArray2() { m[1] = "XX" }

func (m *arrayType) changeArray2p()  { m[1] = "YY" }
func (m *arrayType) changeArray2pt() { (*m)[1] = "ZZ" }

func main() {
	// (X) no change
	changeArray1(array)
	fmt.Println("changeArray1:", array) // [A B C]

	// (O) change
	changeArray1pt(&array)
	fmt.Println("changeArray1pt:", array) // [Y B C]

	// (X) no change
	arrayType(array).changeArray2()
	fmt.Println(".changeArray2():", array) // [Y B C]

	// (O) change
	(*arrayType)(&array).changeArray2p()
	fmt.Println(".changeArray2pt():", array) // [Y YY C]

	// (O) change
	(*arrayType)(&array).changeArray2pt()
	fmt.Println(".changeArray2pt():", array) // [Y ZZ C]
}
