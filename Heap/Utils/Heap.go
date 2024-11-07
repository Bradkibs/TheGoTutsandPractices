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

// Pops the root value of the heap
func DeleteFromHeap(array []int, Type string) []int {
	if len(array) == 0 {
		return array
	}
	array[0] = array[len(array)-1]      // Move the last element to the root
	array = array[:len(array)-1]        // Remove the last element
	heapifyDown := Heapify(array, Type) // Get the correct heapify function
	heapifyDown()                       // Restore the heap property
	return array
}

// Heapify returns a closure that heapifies the array according to the given type
func Heapify(array []int, Type string) func() {
	if Type == "max" {
		return func() {
			// Max heapify: Adjust nodes to ensure parent nodes are greater than children
			for i := len(array)/2 - 1; i >= 0; i-- {
				maxHeapifyDown(array, i)
			}
		}
	} else {
		return func() {
			// Min heapify: Adjust nodes to ensure parent nodes are less than children
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

		// Swap with larger child if heap property is violated
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

		// Swap with smaller child if heap property is violated
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

// printTree displays the heap array as a structured tree
func PrintTree(array []int) {
	n := len(array)
	if n == 0 {
		fmt.Println("Heap is empty")
		return
	}

	// Calculate the number of levels in the tree
	level := int(math.Floor(math.Log2(float64(n)))) + 1
	maxWidth := int(math.Pow(2, float64(level)))

	for i, j := 0, 0; i < level; i++ {
		// Calculate the number of nodes at the current level
		numNodes := int(math.Pow(2, float64(i)))

		// Calculate spacing for the nodes
		spacing := maxWidth / numNodes
		padding := spacing / 2

		// Print spaces before the first node on each level
		fmt.Print(spaces(padding, ""))

		// Print all nodes at the current level
		for k := 0; k < numNodes && j < n; k++ {
			fmt.Printf("%d", array[j])
			fmt.Print(spaces(spacing-1, ""))
			j++
		}
		fmt.Println() // Move to the next level
	}
}

// spaces returns a string with the specified number of spaces
func spaces(count int, Type string) string {
	return fmt.Sprintf("%*s", count, Type)
}
