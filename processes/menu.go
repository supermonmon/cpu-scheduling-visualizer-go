package processes

import (
	"fmt"
)

func DisplayMenu(currentAlgorithm string) {

	fmt.Println("\n=== CPU Scheduling Menu ===")
	fmt.Println("\n1. Choose an algorithm   (Current:", currentAlgorithm, ")")
	fmt.Println("2. Add a process")
	fmt.Println("3. Delete a process")
	fmt.Println("4. View processes")
	fmt.Println("5. Visualize")
	fmt.Println("6. Exit\n")
	fmt.Print("Enter your choice: ")
}
