package main

import "fmt"

var slice = []string{"A", "B", "C"}

func changeSlice1(m []string) { m[0] = "X" }

func changeSlice1pt(m *[]string) { (*m)[0] = "Y" }

type sliceType []string

// var slice = sliceType{"A", "B", "C"}

func (m sliceType) changeSlice2() { m[1] = "XX" }

// func (m *sliceType) changeSlice2p() { m[1] = "YY" }

func (m *sliceType) changeSlice2pt() { (*m)[1] = "YY" }

func changeSlice3(m []string) { m = append(m, "XXX") }

func changeSlice3pt(m *[]string) { *m = append(*m, "YYY") }

func (m sliceType) changeSlice4() { m = append(m, "XXXX") }

func (m *sliceType) changeSlice4pt() { *m = append(*m, "YYYY") }

func main() {
	// (O) change
	changeSlice1(slice)
	fmt.Println("changeSlice1:", slice) // [X B C]

	// (O) change
	changeSlice1pt(&slice)
	fmt.Println("changeSlice1pt:", slice) // [Y B C]

	// (O) change
	sliceType(slice).changeSlice2()
	fmt.Println(".changeSlice2():", slice) // [Y XX C]

	// (O) change
	(*sliceType)(&slice).changeSlice2pt()
	fmt.Println(".changeSlice2pt():", slice) // [Y YY C]

	// (X) no change
	changeSlice3(slice)
	fmt.Println("changeSlice3:", slice) // [Y YY C]

	// (O) change
	changeSlice3pt(&slice)
	fmt.Println("changeSlice3pt:", slice) // [Y YY C YYY]

	// (X) no change
	sliceType(slice).changeSlice4()
	fmt.Println(".changeSlice4():", slice) // [Y YY C YYY]

	// (O) change
	(*sliceType)(&slice).changeSlice4pt()
	fmt.Println(".changeSlice4pt():", slice) // [Y YY C YYY YYYY]
}
