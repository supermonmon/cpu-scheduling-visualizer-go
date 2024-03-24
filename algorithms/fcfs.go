package algorithms

import (
	"cpu-scheduling-algorithms/models"
)

func FirstComeFirstServe(processesList []models.Process) (float64, float64) {
	totalTAT := 0
	totalWaiting := 0
	currentTime := 0

	for _, process := range processesList {
		if process.ArrivalTime > currentTime {
			currentTime = process.ArrivalTime
		}
		process.CompletionTime = currentTime + process.BurstTime
		totalTAT += process.CompletionTime - process.ArrivalTime
		totalWaiting += currentTime - process.ArrivalTime
		currentTime = process.CompletionTime
	}

	avgTAT := float64(totalTAT) / float64(len(processesList))
	avgWaiting := float64(totalWaiting) / float64(len(processesList))
	return avgTAT, avgWaiting
}
