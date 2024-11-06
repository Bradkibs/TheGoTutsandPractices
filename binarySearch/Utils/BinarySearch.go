package Utils

import (
	"sort"
)

func SortArray(array []int) {
	sort.Ints(array)
}

// BinarySearch performs a recursive binary search on a sorted slice.
func BinarySearch(array []int, value int) bool {
	return binarySearchHelper(array, value, 0, len(array)-1)
}

// binarySearchHelper is a helper function that performs the actual recursive search.
func binarySearchHelper(array []int, value, first, last int) bool {
	if first > last {
		return false
	}

	midpoint := (first + last) / 2

	if array[midpoint] == value {
		return true
	} else if value < array[midpoint] {
		return binarySearchHelper(array, value, first, midpoint-1)
	} else {
		return binarySearchHelper(array, value, midpoint+1, last)
	}
}
