package algorithms

import (
	"sort"
)

// FCFS implements the First-Come, First-Served (FCFS) scheduling algorithm
func FCFS(processID []string, arrivalTime, burstTime []int) FCFSResult {

	type FCFSData struct {
		pid         string
		arrivalTime int
		burstTime   int
	}

	var totalTime int
	for _, bt := range burstTime {
		totalTime += bt
	}

	FCFSDataSlice := make([]FCFSData, len(processID))
	for i := range processID {
		FCFSDataSlice[i] = FCFSData{processID[i], arrivalTime[i], burstTime[i]}
	}
	sort.Slice(FCFSDataSlice, func(i, j int) bool {
		return FCFSDataSlice[i].arrivalTime < FCFSDataSlice[j].arrivalTime
	})

	// Extract sorted data back to original slices
	for i := range FCFSDataSlice {
		processID[i] = FCFSDataSlice[i].pid
		arrivalTime[i] = FCFSDataSlice[i].arrivalTime
		burstTime[i] = FCFSDataSlice[i].burstTime
	}

	var waitingTime, completionTime, turnAroundTime []int
	var current int = 0 // Current time

	for i := range processID {
		// Handle processes that arrive after current time
		if arrivalTime[i] > current {
			current = arrivalTime[i]
		}
		waitingTime = append(waitingTime, current-arrivalTime[i])
		current += burstTime[i]
		completionTime = append(completionTime, current)
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
	cpuUtilization := float64(totalTime) / float64(current) * 100

	// Print Gantt chart using outputGantt function
	gantt := FCFSGantt(processID, burstTime)

	return FCFSResult{
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

func FCFSGantt(processID []string, burstTime []int) []TimeSlice { // Updated processID type
	gantt := make([]TimeSlice, len(processID))
	var current int = 0
	for i := range processID {
		gantt[i] = TimeSlice{PID: processID[i], Start: current} // Use processID as string
		current += burstTime[i]
		gantt[i].Stop = current
	}
	return gantt
}
