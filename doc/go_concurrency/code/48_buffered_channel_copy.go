package main

import "fmt"

func main() {
	ch1 := make(chan string, 5000)
	fmt.Println("ch1:", len(ch1), cap(ch1)) // ch1: 0 5000
	ch1 <- "aaa"
	fmt.Println("ch1:", len(ch1), cap(ch1)) // ch1: 1 5000
	fmt.Println(<-ch1)                      // aaa
	fmt.Println("ch1:", len(ch1), cap(ch1)) // ch1: 0 5000

	ch2 := getCh2()
	fmt.Println("ch2:", len(ch2), cap(ch2)) // ch2: 0 5000
	ch2 <- "aaa"
	fmt.Println("ch2:", len(ch2), cap(ch2)) // ch2: 1 5000

	ds := createDatas()
	for _, d := range ds {
		fmt.Println("ds ch:", len(d.ch), cap(d.ch))
	}

	var di DataInterface
	di = &ds[0]
	chd := di.Chan()
	fmt.Println("chd:", len(chd), cap(chd)) // chd: 1 5000
}

func getCh2() chan string {
	return make(chan string, 5000)
}

type Data struct {
	ch chan string
}

func createDatas() []Data {
	ds := make([]Data, 5)
	bufCh := make(chan string, 5000)
	for i := range ds {
		ds[i].ch = bufCh
	}
	return ds
}

type DataInterface interface {
	Chan() chan string
}

func (d *Data) Chan() chan string {
	return d.ch
}
