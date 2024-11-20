package main

import (
	quickSort "github.com/Bradkibs/TheGoTutsandPractices/quickSort/sort"
	"log"
)

func main() {
	unsortedList := []int{1, 4, 23, 6, 5, 22, 98, 0, 1, 45, 23, 43, 32}
	//Adding a large number at the end of the list so that it is naturally sorted.
	unsortedListForQuickSort := append(unsortedList, 100000000)
	log.Println("-----------unsortedListForQuickSort--------------")
	log.Println(unsortedListForQuickSort)
	quickSort.Quicksort(unsortedListForQuickSort, 0, len(unsortedListForQuickSort)-1)
	log.Println("----------------SortedList-----------------")
	// removing the large number added at the end of the list
	unsortedListForQuickSort = unsortedListForQuickSort[:len(unsortedListForQuickSort)-1]
	log.Println(unsortedListForQuickSort)
}
