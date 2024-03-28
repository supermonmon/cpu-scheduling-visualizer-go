package algorithms

import (
	"fmt"
)

func RR(processID []string, arrivalTime, burstTime []int, timeQuantum int) {
	fmt.Println("Round Robin Scheduling Results (Time Quantum:", timeQuantum, "):")

	// Variables to track completion time, waiting time, and turnaround time
	var (
		completionTime = make([]int, len(processID))
		waitingTime    = make([]int, len(processID))
		turnaroundTime = make([]int, len(processID))
		remainingTime  = make([]int, len(processID)) // Track remaining time for each process
		currentTime    int                           // Track current time
		totalProcesses = len(processID)
		totalCompleted int // Track the number of completed processes
	)

	// Initialize remainingTime array with burstTime
	copy(remainingTime, burstTime)

	// Iterate until all processes are completed
	for totalCompleted < totalProcesses {
		// Iterate through each process
		for i := 0; i < totalProcesses; i++ {
			// If remaining time for process i is greater than 0
			if remainingTime[i] > 0 {
				// Execute the process for the time quantum or the remaining time, whichever is smaller
				if remainingTime[i] <= timeQuantum {
					currentTime += remainingTime[i]
					completionTime[i] = currentTime
					remainingTime[i] = 0
				} else {
					currentTime += timeQuantum
					remainingTime[i] -= timeQuantum
				}

				// Check if the process is completed
				if remainingTime[i] == 0 {
					totalCompleted++
					turnaroundTime[i] = completionTime[i] - arrivalTime[i]
					waitingTime[i] = turnaroundTime[i] - burstTime[i]
				}
			}
		}
	}

	// Print the table
	fmt.Println("+-------+------------+-----------+------------+--------------+--------------+")
	fmt.Println("| PID   | AT         | BT        | CT         | WT           | TAT          |")
	fmt.Println("+-------+------------+-----------+------------+--------------+--------------+")
	totalWaitingTime := 0
	totalTurnaroundTime := 0
	for i := 0; i < totalProcesses; i++ {
		fmt.Printf("| %4s  | %10d | %9d | %10d | %12d | %12d |\n", processID[i], arrivalTime[i], burstTime[i], completionTime[i], waitingTime[i], turnaroundTime[i])
		totalWaitingTime += waitingTime[i]
		totalTurnaroundTime += turnaroundTime[i]
	}
	fmt.Println("+-------+------------+-----------+------------+--------------+--------------+")

	// Calculate and display average waiting time, turnaround time, and completion time
	avgWaitingTime := float64(totalWaitingTime) / float64(totalProcesses)
	avgTurnaroundTime := float64(totalTurnaroundTime) / float64(totalProcesses)

	fmt.Printf("Average Waiting Time: %.2f\n", avgWaitingTime)
	fmt.Printf("Average Turnaround Time: %.2f\n", avgTurnaroundTime)
}
