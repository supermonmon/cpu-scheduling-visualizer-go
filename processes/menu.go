package processes

import (
	"fmt"
)

func DisplayMenu() {
	fmt.Println("+-----------------------------------------------------------------------------+")

	fmt.Println("\n\033[48;5;24;38;5;15m⚙️  CPU Scheduling Algorithms \033[0m")
	fmt.Print("\n")

	fmt.Println(`Please choose an algorithm:`)
	fmt.Println("\n\x1b[40m\x1b[34m 1. First Come First Serve \x1b[0m")
	fmt.Print("\n")
	fmt.Println("\x1b[40m\x1b[34m 2. Shortest Job First \x1b[0m")
	fmt.Print("\n")
	fmt.Println("\x1b[40m\x1b[34m 3. Shortest Remaining Time First \x1b[0m")
	fmt.Print("\n")
	fmt.Println("\x1b[40m\x1b[34m 4. Priority \x1b[0m")
	fmt.Print("\n")
	fmt.Println("\x1b[40m\x1b[34m 5. Round Robin \x1b[0m")
	fmt.Print("\n")

	fmt.Println("\n\x1b[47m Press 'Q' to exit \x1b[0m")
	fmt.Print("\n")

	fmt.Print(`Enter your choice: `)
}
