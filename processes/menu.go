package processes

import (
	"fmt"
)

func DisplayMenu() {
	fmt.Print(`
=== CPU Scheduling Menu ===

Please choose an algorithm:

1. First come First Serve (FCFS)
2. Shortest Job First (SJF)
3. Shortest Remaining Time First (SRTF)
4. Priority Queue (PQ)
5. Round Robin (RR)

Press 'Q' to exit.

Enter your choice: `)
}
