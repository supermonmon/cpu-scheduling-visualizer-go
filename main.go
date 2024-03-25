package main

import (
	"cpu-scheduling-algorithms/processes"
	"fmt"
)

func main() {

	for {
		processes.DisplayMenu()

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			processes.RunFCFS()
		case 2:

		case 3:

		case 4:

		case 5:

		case 6:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("\nError: " + "\x1b[31mInvalid. Please try again!\x1b[0m")

		}
	}
}
