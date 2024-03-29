package algorithms

import (
	"fmt"
	"os"
)

// SRTF implements the Shortest Remaining Time First (SRTF) scheduling algorithm with preemptive execution
func SRTF(processID []string, arrivalTime, burstTime []int) {
	fmt.Println("+-----------------------------------------------------------------------------+")
	fmt.Println("\n\033[48;5;24;38;5;15m⚙️  Shortest Remaining Time First Scheduling \033[0m\n")

	type SRTFData struct {
		pid         string
		arrivalTime int
		burstTime   int
		remaining   int
		completed   bool
	}

	// Create a data structure to store process information with burst time, remaining burst time, and completion flag
	SRTFDataSlice := make([]SRTFData, len(processID))
	for i := range processID {
		SRTFDataSlice[i] = SRTFData{processID[i], arrivalTime[i], burstTime[i], burstTime[i], false}
	}

	// Initialize slices for waiting time, completion time, and turnaround time
	waitingTime := make([]int, len(processID))
	completionTime := make([]int, len(processID))
	turnAroundTime := make([]int, len(processID))

	var currentTime int = 0    // Track current time
	var completedProcesses int // Track number of completed processes

	var prevPID string
	var startTime int

	var gantt []TimeSlice

	for completedProcesses < len(processID) {
		// Find the available process with the shortest remaining burst time
		shortestJob := -1
		for j := 0; j < len(processID); j++ {
			if !SRTFDataSlice[j].completed && SRTFDataSlice[j].arrivalTime <= currentTime {
				if shortestJob == -1 || SRTFDataSlice[j].remaining < SRTFDataSlice[shortestJob].remaining {
					shortestJob = j
				}
			}
		}

		// No available processes at current time
		if shortestJob == -1 {
			currentTime++
			continue
		}

		// Check if the process has changed
		if prevPID != SRTFDataSlice[shortestJob].pid {
			if prevPID != "" {
				// Update the stop time for the previous process in the Gantt chart
				ganttSlice := TimeSlice{PID: prevPID, Start: startTime, Stop: currentTime}
				gantt = append(gantt, ganttSlice)
			}
			// Update the start time and previous PID
			prevPID = SRTFDataSlice[shortestJob].pid
			startTime = currentTime
		}

		// Execute the process for 1 unit of time
		SRTFDataSlice[shortestJob].remaining--
		currentTime++

		// Check if the process is completed
		if SRTFDataSlice[shortestJob].remaining == 0 {
			completedProcesses++
			SRTFDataSlice[shortestJob].completed = true
			completionTime[shortestJob] = currentTime
			turnAroundTime[shortestJob] = completionTime[shortestJob] - SRTFDataSlice[shortestJob].arrivalTime
			waitingTime[shortestJob] = turnAroundTime[shortestJob] - burstTime[shortestJob]

			// Update the stop time for the completed process in the Gantt chart
			ganttSlice := TimeSlice{PID: SRTFDataSlice[shortestJob].pid, Start: startTime, Stop: currentTime}
			gantt = append(gantt, ganttSlice)
			prevPID = "" // Reset previous PID
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

	fmt.Printf("\nAverage Waiting Time: \033[20;5;35m%.2f\033[0m\n", avgWaitingTime)
	fmt.Printf("Average Turnaround Time: \033[20;5;35m%.2f\033[0m\n", avgTurnAroundTime)
	fmt.Printf("\n")

	// Print Gantt chart using outputGantt function
	outputGantt(os.Stdout, gantt)
}
