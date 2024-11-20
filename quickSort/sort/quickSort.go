package sort

import (
	"log"
)

func Quicksort(list []int, low int, high int) {
	if low < high {
		pivotPoint := partition(list, low, high)
		log.Println("The pivotPoint is at position:", pivotPoint)
		Quicksort(list, low, pivotPoint-1)
		Quicksort(list, pivotPoint+1, high)
	}
}
