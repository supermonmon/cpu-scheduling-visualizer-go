package processes

import (
	"cpu-scheduling-algorithms/algorithms"
	"cpu-scheduling-algorithms/models"
	"fmt"
)

const (
	FCFSAlgorithm = "FCFS"
	SJFAlgorithm  = "SJF"
)

func VisualizeOption(processesList []models.Process, currentAlgorithm string) {
	if currentAlgorithm == "None" {
		fmt.Println("\n\n\033[1;31mPlease choose an algorithm first!\033[0m")
		return
	}

	switch currentAlgorithm {
	case FCFSAlgorithm:
		avgTAT, avgWaiting := algorithms.FirstComeFirstServe(processesList)
		fmt.Printf("Average Turnaround Time (FCFS): %.2f\n", avgTAT)
		fmt.Printf("Average Waiting Time (FCFS): %.2f\n", avgWaiting)
	case SJFAlgorithm:
		avgTAT, avgWaiting := calculateSJFAverages(processesList)
		fmt.Printf("Average Turnaround Time (SJF): %.2f\n", avgTAT)
		fmt.Printf("Average Waiting Time (SJF): %.2f\n", avgWaiting)
	default:
		fmt.Println("Invalid algorithm chosen!")
	}
}

// calculateSJFAverages calculates the average Turnaround Time and average Waiting Time for SJF algorithm
func calculateSJFAverages(processesList []models.Process) (float64, float64) {
	// Implement SJF algorithm and calculate TAT and Waiting Time here
	fmt.Println("WORKING SJF ALGORITHM...")
	return 0, 0
}
