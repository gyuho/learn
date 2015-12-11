package main

import (
	"fmt"
	"strings"
)

func deterministic_map_update_0() {
	for i := 0; i < 500; i++ {
		m := make(map[int]bool)
		for i := 0; i < 500; i++ {
			m[i] = true
		}
		if len(m) != 500 {
			fmt.Println("deterministic_map_update_0 got non-determinstic:", len(m), "at", i)
			return
		}
	}
}

func deterministic_map_update_1() {
	for i := 0; i < 500; i++ {
		m := make(map[int]bool)
		for i := 0; i < 500; i++ {
			m[i] = true
		}
		for k := range m {
			m[k] = false
		}
		for _, v := range m {
			if v {
				fmt.Println("deterministic_map_update_1 got non-determinstic:", len(m), "at", i)
				return
			}
		}
		for k := range m {
			m[(k+1)*-1] = true
		}
		if len(m) != 2*500 {
			fmt.Println("deterministic_map_update_1 got non-determinstic:", len(m), "at", i)
			return
		}
	}
}

func deterministic_map_update_2() {
	for i := 0; i < 500; i++ {
		m := make(map[int]bool)
		for i := 0; i < 500; i++ {
			m[i] = true
		}
		for k, v := range m {
			_ = v
			m[k] = false
		}
		for _, v := range m {
			if v {
				fmt.Println("deterministic_map_update_2 got non-determinstic:", len(m), "at", i)
				return
			}
		}
		for k, v := range m {
			_ = v
			m[(k+1)*-1] = true
		}
		if len(m) != 2*500 {
			fmt.Println("deterministic_map_update_2 got non-determinstic:", len(m), "at", i)
			return
		}
	}
}

func deterministic_map_delete_0() {
	for i := 0; i < 500; i++ {
		m := make(map[int]bool)
		for i := 0; i < 500; i++ {
			m[i] = true
		}
		for k := range m {
			delete(m, k)
		}
		if len(m) != 0 {
			fmt.Println("deterministic_map_delete_0 got non-determinstic:", len(m), "at", i)
			return
		}
	}
}

func deterministic_map_delete_1() {
	for i := 0; i < 500; i++ {
		m := make(map[int]bool)
		for i := 0; i < 500; i++ {
			m[i] = true
		}
		for k, v := range m {
			_ = v
			delete(m, k)
		}
		if len(m) != 0 {
			fmt.Println("deterministic_map_delete_1 got non-determinstic:", len(m), "at", i)
			return
		}
	}
}

func non_deterministic_map_0() {
	for i := 0; i < 500; i++ {
		m := map[string]int{
			"a": 1,
			"b": 2,
			"c": 3,
			"d": 4,
			"e": 5,
		}
		len1 := len(m)
		for k := range m {
			m[strings.ToUpper(k)] = 100
			delete(m, k)
		}
		len2 := len(m)
		if len1 != len2 {
			fmt.Println("non_deterministic_map_0 is non-determinstic:", len1, len2, "at", i, "/", m)
			return
		}
	}
}

func non_deterministic_map_1() {
	for i := 0; i < 500; i++ {
		m := map[string]int{
			"a": 1,
			"b": 2,
			"c": 3,
			"d": 4,
			"e": 5,
		}
		len1 := len(m)
		for k, v := range m {
			m[strings.ToUpper(k)] = v * v
			delete(m, k)
		}
		len2 := len(m)
		if len1 != len2 {
			fmt.Println("non_deterministic_map_1 is non-determinstic:", len1, len2, "at", i, "/", m)
			return
		}
	}
}

func main() {
	deterministic_map_update_0()
	deterministic_map_update_1()
	deterministic_map_update_2()

	deterministic_map_delete_0()
	deterministic_map_delete_1()

	// non-deterministic when updating, deleting at the same time
	non_deterministic_map_0()
	non_deterministic_map_1()
	/*
	   non_deterministic_map_0 is non-determinstic: 5 4 at 0 / map[B:100 C:100 D:100 E:100]
	   non_deterministic_map_1 is non-determinstic: 5 4 at 0 / map[A:1 B:4 D:16 E:25]
	*/
}
