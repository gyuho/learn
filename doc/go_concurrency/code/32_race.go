// race.go
package main

import "sync"

func updateSliceData(sliceData *[]int, num int, wg *sync.WaitGroup) {
	defer wg.Done()
	*sliceData = append(*sliceData, num)
}

func updateMapData(mapData *map[int]bool, num int, wg *sync.WaitGroup) {
	defer wg.Done()
	(*mapData)[num] = true
}

var wg sync.WaitGroup

func main() {
	var sliceData = []int{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go updateSliceData(&sliceData, i, &wg)
	}
	wg.Wait()

	var mapData = map[int]bool{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go updateMapData(&mapData, i, &wg)
	}
	wg.Wait()
}

/*
==================
WARNING: DATA RACE
Read by goroutine 5:
  main.updateSliceData()
      /home/ubuntu/race.go:7 +0x5f

Previous write by goroutine 4:
  main.updateSliceData()
      /home/ubuntu/race.go:7 +0x147

Goroutine 5 (running) created at:
  main.main()
      /home/ubuntu/race.go:21 +0xfd

Goroutine 4 (finished) created at:
  main.main()
      /home/ubuntu/race.go:21 +0xfd
==================
==================
WARNING: DATA RACE
Read by goroutine 5:
  runtime.growslice()
      /usr/local/go/src/runtime/slice.go:37 +0x0
  main.updateSliceData()
      /home/ubuntu/race.go:7 +0xcd

Previous write by goroutine 4:
  main.updateSliceData()
      /home/ubuntu/race.go:7 +0x104

Goroutine 5 (running) created at:
  main.main()
      /home/ubuntu/race.go:21 +0xfd

Goroutine 4 (finished) created at:
  main.main()
      /home/ubuntu/race.go:21 +0xfd
==================
==================
WARNING: DATA RACE
Write by goroutine 5:
  runtime.mapassign1()
      /usr/local/go/src/runtime/hashmap.go:383 +0x0
  main.updateMapData()
      /home/ubuntu/race.go:12 +0x94

Previous write by goroutine 4:
  runtime.mapassign1()
      /usr/local/go/src/runtime/hashmap.go:383 +0x0
  main.updateMapData()
      /home/ubuntu/race.go:12 +0x94

Goroutine 5 (running) created at:
  main.main()
      /home/ubuntu/race.go:28 +0x1cf

Goroutine 4 (finished) created at:
  main.main()
      /home/ubuntu/race.go:28 +0x1cf
==================
Found 3 data race(s)
exit status 66

*/
