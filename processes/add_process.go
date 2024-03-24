package processes

import (
	"cpu-scheduling-algorithms/models"
	"fmt"
)

func AddProcess(processes *[]models.Process, currentAlgorithm string) {

	if currentAlgorithm == "None" {
		fmt.Println("\n\n\033[1;31mPlease choose an algorithm first!\033[0m")
		return
	}

	var arrivalTimes, burstTimes []int
	var numProcesses int

	fmt.Print("Enter number of processes: ")
	fmt.Scan(&numProcesses)

	arrivalTimes = make([]int, numProcesses)
	burstTimes = make([]int, numProcesses)

	fmt.Println("Enter arrival times separated by spaces:")
	for i := 0; i < numProcesses; i++ {
		fmt.Scan(&arrivalTimes[i])
	}

	fmt.Println("Enter burst times separated by spaces:")
	for i := 0; i < numProcesses; i++ {
		fmt.Scan(&burstTimes[i])
	}

	for i := 0; i < numProcesses; i++ {
		newProcess := models.Process{
			ID:          len(*processes) + 1,
			ArrivalTime: arrivalTimes[i],
			BurstTime:   burstTimes[i],
		}
		*processes = append(*processes, newProcess)
	}

	fmt.Println("\n\n\033[1;32mProcesses added successfully!\033[0m")
}
