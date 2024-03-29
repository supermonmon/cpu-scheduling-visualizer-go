package algorithms

import (
	"fmt"
	"os"
)

func DisplayFCFS(result FCFSResult) {

	processID := result.ProcessID
	arrivalTime := result.ArrivalTime
	burstTime := result.BurstTime

	processedProcessID := make([]string, len(processID))
	copy(processedProcessID, processID)
	processedArrivalTime := make([]int, len(arrivalTime))
	copy(processedArrivalTime, arrivalTime)
	processedBurstTime := make([]int, len(burstTime))
	copy(processedBurstTime, burstTime)

	for i := range result.CompletionTime {
		processedProcessID[i] = fmt.Sprintf("P%d", i+1)
		processedArrivalTime[i] = i * 5 // Modify this as needed
		processedBurstTime[i] = result.CompletionTime[i] - arrivalTime[i]
	}

	fmt.Println("+-----------------------------------------------------------------------------+")
	fmt.Println("\n\033[48;5;24;38;5;15m⚙️  First Come First Serve Scheduling \033[0m")
	fmt.Print("\n")

	fmt.Println("+-------+------------+-----------+------------+--------------+--------------+")
	fmt.Println("| PID  | AT         | BT         | CT         | WT           | TAT          |")
	fmt.Println("+-------+------------+-----------+------------+--------------+--------------+")
	for i := range processID {
		fmt.Printf("| %4s | %10d | %10d | %10d | %12d | %12d |\n", processID[i], arrivalTime[i], burstTime[i], result.CompletionTime[i], result.WaitingTime[i], result.TurnAroundTime[i])
	}
	fmt.Println("+-------+------------+-----------+------------+--------------+--------------+")

	fmt.Printf("\nAverage Waiting Time: \033[20;5;35m%.2f\033[0m\n", result.AvgWaitingTime)
	fmt.Printf("Average Turnaround Time: \033[20;5;35m%.2f\033[0m\n", result.AvgTurnAroundTime)
	fmt.Printf("\n")

	outputGantt(os.Stdout, result.GanttChart)
}

func DisplaySJF(result SJFResult) {

	processID := result.ProcessID
	arrivalTime := result.ArrivalTime
	burstTime := result.BurstTime

	fmt.Println("+-----------------------------------------------------------------------------+")
	fmt.Println("\n\033[48;5;24;38;5;15m⚙️  Shortest Job First Scheduling \033[0m")
	fmt.Print("\n")

	fmt.Println("+-------+------------+-----------+------------+--------------+--------------+")
	fmt.Println("| PID  | AT         | BT         | CT         | WT           | TAT          |")
	fmt.Println("+-------+------------+-----------+------------+--------------+--------------+")
	for i := range processID {
		fmt.Printf("| %4s | %10d | %10d | %10d | %12d | %12d |\n", processID[i], arrivalTime[i], burstTime[i], result.CompletionTime[i], result.WaitingTime[i], result.TurnAroundTime[i])
	}
	fmt.Println("+-------+------------+-----------+------------+--------------+--------------+")

	fmt.Printf("\nAverage Waiting Time: \033[20;5;35m%.2f\033[0m\n", result.AvgWaitingTime)
	fmt.Printf("Average Turnaround Time: \033[20;5;35m%.2f\033[0m\n", result.AvgTurnAroundTime)
	fmt.Printf("\n")

	outputGantt(os.Stdout, result.GanttChart)
}

func DisplaySRTF(result SRTFResult) {

	fmt.Println("+-----------------------------------------------------------------------------+")
	fmt.Println("\n\033[48;5;24;38;5;15m⚙️  Shortest Remaining Time First Scheduling \033[0m")
	fmt.Print("\n")

	fmt.Println("+-------+------------+-----------+------------+--------------+--------------+")
	fmt.Println("| PID   | AT         | BT        | CT         | WT           | TAT          |")
	fmt.Println("+-------+------------+-----------+------------+--------------+--------------+")
	for i := range result.ProcessID {
		fmt.Printf("| %4s  | %10d | %9d | %10d | %12d | %12d |\n", result.ProcessID[i], result.ArrivalTime[i], result.BurstTime[i], result.CompletionTime[i], result.WaitingTime[i], result.TurnAroundTime[i])
	}
	fmt.Println("+-------+------------+-----------+------------+--------------+--------------+")

	fmt.Printf("\nAverage Waiting Time: \033[20;5;35m%.2f\033[0m\n", result.AvgWaitingTime)
	fmt.Printf("Average Turnaround Time: \033[20;5;35m%.2f\033[0m\n", result.AvgTurnAroundTime)
	fmt.Printf("\n")

	outputGantt(os.Stdout, result.GanttChart)
}

func DisplayNPP(result NPPResult) {
	fmt.Println("+-----------------------------------------------------------------------------+")
	fmt.Println("\n\033[48;5;24;38;5;15m⚙️  Priority Scheduling \033[0m")
	fmt.Print("\n")

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

	// Calculate average waiting time and turnaround time
	avgWaitingTime := result.AvgWaitingTime
	avgTurnAroundTime := result.AvgTurnAroundTime

	fmt.Printf("\nAverage Waiting Time: \033[20;5;35m%.2f\033[0m\n", avgWaitingTime)
	fmt.Printf("Average Turnaround Time: \033[20;5;35m%.2f\033[0m\n", avgTurnAroundTime)
	fmt.Printf("\n")

	// Print Gantt chart
	outputGantt(os.Stdout, result.GanttChart)
}

func DisplayRR(result RRResult) {

	fmt.Println("+-----------------------------------------------------------------------------+")
	fmt.Println("\n\033[48;5;24;38;5;15m⚙️  Round Robin Scheduling \033[0m")
	fmt.Print("\n")

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

	// Calculate average waiting time and turnaround time
	avgWaitingTime := result.AvgWaitingTime
	avgTurnAroundTime := result.AvgTurnAroundTime

	fmt.Printf("\nAverage Waiting Time: \033[20;5;35m%.2f\033[0m\n", avgWaitingTime)
	fmt.Printf("Average Turnaround Time: \033[20;5;35m%.2f\033[0m\n", avgTurnAroundTime)
	fmt.Printf("\n")

	// Print Gantt chart
	outputGantt(os.Stdout, result.GanttChart)
}
