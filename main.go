// main.go

package main

import (
	"cpu-scheduling-algorithms/processes"
	"fmt"
)

func main() {
	var chosenAlgorithm string
	var allProcesses []processes.Process

	for {
		processes.DisplayMenu(chosenAlgorithm)
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter the name of the algorithm: ")
			fmt.Scanln(&chosenAlgorithm)
		case 2:
			processes.AddProcess(&allProcesses, chosenAlgorithm)
		case 3:
			processes.DeleteProcess(&allProcesses, chosenAlgorithm)
		case 4:
			processes.DisplayProcesses(allProcesses)
		case 5:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please choose again.")
		}
	}
}
