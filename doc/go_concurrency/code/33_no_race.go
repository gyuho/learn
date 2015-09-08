package example

import (
	"sync"
	"testing"
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

var (
	wg    sync.WaitGroup
	mutex sync.Mutex
)

func TestRace(t *testing.T) {
	var sliceData = []int{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go updateSliceDataWithLock(&sliceData, i, &wg, &mutex)
	}
	wg.Wait()

	var mapData = map[int]bool{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go updateMapDataWithLock(&mapData, i, &wg, &mutex)
	}
	wg.Wait()
}

// go test -race
// no race condition
