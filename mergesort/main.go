package main

import "fmt"

func main() {
	newList := [14]int{21, 1, 3, 2, 45, 12, 67, 5, 4, 3, 87, 76, 12, 1}
	fmt.Println("-------------List before merge sort--------------------")
	fmt.Println(newList)
	sliceOfNewList := newList[:]
	fmt.Println("--------------The mergesort algorithm visualization-------------")
	mergeSort(&sliceOfNewList, 0, 13)

	fmt.Println("---------The list After merge sort------------")
	fmt.Println(sliceOfNewList)
}
