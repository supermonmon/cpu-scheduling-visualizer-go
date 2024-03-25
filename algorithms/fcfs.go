package algorithms

import (
	"cpu-scheduling-algorithms/models"
	"fmt"
	"sort"
)

type FCFSResult struct {
	Processes      []*models.Process
	AvgTAT         float64
	AvgWT          float64
	CPUUtilization float64
	GanttChart     []int
}

func FirstComeFirstServe() (*FCFSResult, error) {
	var processes []*models.Process
	var arrivalTimes, burstTimes []int
	var numProcesses int

	fmt.Print("Enter number of processes: ")
	_, err := fmt.Scan(&numProcesses)
	if err != nil {
		return nil, fmt.Errorf("\x1b[31mFailed to read number of processes: %v\x1b[0m", err)
	}

	arrivalTimes = make([]int, numProcesses)
	burstTimes = make([]int, numProcesses)

	// Get arrival times from user
	fmt.Print("Enter arrival times separated by spaces: ")
	for i := 0; i < numProcesses; i++ {
		_, err := fmt.Scan(&arrivalTimes[i])
		if err != nil {
			return nil, fmt.Errorf("\x1b[31mFailed to read arrival time for process %d: %v\x1b[0m", i+1, err)
		}
	}

	// Get burst times from user
	fmt.Print("Enter burst times separated by spaces: ")
	for i := 0; i < numProcesses; i++ {
		_, err := fmt.Scan(&burstTimes[i])
		if err != nil {
			return nil, fmt.Errorf("\x1b[31mFailed to read burst time for process %d: %v\x1b[0m", i+1, err)
		}
	}

	// Create processes
	for i := 0; i < numProcesses; i++ {
		newProcess := &models.Process{
			ID:          i + 1,
			ArrivalTime: arrivalTimes[i],
			BurstTime:   burstTimes[i],
		}
		processes = append(processes, newProcess)
	}

	// Sort processes by arrival time
	sort.Slice(processes, func(i, j int) bool {
		return processes[i].ArrivalTime < processes[j].ArrivalTime
	})

	// Calculate completion time, turnaround time, and waiting time for each process
	var ganttChart []int
	var totalTAT, totalWT float64

	currentTime := 0
	for _, process := range processes {
		// Update current time to match process arrival time if needed
		if process.ArrivalTime > currentTime {
			currentTime = process.ArrivalTime
		}
		ganttChart = append(ganttChart, process.ID)

		process.CompletionTime = currentTime + process.BurstTime
		process.TurnaroundTime = process.CompletionTime - process.ArrivalTime
		process.WaitingTime = process.TurnaroundTime - process.BurstTime

		totalTAT += float64(process.TurnaroundTime)
		totalWT += float64(process.WaitingTime)

		currentTime = process.CompletionTime
	}

	// Calculate AVG TAT, AVG WT, and CPU Utilization
	avgTAT := totalTAT / float64(numProcesses)
	avgWT := totalWT / float64(numProcesses)
	cpuUtilization := (float64(currentTime) / float64(currentTime)) * 100 // Total burst time / Total time

	result := &FCFSResult{
		Processes:      processes,
		AvgTAT:         avgTAT,
		AvgWT:          avgWT,
		CPUUtilization: cpuUtilization,
		GanttChart:     ganttChart,
	}

	return result, nil
}
