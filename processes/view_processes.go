package processes

import (
	"cpu-scheduling-algorithms/models"
	"fmt"
)

func DisplayProcesses(processes []models.Process) {

	if len(processes) == 0 {
		fmt.Println("\n\n\033[1;31mThere are no processes!\033[0m")
		return
	}

	fmt.Println("\n\n")
	fmt.Println("-----------------------------------------------------------------")
	fmt.Println("| Process ID | Arrival Time | Burst Time | Completion Time |")
	fmt.Println("-----------------------------------------------------------------")
	for _, process := range processes {
		fmt.Printf("|     %2d     |      %8d |      %7d |        %10d |\n", process.ID, process.ArrivalTime, process.BurstTime, process.CompletionTime)
	}
	fmt.Println("-----------------------------------------------------------------")
}
