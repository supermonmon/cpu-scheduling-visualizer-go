package processes

import (
	"cpu-scheduling-algorithms/models"
	"fmt"
)

func DeleteProcess(processes *[]models.Process, currentAlgorithm string) {
	if currentAlgorithm == "None" {
		fmt.Println("\n\n\033[1;31mPlease choose an algorithm first!\033[0m")
		return
	}

	if len(*processes) == 0 {
		fmt.Println("\n\n\033[1;31mThere are no processes!\033[0m")
		return
	}

	fmt.Println("Current Processes:")
	DisplayProcesses(*processes)

	var id int
	fmt.Print("\n\nEnter Process ID to delete: ")
	fmt.Scan(&id)

	found := false
	for i, process := range *processes {
		if process.ID == id {
			*processes = append((*processes)[:i], (*processes)[i+1:]...)
			fmt.Println("\033[1;32mProcess with ID:\033[0m", id, "\033[1;32mdeleted successfully!\033[0m")
			found = true
			break
		}
	}

	if !found {
		fmt.Println("\n\n\033[1;31mProcess With ID:\033[0m", id, "\033[1;31mcannot be found!\033[0m")
	}
}
