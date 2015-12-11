package main

import "fmt"

type Flags struct {
	Key string
}

type Data struct {
	Flags Flags
}

type DataPointer struct {
	Flags *Flags
}

// X
func functionData1(d Data) {
	d.Flags.Key = "updated by functionData1"
}

// O
func functionData2(d *Data) {
	d.Flags.Key = "updated by functionData2"
}

// X
func (d Data) methodData1() {
	d.Flags.Key = "updated by methodData1"
}

// O
func (d *Data) methodData2() {
	d.Flags.Key = "updated by methodData2"
}

// O
func functionDataPointer1(d DataPointer) {
	d.Flags.Key = "updated by functionDataPointer1"
}

// O
func functionDataPointer2(d *DataPointer) {
	d.Flags.Key = "updated by functionDataPointer2"
}

// O
func (d DataPointer) methodDataPointer1() {
	d.Flags.Key = "updated by methodDataPointer1"
}

// O
func (d *DataPointer) methodDataPointer2() {
	d.Flags.Key = "updated by methodDataPointer2"
}

func main() {
	f := Flags{Key: ""}

	data := Data{}
	data.Flags = f

	dataPointer := DataPointer{}
	dataPointer.Flags = &f

	// X
	functionData1(data)
	fmt.Println("after functionData1:", data)

	// O
	functionData2(&data)
	fmt.Println("after functionData2:", data)

	// X
	data.methodData1()
	fmt.Println("after methodData1:", data)

	// O
	(&data).methodData2()
	fmt.Println("after methodData2:", data)

	// O
	functionDataPointer1(dataPointer)
	fmt.Println("after functionDataPointer1:", dataPointer.Flags)

	// O
	functionDataPointer2(&dataPointer)
	fmt.Println("after functionDataPointer2:", dataPointer.Flags)

	// O
	dataPointer.methodDataPointer1()
	fmt.Println("after methodDataPointer1:", dataPointer.Flags)

	// O
	(&dataPointer).methodDataPointer2()
	fmt.Println("after methodDataPointer2:", dataPointer.Flags)
}

/*
after functionData1: {{}}
after functionData2: {{updated by functionData2}}
after methodData1: {{updated by functionData2}}
after methodData2: {{updated by methodData2}}
after functionDataPointer1: &{updated by functionDataPointer1}
after functionDataPointer2: &{updated by functionDataPointer2}
after methodDataPointer1: &{updated by methodDataPointer1}
after methodDataPointer2: &{updated by methodDataPointer2}
*/
