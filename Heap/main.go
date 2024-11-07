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
	fmt.Println("----------Deleting for the 1st time the min heap-------------------")
	newHeap := Utils.DeleteFromHeap(data[:len(data)-1], "min")
	Utils.PrintTree(newHeap)
	fmt.Println("----------Deleting for the 2nd  time the min heap-------------------")
	newHeap = Utils.DeleteFromHeap(data[:len(data)-2], "min")
	Utils.PrintTree(newHeap)

	fmt.Println("----------Deleting for the 3rd time the min heap-------------------")
	newHeap = Utils.DeleteFromHeap(data[:len(data)-3], "min")
	Utils.PrintTree(newHeap)
	for i := 4; i < len(data); i++ {
		fmt.Printf("----------Deleting for the %v time the min heap-------------------\n", i)
		newHeap = Utils.DeleteFromHeap(data[:len(data)-i], "min")
		Utils.PrintTree(newHeap)
	}

}
