package algorithms

import "sort"

func RR(processID []string, arrivalTime, burstTime []int, timeQuantum int) Result {

	var (
		completionTime = make([]int, len(processID))
		waitingTime    = make([]int, len(processID))
		turnaroundTime = make([]int, len(processID))
		remainingTime  = make([]int, len(processID))
		currentTime    int
		totalProcesses = len(processID)
		totalCompleted int
	)

	var totalTime int
	for _, bt := range burstTime {
		totalTime += bt
	}

	// Initialize remainingTime array with burstTime
	copy(remainingTime, burstTime)

	// Sort processes based on arrival time
	type Process struct {
		ID           string
		ArrivalTime  int
		BurstTime    int
		RemainingTime int
	}
	var processes []Process
	for i := 0; i < totalProcesses; i++ {
		processes = append(processes, Process{
			ID:           processID[i],
			ArrivalTime:  arrivalTime[i],
			BurstTime:    burstTime[i],
			RemainingTime: burstTime[i],
		})
	}
	sort.Slice(processes, func(i, j int) bool {
		return processes[i].ArrivalTime < processes[j].ArrivalTime
	})

	// Initialize Gantt chart slices
	var gantt []TimeSlice

	// Iterate until all processes are completed
	for totalCompleted < totalProcesses {
		// Iterate through each process
		for i := 0; i < totalProcesses; i++ {
			// If remaining time for process i is greater than 0 and it has arrived
			if processes[i].RemainingTime > 0 && processes[i].ArrivalTime <= currentTime {
				// Execute the process for the time quantum or the remaining time, whichever is smaller
				executionTime := min(timeQuantum, processes[i].RemainingTime)
				currentTime += executionTime
				processes[i].RemainingTime -= executionTime

				// Update Gantt chart
				ganttSlice := TimeSlice{PID: processes[i].ID, Start: currentTime - executionTime, Stop: currentTime}
				gantt = append(gantt, ganttSlice)

				// Check if the process is completed
				if processes[i].RemainingTime == 0 {
					totalCompleted++
					completionTime[i] = currentTime
					turnaroundTime[i] = completionTime[i] - arrivalTime[i]
					waitingTime[i] = turnaroundTime[i] - burstTime[i]
				}
			}
		}
	}

	totalWaitingTime := 0
	totalTurnaroundTime := 0
	for i := 0; i < totalProcesses; i++ {
		totalWaitingTime += waitingTime[i]
		totalTurnaroundTime += turnaroundTime[i]
	}

	avgWaitingTime := float64(totalWaitingTime) / float64(totalProcesses)
	avgTurnaroundTime := float64(totalTurnaroundTime) / float64(totalProcesses)
	cpuUtilization := float64(totalTime) / float64(currentTime) * 100

	return Result{
		Algorithm:         "RR",
		ProcessID:         processID,
		ArrivalTime:       arrivalTime,
		BurstTime:         burstTime,
		TimeQuantum:       timeQuantum,
		CompletionTime:    completionTime,
		WaitingTime:       waitingTime,
		TurnAroundTime:    turnaroundTime,
		GanttChart:        gantt,
		AvgWaitingTime:    avgWaitingTime,
		AvgTurnAroundTime: avgTurnaroundTime,
		CPUUtilization:    cpuUtilization,
	}
}

// Helper function to find minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
