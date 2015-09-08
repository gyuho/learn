package main

import "fmt"

func main() {
	// direct access doesn't need any pointer
	array := [5]int{0, 1, 2, 3, 4}
	for i := 0; i < len(array); i++ {
		array[i]++ // DOES CHANGE
	}
	fmt.Println(array) // [1 2 3 4 5]

	// array needs pointer to update its element
	swapArray1(2, 3, array) // NO CHANGE
	fmt.Println(array)      // [1 2 3 4 5]

	swapArray2(2, 3, &array) // DOES CHANGE
	fmt.Println(array)       // [1 2 3 4 5]

	// direct access doesn't need any pointer
	slice := []int{0, 1, 2, 3, 4}
	for i := 0; i < len(slice); i++ {
		slice[i]++ // DOES CHANGE
	}
	fmt.Println(slice) // [1 2 3 4 5]

	// slice elements are pointers
	swapSlice1(2, 3, slice) // DOES CHANGE
	fmt.Println(slice)      // [1 2 4 3 5]

	swapSlice2(2, 3, &slice) // DOES CHANGE
	fmt.Println(slice)       // [1 2 3 4 5]
}

func swapArray1(i, j int, array [5]int) {
	array[i], array[j] = array[j], array[i]
}

func swapArray2(i, j int, array *[5]int) {
	(*array)[i], (*array)[j] = (*array)[j], (*array)[i]
}

func swapSlice1(i, j int, slice []int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func swapSlice2(i, j int, slice *[]int) {
	(*slice)[i], (*slice)[j] = (*slice)[j], (*slice)[i]
}
