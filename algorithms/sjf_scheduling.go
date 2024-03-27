package algorithms

import (
	"fmt"
	"os"
	"sort"
)

// SJF implements the Shortest Job First (SJF) scheduling algorithm with non-preemptive execution
func SJF(processID []string, arrivalTime, burstTime []int) {
	fmt.Println("SJF Scheduling Results (Non-preemptive):")

	// Create a data structure to store process information with burst time
	type processData struct {
		pid         string
		arrivalTime int
		burstTime   int
	}
	processDataSlice := make([]processData, len(processID))
	for i := range processID {
		processDataSlice[i] = processData{processID[i], arrivalTime[i], burstTime[i]}
	}

	// Sort processes by arrival time (ascending)
	sort.Slice(processDataSlice, func(i, j int) bool {
		return processDataSlice[i].arrivalTime < processDataSlice[j].arrivalTime
	})

	// Extract sorted data back to original slices
	for i := range processDataSlice {
		processID[i] = processDataSlice[i].pid
		arrivalTime[i] = processDataSlice[i].arrivalTime
		burstTime[i] = processDataSlice[i].burstTime
	}

	var waitingTime, completionTime, turnAroundTime []int
	var currentProcess int = -1              // Index of the currently executing process
	var currentBurst, currentTime int = 0, 0 // Track current burst time and time

	// SJF scheduling loop with non-preemption
	for i := range processID {
		// Find the first available process (not yet completed) with the lowest burst time
		shortestJob := -1
		for j := 0; j < len(processID); j++ {
			if arrivalTime[j] <= currentTime && burstTime[j] > 0 && (shortestJob == -1 || burstTime[j] < burstTime[shortestJob]) {
				shortestJob = j
			}
		}

		// Handle no available processes
		if shortestJob == -1 {
			currentTime = arrivalTime[i] // Move time to next process arrival
			continue
		}

		// Update current process, burst time, and completion time
		if currentProcess == -1 {
			currentProcess = shortestJob
		}
		currentBurst = burstTime[shortestJob]
		currentTime += currentBurst
		burstTime[shortestJob] = 0 // Mark completed process

		// Calculate waiting and turnaround times
		waitingTime = append(waitingTime, currentTime-arrivalTime[shortestJob]-currentBurst)
		completionTime = append(completionTime, currentTime)
		turnAroundTime = append(turnAroundTime, completionTime[i]-arrivalTime[i])
	}

	// Calculate average waiting time and turnaround time
	totalWaitingTime := 0
	totalTurnAroundTime := 0
	for _, wt := range waitingTime {
		totalWaitingTime += wt
	}
	for _, tat := range turnAroundTime {
		totalTurnAroundTime += tat
	}
	avgWaitingTime := float64(totalWaitingTime) / float64(len(processID))
	avgTurnAroundTime := float64(totalTurnAroundTime) / float64(len(processID))

	// Print results similar to FCFS
	fmt.Println("+-------+------------+-----------+------------+--------------+--------------+")
	fmt.Println("| PID  | AT         | BT         | CT         | WT           | TAT          |")
	fmt.Println("+-------+------------+-----------+------------+--------------+--------------+")
	for i := range processID {
		fmt.Printf("| %4s | %10d | %10d | %10d | %12d | %12d |\n", processID[i], arrivalTime[i], burstTime[i], completionTime[i], waitingTime[i], turnAroundTime[i])
	}
	fmt.Println("+-------+------------+-----------+------------+--------------+--------------+")
	fmt.Printf("Average Waiting Time: %.2f\n", avgWaitingTime)
	fmt.Printf("Average Turnaround Time: %.2f\n", avgTurnAroundTime)

	// Print Gantt chart using outputGantt function
	gantt := SJFGantt(processID, completionTime)
	outputGantt(os.Stdout, gantt)
}
