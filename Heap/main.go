package main

import (
	"Heap/Utils"
	"fmt"
)

func main() {
	data := []int{23, 13, 13, 12, 1, 27, 89, 78, 90, 34}
	fmt.Println("-----------------Before Being Heapified-------------------------")
	Utils.PrintTree(data)
	fmt.Println("-----------------After  Being Max Heapified------------------------")
	maxHeapify := Utils.Heapify(data, "max")
	maxHeapify()
	Utils.PrintTree(data)
	fmt.Println("-----------------After  Being Min Heapified------------------------")
	minHeapify := Utils.Heapify(data, "min")
	minHeapify()
	Utils.PrintTree(data)
	fmt.Println("----------Deleting from the min heap-------------------")
	Utils.DeleteFromHeap(data, "min")
	Utils.PrintTree(data)
	fmt.Println("----------Deleting again from the min heap-------------------")
	Utils.DeleteFromHeap(data, "min")
	Utils.PrintTree(data)
}
