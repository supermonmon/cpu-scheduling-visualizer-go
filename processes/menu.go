package processes

import (
	"fmt"
)

func DisplayMenu() {
	fmt.Println("+-----------------------------------------------------------------------------+")

	fmt.Println("\n\033[48;5;24;38;5;15m⚙️  CPU Scheduling Algorithms \033[0m\n")

	fmt.Print(`Please choose an algorithm:

1. First come First Serve (FCFS)
2. Shortest Job First (SJF)
3. Shortest Remaining Time First (SRTF)
4. Priority Queue (PQ)
5. Round Robin (RR)
`)

	fmt.Println("\n\x1b[40m Press 'Q' to exit \x1b[0m\n")

	fmt.Print(`Enter your choice: `)
}
