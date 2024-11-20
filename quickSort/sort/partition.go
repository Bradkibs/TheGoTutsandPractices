package sort

import "log"

func partition(list []int, low int, high int) int {

	pivot := list[low]
	i, j := low+1, high
	for i <= j {
		for list[i] <= pivot && i <= high {
			i++
		}
		for list[j] > pivot && j >= low {
			j--
		}
		if i < j {
			swap(list, i, j)
		}

	}
	swap(list, low, j)
	return j
}

func swap(list []int, value1 int, value2 int) {

	list[value1], list[value2] = list[value2], list[value1]
	log.Println("The new list is:", list)
}
