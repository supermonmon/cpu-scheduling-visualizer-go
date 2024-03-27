package algorithms

import (
	"fmt"
)

type processData struct {
	pid         string
	arrivalTime int
	burstTime   int
	completed   bool
}

func allCompleted(data []processData) bool {
	for _, process := range data {
		if !process.completed {
			return false
		}
	}
	return true
}

// SRTF implements the Shortest Remaining Time First (SRTF) scheduling algorithm with preemptive execution
func SRTF(processID []string, arrivalTime, burstTime []int) {
	fmt.Println("SRTF Scheduling Results (Preemptive):")

	// Create a data structure to store process information with burst time and completion flag
	processDataSlice := make([]processData, len(processID))
	for i := range processID {
		processDataSlice[i] = processData{processID[i], arrivalTime[i], burstTime[i], false}
	}

	// Initialize slices for waiting time, completion time, and turnaround time
	waitingTime := make([]int, len(processID))
	completionTime := make([]int, len(processID))
	turnAroundTime := make([]int, len(processID))

	var currentTime int = 0    // Track current time
	var completedProcesses int // Track number of completed processes

	for completedProcesses < len(processID) {
		// Find the available process with the shortest remaining burst time
		shortestJob := -1
		for j := 0; j < len(processID); j++ {
			if !processDataSlice[j].completed && processDataSlice[j].arrivalTime <= currentTime {
				if shortestJob == -1 || processDataSlice[j].burstTime < processDataSlice[shortestJob].burstTime {
					shortestJob = j
				}
			}
		}

		// No available processes at current time
		if shortestJob == -1 {
			currentTime++
			continue
		}

		// Execute the process for 1 unit of time
		processDataSlice[shortestJob].burstTime--
		currentTime++

		// Check if the process is completed
		if processDataSlice[shortestJob].burstTime == 0 {
			completedProcesses++
			processDataSlice[shortestJob].completed = true
			completionTime[shortestJob] = currentTime
			turnAroundTime[shortestJob] = completionTime[shortestJob] - processDataSlice[shortestJob].arrivalTime
			waitingTime[shortestJob] = turnAroundTime[shortestJob] - burstTime[shortestJob]
		}
	}

	// Print results similar to SRTF
	fmt.Println("+-------+------------+-----------+------------+--------------+--------------+")
	fmt.Println("| PID   | AT         | BT        | CT         | WT           | TAT          |")
	fmt.Println("+-------+------------+-----------+------------+--------------+--------------+")
	totalWT := 0
	totalTAT := 0
	for i := range processID {
		fmt.Printf("| %4s  | %10d | %9d | %10d | %12d | %12d |\n", processID[i], arrivalTime[i], burstTime[i], completionTime[i], waitingTime[i], turnAroundTime[i])
		totalWT += waitingTime[i]
		totalTAT += turnAroundTime[i]
	}
	fmt.Println("+-------+------------+-----------+------------+--------------+--------------+")

	// Calculate average waiting time and turnaround time
	avgWaitingTime := float64(totalWT) / float64(len(processID))
	avgTurnAroundTime := float64(totalTAT) / float64(len(processID))

	// Print average waiting time and turnaround time
	fmt.Printf("Average Waiting Time: %.2f\n", avgWaitingTime)
	fmt.Printf("Average Turnaround Time: %.2f\n", avgTurnAroundTime)
}
