package Utils

import (
	"fmt"
	"math"
)

// Adds a single value to the heap
func addToHeap(array []int, value int, Type string) []int {
	array = append(array, value)      // Add the new value at the end
	heapifyUp := Heapify(array, Type) // Get the correct heapify function
	heapifyUp()                       // Restore the heap property
	return array
}

// Pops the root value of the heap and reduces the size correctly
func DeleteFromHeap(array []int, Type string) []int {
	if len(array) == 0 {
		return array
	}
	// Replace the root with the last element and reduce the slice length
	array[0] = array[len(array)-1]
	array = array[:len(array)-1] // Properly reduce the slice length

	// Restore the heap property
	heapifyDown := Heapify(array, Type)
	heapifyDown()

	return array
}

// Heapify returns a closure that heapifies the array according to the given type
func Heapify(array []int, Type string) func() {
	if Type == "max" {
		return func() {
			for i := len(array)/2 - 1; i >= 0; i-- {
				maxHeapifyDown(array, i)
			}
		}
	} else {
		return func() {
			for i := len(array)/2 - 1; i >= 0; i-- {
				minHeapifyDown(array, i)
			}
		}
	}
}

// Helper function for max-heap property restoration
func maxHeapifyDown(array []int, index int) {
	length := len(array)
	for index < length {
		leftChild := 2*index + 1
		rightChild := 2*index + 2
		swapIndex := index

		if leftChild < length && array[leftChild] > array[swapIndex] {
			swapIndex = leftChild
		}
		if rightChild < length && array[rightChild] > array[swapIndex] {
			swapIndex = rightChild
		}
		if swapIndex != index {
			array[index], array[swapIndex] = array[swapIndex], array[index]
			index = swapIndex
		} else {
			break
		}
	}
}

// Helper function for min-heap property restoration
func minHeapifyDown(array []int, index int) {
	length := len(array)
	for index < length {
		leftChild := 2*index + 1
		rightChild := 2*index + 2
		swapIndex := index

		if leftChild < length && array[leftChild] < array[swapIndex] {
			swapIndex = leftChild
		}
		if rightChild < length && array[rightChild] < array[swapIndex] {
			swapIndex = rightChild
		}
		if swapIndex != index {
			array[index], array[swapIndex] = array[swapIndex], array[index]
			index = swapIndex
		} else {
			break
		}
	}
}

// PrintTree displays the heap array as a structured tree with only active elements
func PrintTree(array []int) {
	n := len(array)
	if n == 0 {
		fmt.Println("Heap is empty")
		return
	}

	// Calculate the number of levels in the tree
	level := int(math.Floor(math.Log2(float64(n)))) + 1
	currentIndex := 0

	for i := 0; i < level && currentIndex < n; i++ {
		// Calculate the number of nodes at the current level
		numNodes := int(math.Min(float64(n-currentIndex), math.Pow(2, float64(i))))
		padding := int(math.Pow(2, float64(level-i-1)) - 1)

		// Print spaces before the first node on each level
		fmt.Print(spaces(padding))

		// Print nodes at the current level
		for j := 0; j < numNodes && currentIndex < n; j++ {
			fmt.Printf("%d", array[currentIndex])
			currentIndex++
			// Print spaces between nodes
			if j < numNodes-1 {
				fmt.Print(spaces(2*padding + 1))
			}
		}
		fmt.Println() // Move to the next level
	}
}

// spaces returns a string with the specified number of spaces
func spaces(count int) string {
	return fmt.Sprintf("%*s", count, "")
}
