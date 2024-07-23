package main

import "fmt"

var a, b  = 1, 2

// fibonacci returns a function that returns successive Fibonacci numbers.
func fibonacci() func() int {
	return func() int {
		next := a
		a, b = b, next + b
		return a
	}
}


// fibonacciAppend appends elements to a slice using a Fibonacci resizing strategy.
func fibonacciAppend(slice []int, elems ...int) []int {
	// Initialize the Fibonacci generator
	nextFib := fibonacci()



	// Loop over the elements to be appended
	for _, elem := range elems {
		// Check if resizing is needed
		if len(slice) == cap(slice) {
			// Get the next Fibonacci capacity
			newCap := nextFib()
			fmt.Println(newCap)

			// Create a new slice with the new capacity
			newSlice := make([]int, len(slice), newCap)
			copy(newSlice, slice)
			slice = newSlice
			//
			//// Update the current capacity
			//currentCap = newCap
		}
		// Append the element to the slice
		slice = append(slice, elem)
	}
	return slice
}

func main() {
	slice := make([]int, 0, 1)
	for i := 0; i < 50; i++ {
		slice = fibonacciAppend(slice, i)
		fmt.Printf("After appending %d: len=%d cap=%d\n", i, len(slice), cap(slice))
	}
}
