package main

import "fmt"

/*
Go supports multiple assignment
Swap(&slice[i+1], &slice[lidx])

func Swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

Pseudocode from CLRS

quickSort(A, p, r)
	if p < r
		q = partition(A, p, r)
		quickSort(A, p, q-1)
		quickSort(A, q+1, r)

Partition(A, p, r)
	x = A[r]
	i = p - 1
	for j = p to r - 1
		if A[j] =< x
			i = i + 1
			exchange A[i] with A[j]
	exchange A[i+1] with A[r]
	return i+1
*/

// O(n * log n), in place with O(log(n)) stack space
// First choose a pivot element.
// And partition the array around the pivot.
// Around pivot, bigger ones moved to the right.
// Smaller ones moved to the left.
// Repeat (Recursion)
func quickSort(slice []int, fidx, lidx int) {
	if fidx < lidx {
		mid := partition(slice, fidx, lidx)
		quickSort(slice, fidx, mid-1)
		quickSort(slice, mid+1, lidx)
	}
}

// O(n)
// partition literally partition an integer array around a pivot.
// Elements bigger than pivot move to the right.
// Smaller ones move to left.
// It returns the index of the pivot element.
// After partition, the pivot element is places
// where it should be in the final sorted array.
// fidx = index of first element, usually 0
// lidx = index of last element, usually slice.size - 1
func partition(slice []int, fidx int, lidx int) int {
	x := slice[lidx]
	i := fidx - 1

	for j := fidx; j < lidx; j++ {
		if slice[j] <= x {
			i++
			slice[i], slice[j] = slice[j], slice[i] // to swap
		}
	}
	slice[i+1], slice[lidx] = slice[lidx], slice[i+1]
	return i + 1
}

func main() {
	slice := []int{9, -13, 4, -2, 3, 1, -10, 21, 12}
	quickSort(slice, 0, len(slice)-1)
	fmt.Println(slice)
	// [-13 -10 -2 1 3 4 9 12 21]
}
