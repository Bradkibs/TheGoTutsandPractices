package main

import "fmt"

func fibonacciAppend(slice []int, elems ...int) []int {
	/* An implementation of append() method for slices but uses fibonacci resizing
	strategy instead of using basic doubling strategy.
	*/
	fib1, fib2 := 1, 2
	for _, elem := range elems {
		if len(slice) == cap(slice) {
			newCap := fib1 + fib2
			fib1, fib2 = fib2, newCap
			newSlice := make([]int, len(slice), newCap)
			copy(newSlice, slice)
			slice = newSlice
		}
		slice = append(slice, elem)
	}
	return slice
}

func main() {
	slice := make([]int, 0, 1)
	for i := 0; i < 20; i++ {
		slice = fibonacciAppend(slice, i)
		fmt.Printf("After appending %d: len=%d cap=%d\n", i, len(slice), cap(slice))
	}
}
