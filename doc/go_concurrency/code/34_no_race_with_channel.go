// race.go
package main

func sendWithChannel(ch chan int, num int) {
	ch <- num
}

func main() {
	ch1 := make(chan int)
	for i := 0; i < 100; i++ {
		go sendWithChannel(ch1, i)
	}
	cn := 0
	var sliceData = []int{}
	for v := range ch1 {
		sliceData = append(sliceData, v)
		cn++
		if cn == 100 {
			close(ch1)
		}
	}

	ch2 := make(chan int)
	var mapData = map[int]bool{}
	for i := 0; i < 100; i++ {
		go sendWithChannel(ch2, i)
	}
	cn = 0
	for v := range ch2 {
		mapData[v] = true
		cn++
		if cn == 100 {
			close(ch2)
		}
	}
}

// go run -race race.go
// no race condition
