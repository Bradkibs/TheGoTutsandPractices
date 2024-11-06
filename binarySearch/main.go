package main

import (
	Utils "binarySearch/Utils"
	"fmt"
)

func main() {
	List := []int{23, 1, 2, 43, 6, 7, 8, 5, 12, 78}
	Utils.SortArray(List)
	valueToBeSearched := 43
	isPresent := Utils.BinarySearch(List, valueToBeSearched)
	if isPresent == true {
		fmt.Printf("The value %v is in the list", valueToBeSearched)
		return
	}
	fmt.Printf("The value %v is NOT in the list", valueToBeSearched)
}
