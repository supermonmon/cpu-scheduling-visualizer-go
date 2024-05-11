package algorithms

import (
	"sort"
)

func SJF(processID []string, arrivalTime, burstTime []int) Result{

	var totalTime int
	for _, bt := range burstTime {
		totalTime += bt
	}

	SJFDataSlice := make([]SJFData, len(processID))
	for i := range processID {
		SJFDataSlice[i] = SJFData{processID[i], arrivalTime[i], burstTime[i]}
	}

	// Sort processes by arrival time (ascending)
	sort.Slice(SJFDataSlice, func(i, j int) bool {
		return SJFDataSlice[i].arrivalTime < SJFDataSlice[j].arrivalTime
	})

	// Extract sorted data back to original slices
	for i := range SJFDataSlice {
		processID[i] = SJFDataSlice[i].pid
		arrivalTime[i] = SJFDataSlice[i].arrivalTime
		burstTime[i] = SJFDataSlice[i].burstTime
	}

	var waitingTime, completionTime, turnAroundTime []int
	var currentProcess int = -1              
	var currentBurst, currentTime int = 0, 0 

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

		waitingTime = append(waitingTime, currentTime-arrivalTime[shortestJob]-currentBurst)
		completionTime = append(completionTime, currentTime)
		turnAroundTime = append(turnAroundTime, completionTime[i]-arrivalTime[i])
	}

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
	cpuUtilization := float64(totalTime) / float64(currentTime) * 100

	gantt := SJFGantt(processID, completionTime)

	return Result{
		Algorithm:		   "SJF",
		ProcessID:         processID,
		ArrivalTime:       arrivalTime,
		BurstTime:         burstTime,
		CompletionTime:    completionTime,
		WaitingTime:       waitingTime,
		TurnAroundTime:    turnAroundTime,
		AvgWaitingTime:    avgWaitingTime,
		AvgTurnAroundTime: avgTurnAroundTime,
		GanttChart:        gantt,
		CPUUtilization:    cpuUtilization,
	}
}

func SJFGantt(processID []string, completionTime []int) []TimeSlice {
	gantt := make([]TimeSlice, len(processID))
	var prevCompletionTime int = 0
	for i := range processID {
		gantt[i] = TimeSlice{PID: processID[i], Start: prevCompletionTime}
		prevCompletionTime = completionTime[i]
		gantt[i].Stop = prevCompletionTime
	}
	return gantt
}
