package processes

import (
	"fmt"
)

// Process represents a single process
type Process struct {
	ID             int
	ArrivalTime    int
	BurstTime      int
	CompletionTime int // New member to track completion time
}

// DisplayMenu displays the main menu
func DisplayMenu(currentAlgorithm string) {
	fmt.Println("\nMenu")
	fmt.Println("1. Choose an algorithm", currentAlgorithm)
	fmt.Println("2. Add a process")
	fmt.Println("3. Delete a process")
	fmt.Println("4. View processes")
	fmt.Println("5. Exit")
	fmt.Print("Enter your choice: ")
}

// AddProcess adds a new process
func AddProcess(processes *[]Process, currentAlgorithm string) {
	if currentAlgorithm == "None" {
		fmt.Println("Please choose an algorithm first!")
		return
	}

	var newProcess Process
	fmt.Print("Enter Arrival Time: ")
	fmt.Scan(&newProcess.ArrivalTime)
	fmt.Print("Enter Burst Time: ")
	fmt.Scan(&newProcess.BurstTime)
	newProcess.ID = len(*processes) + 1

	*processes = append(*processes, newProcess)
	fmt.Println("Process added successfully!")
}

// DeleteProcess deletes a process by ID
func DeleteProcess(processes *[]Process, currentAlgorithm string) {
	if currentAlgorithm == "None" {
		fmt.Println("Please choose an algorithm first!")
		return
	}

	if len(*processes) == 0 {
		fmt.Println("No processes to delete!")
		return
	}

	fmt.Println("Current Processes:")
	DisplayProcesses(*processes)

	var id int
	fmt.Print("Enter Process ID to delete: ")
	fmt.Scan(&id)

	found := false
	for i, process := range *processes {
		if process.ID == id {
			*processes = append((*processes)[:i], (*processes)[i+1:]...)
			fmt.Println("Process with ID", id, "deleted successfully!")
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Process with ID", id, "not found!")
	}
}

// DisplayProcesses displays all processes
func DisplayProcesses(processes []Process) {
	fmt.Println("-----------------------------------------------------------------")
	fmt.Println("| Process ID | Arrival Time | Burst Time | Completion Time |")
	fmt.Println("-----------------------------------------------------------------")
	for _, process := range processes {
		fmt.Printf("|     %2d     |      %8d |      %7d |        %10d |\n", process.ID, process.ArrivalTime, process.BurstTime, process.CompletionTime)
	}
	fmt.Println("-----------------------------------------------------------------")
}
