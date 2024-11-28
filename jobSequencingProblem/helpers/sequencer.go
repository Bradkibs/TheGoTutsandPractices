package helpers

import "sort"

type job struct {
	name     string
	profit   int
	deadline int
}

func Sequencer(sequenceMatrix map[string][]int, sequenceType string) []string {
	// sorting the sequenceMatrix
	sortedJobs := sorter(sequenceMatrix, sequenceType)

	// if we assume that each task takes 1 unit of time we need to find the job with the greatest deadline for efficient deadline scheduling
	maxDeadline := 0
	for item := range sortedJobs {
		if sortedJobs[item].deadline > maxDeadline {
			maxDeadline = sortedJobs[item].deadline
		}
	}
	Results := make([]string, maxDeadline)

	// Assigning jobs to the available slots (starting from the last available slot)
	for _, job := range sortedJobs {
		// Finding the latest available slot for this job (starting from job.deadline - 1)
		for j := job.deadline - 1; j >= 0; j-- {
			// If the slot is free, assign the job to this slot
			if Results[j] == "" {
				Results[j] = job.name
				break
			}
		}
	}
	return Results

}

func sorter(sequenceMatrix map[string][]int, sequenceType string) []job {
	var jobs []job

	for name, details := range sequenceMatrix {
		jobs = append(jobs, job{
			name:     name,
			profit:   details[0],
			deadline: details[1],
		})
	}
	switch sequenceType {

	case "maximize":

		sort.Slice(jobs, func(i int, j int) bool {
			return jobs[i].profit > jobs[j].profit
		})
	case "minimize":
		sort.Slice(jobs, func(i int, j int) bool {
			return jobs[i].profit < jobs[j].profit
		})
	default:
		sort.Slice(jobs, func(i int, j int) bool {
			return jobs[i].profit > jobs[j].profit
		})
	}

	return jobs
}
