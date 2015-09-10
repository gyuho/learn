/*
go run -race 33_no_race_with_mutex.go
*/
package main

import (
	"fmt"
	"sync"
)

func updateSliceDataWithLock(sliceData *[]int, num int, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	mutex.Lock()
	*sliceData = append(*sliceData, num)
	mutex.Unlock()
}

func updateMapDataWithLock(mapData *map[int]bool, num int, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	mutex.Lock()
	(*mapData)[num] = true
	mutex.Unlock()
}

// Mutexes can be created as part of other structures
type sliceData struct {
	sync.Mutex
	s []int
}

func updateSliceDataWithLockStruct(data *sliceData, num int, wg *sync.WaitGroup) {
	defer wg.Done()
	data.Lock()
	data.s = append(data.s, num)
	data.Unlock()
}

// Mutexes can be created as part of other structures
type mapData struct {
	sync.Mutex
	m map[int]bool
}

func updateMapDataWithLockStruct(data *mapData, num int, wg *sync.WaitGroup) {
	defer wg.Done()
	data.Lock()
	data.m[num] = true
	data.Unlock()
}

func main() {
	var (
		wg    sync.WaitGroup
		mutex sync.Mutex
	)

	var ds1 = []int{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go updateSliceDataWithLock(&ds1, i, &wg, &mutex)
	}
	wg.Wait()
	fmt.Println(ds1)

	var dm1 = map[int]bool{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go updateMapDataWithLock(&dm1, i, &wg, &mutex)
	}
	wg.Wait()
	fmt.Println(dm1)

	ds2 := sliceData{}
	ds2.s = []int{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go updateSliceDataWithLockStruct(&ds2, i, &wg)
	}
	wg.Wait()
	fmt.Println(ds2)

	dm2 := mapData{}
	dm2.m = make(map[int]bool)
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go updateMapDataWithLockStruct(&dm2, i, &wg)
	}
	wg.Wait()
	fmt.Println(dm2)
}
