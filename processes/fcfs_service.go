package processes

import (
	"cpu-scheduling-algorithms/algorithms"
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"
)

func RunFCFS() {
	// Call the FirstComeFirstServe function from the algorithms package
	result, err := algorithms.FirstComeFirstServe()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	processes := result.Processes
	ganttChart := result.GanttChart
	avgTAT := result.AvgTAT
	avgWT := result.AvgWT
	cpuUtilization := result.CPUUtilization

	// Create a new table for FCFS results
	t := table.NewWriter()
	t.SetTitle("First Come First Serve Algorithm")
	t.Style().Options.SeparateColumns = true
	t.Style().Options.SeparateHeader = true

	// Set column headers
	t.AppendHeader(table.Row{"Process ID", "Arrival Time", "Burst Time", "Completion Time", "Turnaround Time", "Waiting Time"})

	// Add rows for each process
	for _, process := range processes {
		t.AppendRow(table.Row{
			process.ID,
			process.ArrivalTime,
			process.BurstTime,
			process.CompletionTime,
			process.TurnaroundTime,
			process.WaitingTime,
		})
	}

	// Set average TAT, average WT, and CPU Utilization as footer
	t.AppendFooter(table.Row{"", "", "", "", fmt.Sprintf("Avg TAT: %.2f", avgTAT), fmt.Sprintf("Avg WT: %.2f", avgWT), fmt.Sprintf("CPU Utilization: %.2f%%", cpuUtilization)})

	// Render the FCFS results table
	fmt.Println(t.Render())

	// Print Gantt chart
	fmt.Println("\n\nGantt Chart:")
	fmt.Print("┌")
	for range ganttChart {
		fmt.Print("──────┬")
	}
	fmt.Println()

	fmt.Print("│ ")
	for _, id := range ganttChart {
		fmt.Printf("P%d    │ ", id)
	}
	fmt.Println()

	fmt.Print("└")
	for range ganttChart {
		fmt.Print("──────┴")
	}
	fmt.Println()

	fmt.Print("0")
	for _, process := range processes {
		fmt.Printf("     %d", process.CompletionTime)
	}
	fmt.Println()
}
