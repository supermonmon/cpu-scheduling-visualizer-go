package algorithms

func SRTF(processID []string, arrivalTime, burstTime []int) Result {

	type SRTFData struct {
		pid         string
		arrivalTime int
		burstTime   int
		remaining   int
		completed   bool
	}

	var totalTime int
	for _, bt := range burstTime {
		totalTime += bt
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

	// Calculate average waiting time and turnaround time
	totalWT := 0
	totalTAT := 0
	for i := range processID {
		totalWT += waitingTime[i]
		totalTAT += turnAroundTime[i]
	}
	avgWaitingTime := float64(totalWT) / float64(len(processID))
	avgTurnAroundTime := float64(totalTAT) / float64(len(processID))
	cpuUtilization := float64(totalTime) / float64(currentTime) * 100

	// Return the result
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
