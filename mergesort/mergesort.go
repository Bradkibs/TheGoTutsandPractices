package main

import "fmt"

// Inplace merge sort
func mergeSort(list *[]int, low int, high int) {
	if low < high {
		mid := (low + high) / 2

		mergeSort(list, low, mid)
		mergeSort(list, mid+1, high)
		merge(list, low, mid, high)
	}
}

func merge(list *[]int, low int, mid int, high int) {
	slice := *list
	left := make([]int, mid-low+1)
	right := make([]int, high-mid)

	copy(left, slice[low:mid+1])
	copy(right, slice[mid+1:high+1])
	fmt.Println("Left list: ", left)
	fmt.Println("Right list: ", right)

	i, j, k := 0, 0, low
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
		k++
	}
	for i < len(left) {
		slice[k] = left[i]
		i++
		k++
	}
	for j < len(right) {
		slice[k] = right[j]
		j++
		k++
	}

}
