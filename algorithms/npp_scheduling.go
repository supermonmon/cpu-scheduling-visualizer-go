package algorithms

import (
	"fmt"
	"os"
)

// NPP implements the Priority scheduling algorithm (non-preemptive)
func NPP(processID []string, arrivalTime, burstTime, priority []int) {
	fmt.Println("Priority Scheduling (Non-Preemptive) Results:")

	type NPPData struct {
		pid         string
		arrivalTime int
		burstTime   int
		priority    int
		completed   bool
	}

	// Create a data structure to store process information with burst time, priority, and completion flag
	NPPDataSlice := make([]NPPData, len(processID))
	for i := range processID {
		NPPDataSlice[i] = NPPData{processID[i], arrivalTime[i], burstTime[i], priority[i], false}
	}

	// Initialize slices for waiting time, completion time, and turnaround time
	waitingTime := make([]int, len(processID))
	completionTime := make([]int, len(processID))
	turnAroundTime := make([]int, len(processID))

	var currentTime int = 0        // Track current time
	var completedProcesses int = 0 // Track number of completed processes

	var gantt []TimeSlice

	for completedProcesses < len(processID) {
		// Find the highest priority process that has arrived
		highestPriorityJob := -1
		for j := 0; j < len(processID); j++ {
			if !NPPDataSlice[j].completed && NPPDataSlice[j].arrivalTime <= currentTime {
				if highestPriorityJob == -1 || NPPDataSlice[j].priority < NPPDataSlice[highestPriorityJob].priority {
					highestPriorityJob = j
				}
			}
		}

		// No available processes at current time
		if highestPriorityJob == -1 {
			currentTime++
			continue
		}

		// Execute the selected process
		NPPDataSlice[highestPriorityJob].completed = true
		completionTime[highestPriorityJob] = currentTime + NPPDataSlice[highestPriorityJob].burstTime
		turnAroundTime[highestPriorityJob] = completionTime[highestPriorityJob] - NPPDataSlice[highestPriorityJob].arrivalTime
		waitingTime[highestPriorityJob] = turnAroundTime[highestPriorityJob] - NPPDataSlice[highestPriorityJob].burstTime

		// Update the number of completed processes
		completedProcesses++

		// Update Gantt chart
		ganttSlice := TimeSlice{PID: NPPDataSlice[highestPriorityJob].pid, Start: currentTime, Stop: completionTime[highestPriorityJob]}
		gantt = append(gantt, ganttSlice)

		// Move current time to the completion time of the executed process
		currentTime = completionTime[highestPriorityJob]
	}

	// Print results
	fmt.Println("+-------+------------+-----------+-----------+------------+--------------+--------------+")
	fmt.Println("| PID   | AT         | BT        | PL        | CT         | WT           | TAT          |")
	fmt.Println("+-------+------------+-----------+-----------+------------+--------------+--------------+")
	totalWT := 0
	totalTAT := 0
	for i := range processID {
		fmt.Printf("| %4s  | %10d | %9d | %9d | %10d | %12d | %12d |\n", processID[i], arrivalTime[i], burstTime[i], priority[i], completionTime[i], waitingTime[i], turnAroundTime[i])
		totalWT += waitingTime[i]
		totalTAT += turnAroundTime[i]
	}
	fmt.Println("+-------+------------+-----------+-----------+------------+--------------+--------------+")

	// Calculate average waiting time and turnaround time
	avgWaitingTime := float64(totalWT) / float64(len(processID))
	avgTurnAroundTime := float64(totalTAT) / float64(len(processID))

	// Print average waiting time and turnaround time
	fmt.Printf("Average Waiting Time: %.2f\n", avgWaitingTime)
	fmt.Printf("Average Turnaround Time: %.2f\n", avgTurnAroundTime)

	// Print Gantt chart
	outputGantt(os.Stdout, gantt)
}
