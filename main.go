package main

import (
	"cpu-scheduling-algorithms/models"
	"cpu-scheduling-algorithms/processes"
	"fmt"
)

func main() {
	var processesList []models.Process
	currentAlgorithm := "None"

	for {
		DisplayMenu()

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Choose an Algorithm")
			fmt.Println("1. First-Come, First-Served (FCFS)")
			fmt.Println("2. Shortest Job First (SJF)")
			fmt.Print("Enter your choice: ")
			fmt.Scan(&choice)

			switch choice {
			case 1:
				currentAlgorithm = "FCFS"
				processesList = nil
			case 2:
				currentAlgorithm = "SJF"
				processesList = nil
			default:
				fmt.Println("Invalid choice!")
			}
		case 2:
			processes.AddProcess(&processesList, currentAlgorithm)
		case 3:
			processes.DeleteProcess(&processesList, currentAlgorithm)
		case 4:
			processes.DisplayProcesses(processesList)
		case 5:
			processes.VisualizeOption(processesList, currentAlgorithm)
		case 6:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice!")
		}
	}
}
