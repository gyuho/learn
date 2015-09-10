package main

import (
	"fmt"
	"strings"
)

func nonDeterministicMapUpdateV1() {
	for i := 0; i < 10; i++ {
		fmt.Println("nonDeterministicMapUpdateV1 TRY =", i)
		mmap := map[string]int{
			"hello": 10,
			"world": 50,
			"here":  5,
			"go":    7,
			"code":  11,
		}
		length := len(mmap)
		for k, v := range mmap {
			mmap[strings.ToUpper(k)] = v * v
			delete(mmap, k)
		}
		if length == len(mmap) {
			fmt.Println("Luckily, Deterministic with nonDeterministicMapUpdateV1:", length, len(mmap))
			return
		}
		fmt.Println("Non-Deterministic with nonDeterministicMapUpdateV1:", length, len(mmap))
	}
}

func nonDeterministicMapUpdateV2() {
	for i := 0; i < 10; i++ {
		fmt.Println("nonDeterministicMapUpdateV2 TRY =", i)
		mmap := map[string]int{
			"hello": 10,
			"world": 50,
			"here":  5,
			"go":    7,
			"code":  11,
		}
		ks := []string{}
		length := len(mmap)
		for k, v := range mmap {
			mmap[strings.ToUpper(k)] = v * v
			ks = append(ks, k)
		}
		for _, k := range ks {
			delete(mmap, k)
		}
		if length == len(mmap) {
			fmt.Println("Luckily, Deterministic with nonDeterministicMapUpdateV2:", length, len(mmap))
			return
		}
		fmt.Println("Non-Deterministic with nonDeterministicMapUpdateV2:", length, len(mmap))
	}
}

func nonDeterministicMapUpdateV3() {
	for i := 0; i < 10; i++ {
		fmt.Println("nonDeterministicMapUpdateV3 TRY =", i)
		mmap := map[string]int{
			"hello": 10,
			"world": 50,
			"here":  5,
			"go":    7,
			"code":  11,
		}
		length := len(mmap)
		for k := range mmap {
			v := mmap[k]
			mmap[strings.ToUpper(k)] = v * v
			delete(mmap, k)
		}
		if length == len(mmap) {
			fmt.Println("Luckily, Deterministic with nonDeterministicMapUpdateV3:", length, len(mmap))
			return
		}
		fmt.Println("Non-Deterministic with nonDeterministicMapUpdateV3:", length, len(mmap))
	}
}

func deterministicMapSet() {
	for i := 0; i < 10000; i++ {
		mmap := make(map[int]bool)
		for i := 0; i < 10000; i++ {
			mmap[i] = true
		}
		length := len(mmap)
		for k := range mmap {
			delete(mmap, k)
		}
		if len(mmap) == 0 {
			fmt.Println("Deterministic with deterministicMapSet:", length, len(mmap))
			return
		}
		fmt.Println("Non-Deterministic with deterministicMapSet:", length, len(mmap))
	}
}

func deterministicMapDelete() {
	for i := 0; i < 10000; i++ {
		fmt.Println("deterministicMapDelete TRY =", i)
		mmap := map[string]int{
			"hello": 10,
			"world": 50,
			"here":  5,
			"go":    7,
			"code":  11,
		}
		length := len(mmap)
		for k := range mmap {
			delete(mmap, k)
		}
		if len(mmap) == 0 {
			fmt.Println("Deterministic with deterministicMapDelete:", length, len(mmap))
			return
		}
		fmt.Println("Non-Deterministic with deterministicMapDelete:", length, len(mmap))
	}
}

func deterministicMapUpdate() {
	for i := 0; i < 10000; i++ {
		fmt.Println("deterministicMapUpdate TRY =", i)
		mmap := map[string]int{
			"hello": 10,
			"world": 50,
			"here":  5,
			"go":    7,
			"code":  11,
		}
		mmapCopy := make(map[string]int)
		length := len(mmap)
		for k, v := range mmap {
			mmapCopy[strings.ToUpper(k)] = v * v
		}
		for k := range mmap {
			delete(mmap, k)
		}
		if length == len(mmapCopy) || len(mmap) != 0 {
			fmt.Println("Deterministic with deterministicMapUpdate:", length, len(mmapCopy))
			return
		} else {
			mmapCopy = make(map[string]int) // to initialize(empty)
			//
			// (X)
			// mmapCopy = nil
		}
		fmt.Println("Non-Deterministic with deterministicMapUpdate:", length, len(mmap))
	}
}

func main() {
	nonDeterministicMapUpdateV1()
	fmt.Println()
	nonDeterministicMapUpdateV2()
	fmt.Println()
	nonDeterministicMapUpdateV3()

	fmt.Println()

	deterministicMapSet()
	fmt.Println()
	deterministicMapDelete()
	fmt.Println()
	deterministicMapUpdate()
}

/*
These are all non-deterministic.
If you are lucky, the map gets updated inside range.

nonDeterministicMapUpdateV1 TRY = 0
Non-Deterministic with nonDeterministicMapUpdateV1: 5 4
nonDeterministicMapUpdateV1 TRY = 1
Non-Deterministic with nonDeterministicMapUpdateV1: 5 4
nonDeterministicMapUpdateV1 TRY = 2
Luckily, Deterministic with nonDeterministicMapUpdateV1: 5 5

nonDeterministicMapUpdateV2 TRY = 0
Non-Deterministic with nonDeterministicMapUpdateV2: 5 2
nonDeterministicMapUpdateV2 TRY = 1
Non-Deterministic with nonDeterministicMapUpdateV2: 5 2
nonDeterministicMapUpdateV2 TRY = 2
Non-Deterministic with nonDeterministicMapUpdateV2: 5 2
nonDeterministicMapUpdateV2 TRY = 3
Non-Deterministic with nonDeterministicMapUpdateV2: 5 2
nonDeterministicMapUpdateV2 TRY = 4
Non-Deterministic with nonDeterministicMapUpdateV2: 5 2
nonDeterministicMapUpdateV2 TRY = 5
Non-Deterministic with nonDeterministicMapUpdateV2: 5 2
nonDeterministicMapUpdateV2 TRY = 6
Non-Deterministic with nonDeterministicMapUpdateV2: 5 3
nonDeterministicMapUpdateV2 TRY = 7
Non-Deterministic with nonDeterministicMapUpdateV2: 5 4
nonDeterministicMapUpdateV2 TRY = 8
Non-Deterministic with nonDeterministicMapUpdateV2: 5 2
nonDeterministicMapUpdateV2 TRY = 9
Non-Deterministic with nonDeterministicMapUpdateV2: 5 4

nonDeterministicMapUpdateV3 TRY = 0
Non-Deterministic with nonDeterministicMapUpdateV3: 5 4
nonDeterministicMapUpdateV3 TRY = 1
Non-Deterministic with nonDeterministicMapUpdateV3: 5 4
nonDeterministicMapUpdateV3 TRY = 2
Non-Deterministic with nonDeterministicMapUpdateV3: 5 4
nonDeterministicMapUpdateV3 TRY = 3
Luckily, Deterministic with nonDeterministicMapUpdateV3: 5 5

Deterministic with deterministicMapSet: 10000 0

deterministicMapDelete TRY = 0
Deterministic with deterministicMapDelete: 5 0

deterministicMapUpdate TRY = 0
Deterministic with deterministicMapUpdate: 5 5
*/
