package algorithms

// Import necessary packages

// Define the SRTF function
func SRTF(processID []string, arrivalTime, burstTime []int) Result {
	// Initialize variables
	var totalTime int
	for _, bt := range burstTime {
		totalTime += bt
	}

	SRTFDataSlice := make([]SRTFData, len(processID))
	for i := range processID {
		SRTFDataSlice[i] = SRTFData{processID[i], arrivalTime[i], burstTime[i], burstTime[i], false}
	}
	waitingTime := make([]int, len(processID))
	completionTime := make([]int, len(processID))
	turnAroundTime := make([]int, len(processID))

	var currentTime int = 0
	var completedProcesses int

	var prevPID string
	var startTime int

	var gantt []TimeSlice

	// Loop until all processes are completed
	for completedProcesses < len(processID) {
		// Find the available process with the shortest remaining burst time among arrived processes
		shortestJob := -1
		for j := 0; j < len(processID); j++ {
			// Check if the process has arrived and not completed
			if !SRTFDataSlice[j].completed && SRTFDataSlice[j].arrivalTime <= currentTime {
				// If no shortest job found yet or the current job has a shorter remaining time, update shortestJob
				if shortestJob == -1 || SRTFDataSlice[j].remaining < SRTFDataSlice[shortestJob].remaining {
					shortestJob = j
				}
			}
		}

		// No available processes at the current time
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
			prevPID = ""
		}
	}

	// Calculate total waiting time and total turnaround time
	totalWT := 0
	totalTAT := 0
	for i := range processID {
		totalWT += waitingTime[i]
		totalTAT += turnAroundTime[i]
	}

	// Calculate average waiting time, average turnaround time, and CPU utilization
	avgWaitingTime := float64(totalWT) / float64(len(processID))
	avgTurnAroundTime := float64(totalTAT) / float64(len(processID))
	cpuUtilization := float64(totalTime) / float64(currentTime) * 100

	// Create and return the result
	return Result{
		Algorithm:         "SRTF",
		ProcessID:         processID,
		ArrivalTime:       arrivalTime,
		BurstTime:         burstTime,
		CompletionTime:    completionTime,
		WaitingTime:       waitingTime,
		TurnAroundTime:    turnAroundTime,
		GanttChart:        gantt,
		AvgWaitingTime:    avgWaitingTime,
		AvgTurnAroundTime: avgTurnAroundTime,
		CPUUtilization:    cpuUtilization,
	}
}
