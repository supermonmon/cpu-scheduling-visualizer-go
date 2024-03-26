package algorithms

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

// FCFS implements the First-Come, First-Served (FCFS) scheduling algorithm
func FCFS(processID []string, arrivalTime, burstTime []int) {
	fmt.Println("FCFS Scheduling Results:")

	// Sort data based on arrival time (ascending order)
	type processData struct {
		pid         string
		arrivalTime int
		burstTime   int
	}
	processDataSlice := make([]processData, len(processID))
	for i := range processID {
		processDataSlice[i] = processData{processID[i], arrivalTime[i], burstTime[i]}
	}
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
	var current int = 0 // Current time

	// Print Gantt chart using outputGantt function
	gantt := createGanttChart(processID, burstTime)
	outputGantt(os.Stdout, gantt)

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

	fmt.Println("+-------+------------+-----------+------------+--------------+--------------+")
	fmt.Println("| PID  | AT         | BT         | CT         | WT           | TAT          |")
	fmt.Println("+-------+------------+-----------+------------+--------------+--------------+")
	for i := range processID {
		fmt.Printf("| %4s | %10d | %10d | %10d | %12d | %12d |\n", processID[i], arrivalTime[i], burstTime[i], completionTime[i], waitingTime[i], turnAroundTime[i])
	}
	fmt.Println("+-------+------------+-----------+------------+--------------+--------------+")
	fmt.Printf("Average Waiting Time: %.2f\n", avgWaitingTime)
	fmt.Printf("Average Turnaround Time: %.2f\n", avgTurnAroundTime)
}

// Function to create Gantt chart data structure
func createGanttChart(processID []string, burstTime []int) []TimeSlice { // Updated processID type
	gantt := make([]TimeSlice, len(processID))
	var current int = 0
	for i := range processID {
		gantt[i] = TimeSlice{PID: processID[i], Start: current} // Use processID as string
		current += burstTime[i]
		gantt[i].Stop = current
	}
	return gantt
}

// Structure to represent a time slice in the Gantt chart
type TimeSlice struct {
	PID   string
	Start int
	Stop  int
}

// Function to print the Gantt chart with improved formatting
func outputGantt(w io.Writer, gantt []TimeSlice) {
	_, _ = fmt.Fprintln(w, "Gantt schedule")
	_, _ = fmt.Fprint(w, "|")
	maxLength := 0
	for _, ts := range gantt {
		pidStr := fmt.Sprint(ts.PID)
		if len(pidStr) > maxLength {
			maxLength = len(pidStr)
		}
	}
	format := strings.Repeat(" ", maxLength+2) + "%v |"
	for _, ts := range gantt {
		_, _ = fmt.Fprintf(w, format, ts.PID)
	}
	_, _ = fmt.Fprintln(w)

	_, _ = fmt.Fprint(w, "0")
	for _, ts := range gantt {
		_, _ = fmt.Fprintf(w, "\t%d", ts.Stop)
	}
	_, _ = fmt.Fprintln(w)
	_, _ = fmt.Fprintf(w, "\n\n")
}
