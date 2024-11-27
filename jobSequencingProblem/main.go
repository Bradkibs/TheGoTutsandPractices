package main

import (
	"fmt"
	"github.com/Bradkibs/TheGoTutsandPractices/jobSequencingProblem/helpers"
)

func main() {
	sequenceMatrix := make(map[string][]int, 10)

	sequenceMatrix["Job1"] = []int{100, 2} // Job1 has profit 100 and deadline 2
	sequenceMatrix["Job2"] = []int{50, 1}  // Job2 has profit 50 and deadline 1
	sequenceMatrix["Job3"] = []int{70, 3}  // Job3 has profit 70 and deadline 3
	sequenceMatrix["Job4"] = []int{25, 5}  // Job4 has profit 25 and deadline 5
	sequenceMatrix["Job5"] = []int{60, 1}  // Job5 has profit 60 and deadline 1
	sequenceMatrix["Job6"] = []int{40, 4}  // Job6 has profit 40 and deadline 4
	sequenceMatrix["Job7"] = []int{30, 2}  // Job7 has profit 30 and deadline 2
	sequenceMatrix["Job8"] = []int{90, 3}  // Job8 has profit 90 and deadline 3
	sequenceMatrix["Job9"] = []int{85, 1}  // Job9 has profit 85 and deadline 1
	sequenceMatrix["Job10"] = []int{10, 2} // Job10 has profit 10 and deadline 2

	result := helpers.Sequencer(sequenceMatrix, "maximize")
	profit := 0
	for _, item := range result {
		profit += sequenceMatrix[item][0]
	}

	fmt.Println("The result of maximizing the profit on the jobs is by taking jobs: ", result, " in their respective order and the profit is: ", profit)
}
