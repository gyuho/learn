package main

import "fmt"

// O(n)
func merge(s1, s2 []int) []int {
	final := make([]int, len(s1)+len(s2))
	i, j := 0, 0
	for i < len(s1) && j < len(s2) {
		if s1[i] <= s2[j] {
			final[i+j] = s1[i]
			i++
			continue
		}
		final[i+j] = s2[j]
		j++
	}
	for i < len(s1) {
		final[i+j] = s1[i]
		i++
	}
	for j < len(s2) {
		final[i+j] = s2[j]
		j++
	}
	return final
}

// O(n * log n)
// Recursively splits the array into subarrays, until only one element.
// From each subarray, merge them into a sorted array.
func mergeSort(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}
	idx := len(slice) / 2
	left := mergeSort(slice[:idx])
	right := mergeSort(slice[idx:])
	return merge(left, right)
}

// O(n * log n)
// Recursively splits the array into subarrays, until only one element.
// From each subarray, merge them into a sorted array.
func mergeSortConcurrency(slice []int, ch chan []int) {
	if len(slice) < 2 {
		ch <- slice
		return
	}

	idx := len(slice) / 2
	ch1, ch2 := make(chan []int), make(chan []int)

	go mergeSortConcurrency(slice[:idx], ch1)
	go mergeSortConcurrency(slice[idx:], ch2)

	left := <-ch1
	right := <-ch2

	// close after waiting to receive all
	close(ch1)
	close(ch2)

	ch <- merge(left, right)
}

func main() {
	sliceA := []int{9, -13, 4, -2, 3, 1, -10, 21, 12}
	fmt.Println(mergeSort(sliceA))

	sliceB := []int{9, -13, 4, -2, 3, 1, -10, 21, 12}
	ch := make(chan []int)
	go mergeSortConcurrency(sliceB, ch)
	fmt.Println(<-ch)
}
