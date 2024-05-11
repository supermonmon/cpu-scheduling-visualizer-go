package algorithms

import (
	"fmt"
	"os"
)


func Display(result Result) {

	fmt.Println("+-----------------------------------------------------------------------------+")
	fmt.Printf("\n\033[48;5;24;38;5;15m⚙️  %s Scheduling \033[0m\n", result.Algorithm)
	fmt.Print("\n")
	
	if result.Algorithm == "NPP"{
		fmt.Println("+-------+------------+-----------+-----------+------------+--------------+--------------+")
		fmt.Println("| PID   | AT         | BT        | PL        | CT         | WT           | TAT          |")
		fmt.Println("+-------+------------+-----------+-----------+------------+--------------+--------------+")
		totalWT := 0
		totalTAT := 0
		for i := range result.ProcessID {
			fmt.Printf("| %4s  | %10d | %9d | %9d | %10d | %12d | %12d |\n", result.ProcessID[i], result.ArrivalTime[i], result.BurstTime[i], result.Priority[i], result.CompletionTime[i], result.WaitingTime[i], result.TurnAroundTime[i])
			totalWT += result.WaitingTime[i]
			totalTAT += result.TurnAroundTime[i]
		}
		fmt.Println("+-------+------------+-----------+-----------+------------+--------------+--------------+")
	} else if (result.Algorithm == "RR") {
		fmt.Println("+-------+------------+-----------+------------+--------------+--------------+")
		fmt.Println("| PID   | AT         | BT        | CT         | WT           | TAT          |")
		fmt.Println("+-------+------------+-----------+------------+--------------+--------------+")
		totalWT := 0
		totalTAT := 0
		for i := range result.ProcessID {
			fmt.Printf("| %4s  | %10d | %9d | %10d | %12d | %12d |\n", result.ProcessID[i], result.ArrivalTime[i], result.BurstTime[i], result.CompletionTime[i], result.WaitingTime[i], result.TurnAroundTime[i])
			totalWT += result.WaitingTime[i]
			totalTAT += result.TurnAroundTime[i]
		}
		fmt.Println("+-------+------------+-----------+-----------+------------+--------------+--------------+")
	} else {
		fmt.Println("+-------+------------+-----------+------------+--------------+--------------+")
		fmt.Println("| PID  | AT         | BT         | CT         | WT           | TAT          |")
		fmt.Println("+-------+------------+-----------+------------+--------------+--------------+")
		for i := range result.ProcessID {
			fmt.Printf("| %4s | %10d | %10d | %10d | %12d | %12d |\n", result.ProcessID[i], result.ArrivalTime[i], result.BurstTime[i], result.CompletionTime[i], result.WaitingTime[i], result.TurnAroundTime[i])
		}
		fmt.Println("+-------+------------+-----------+------------+--------------+--------------+")
	}

	fmt.Printf("\nAverage Waiting Time: \033[20;5;35m%.2f\033[0m\n", result.AvgWaitingTime)
	fmt.Printf("Average Turnaround Time: \033[20;5;35m%.2f\033[0m\n", result.AvgTurnAroundTime)
	fmt.Printf("CPU Utilization: \033[20;5;35m%.2f\033[0m\n", result.CPUUtilization)
	fmt.Printf("\n")

	outputGantt(os.Stdout, result.GanttChart)
}


