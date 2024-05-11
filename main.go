package main

import (
	"bufio"
	"cpu-scheduling-algorithms/algorithms"
	"cpu-scheduling-algorithms/processes"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Check for command-line argument (file path)
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename.csv>")
		return
	}

	fileName := os.Args[1]

	// Open the CSV file and create a reader
	reader, file, err := processes.OpenAndReadCSV(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close() // Close the file after processing data

	// Process CSV data and extract information
	processID, arrivalTime, burstTime, priorityLevel, timeQuantum, err := processes.ProcessCSVData(reader)
	if err != nil {
		fmt.Println("Error processing CSV data:", err)
		return
	}

	// Display the menu and handle user input
	scanner := bufio.NewScanner(os.Stdin)
	for {
		processes.DisplayMenu()
		if !scanner.Scan() {
			break
		}

		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			fcfsResult := algorithms.FCFS(processID, arrivalTime, burstTime)
			algorithms.Display(fcfsResult)
		case "2":
			sjfResult := algorithms.SJF(processID, arrivalTime, burstTime)
			algorithms.Display(sjfResult)
		case "3":
			srtfResult := algorithms.SRTF(processID, arrivalTime, burstTime)
			algorithms.Display(srtfResult)
		case "4":
			nppResult := algorithms.NPP(processID, arrivalTime, burstTime, priorityLevel)
			algorithms.Display(nppResult)
		case "5":
			rrResult := algorithms.RR(processID, arrivalTime, burstTime, timeQuantum)
			algorithms.Display(rrResult)
		case "6":
			// Execute all algorithms and store results
			srtfResult := algorithms.SRTF(processID, arrivalTime, burstTime)
			nppResult := algorithms.NPP(processID, arrivalTime, burstTime, priorityLevel)
			rrResult := algorithms.RR(processID, arrivalTime, burstTime, timeQuantum)
			fcfsResult := algorithms.FCFS(processID, arrivalTime, burstTime)
			sjfResult := algorithms.SJF(processID, arrivalTime, burstTime)

			algorithms := []struct {
				name   string
				avgWT  float64
				avgTAT float64
			}{
				{"FCFS", fcfsResult.AvgWaitingTime, fcfsResult.AvgTurnAroundTime},
				{"SJF", sjfResult.AvgWaitingTime, sjfResult.AvgTurnAroundTime},
				{"SRTF", srtfResult.AvgWaitingTime, srtfResult.AvgTurnAroundTime},
				{"NPP", nppResult.AvgWaitingTime, nppResult.AvgTurnAroundTime},
				{"RR", rrResult.AvgWaitingTime, rrResult.AvgTurnAroundTime},
			}

			highestAWT := algorithms[0].avgWT
			lowestAWT := algorithms[0].avgWT
			highestATT := algorithms[0].avgTAT
			lowestATT := algorithms[0].avgTAT

			for _, algo := range algorithms {
				if algo.avgWT > highestAWT {
					highestAWT = algo.avgWT
				}
				if algo.avgWT < lowestAWT {
					lowestAWT = algo.avgWT
				}
				if algo.avgTAT > highestATT {
					highestATT = algo.avgTAT
				}
				if algo.avgTAT < lowestATT {
					lowestATT = algo.avgTAT
				}
			}

			fmt.Println("+-----------------------------------------------------------------------------+")

			fmt.Println("\n\033[48;5;24;38;5;15m Scheduling Algorithm Comparison \033[0m")

			// Check if all algorithms have identical performance
			samePerformance := true
			for i := 1; i < len(algorithms); i++ {
				if algorithms[i].avgWT != algorithms[i-1].avgWT || algorithms[i].avgTAT != algorithms[i-1].avgTAT {
					samePerformance = false
					break
				}
			}

			fmt.Printf("\n Average Waiting Time:\n")
			for _, algo := range algorithms {
				color := "\033[0m"
				if samePerformance {
					color = "\033[20;5;33m"
				} else {
					if algo.avgWT == highestAWT {
						color = "\033[20;5;91m" // Red for highest
					} else if algo.avgWT == lowestAWT {
						color = "\033[20;5;32m" // Green for lowest
					}
				}
				fmt.Printf("  * %s: %s%.2f\033[0m\n", algo.name, color, algo.avgWT)
			}

			fmt.Printf("\n Average Turnaround Time:\n")
			for _, algo := range algorithms {
				color := "\033[0m"
				if samePerformance {
					color = "\033[20;5;33m"
				} else {
					if algo.avgTAT == highestATT {
						color = "\033[20;5;91m" // Red for highest
					} else if algo.avgTAT == lowestATT {
						color = "\033[20;5;32m" // Green for lowest
					}
				}

				fmt.Printf("  * %s: %s%.2f\033[0m\n", algo.name, color, algo.avgTAT)
			}

			bestOverall := ""
			for _, algo := range algorithms {
				if algo.avgWT == lowestAWT && algo.avgTAT == lowestATT {
					bestOverall = algo.name
					break
				} else if (algo.avgWT <= lowestAWT+0.1) && (algo.avgTAT <= lowestATT+0.1) {
					// Heuristic: Allow a small tolerance for a balanced algorithm
					if bestOverall == "" {
						bestOverall = algo.name
					}
				}
			}

			worstOverall := ""
			for _, algo := range algorithms {
				if algo.avgWT == highestAWT && algo.avgTAT == highestATT {
					worstOverall = algo.name
					break
				} else if (algo.avgWT >= highestAWT-0.1) && (algo.avgTAT >= highestATT-0.1) {
					// Heuristic: Allow a small tolerance for a balanced algorithm
					if worstOverall == "" {
						worstOverall = algo.name
					}
				}
			}

			if samePerformance {
				fmt.Println("\nAll algorithms are the same in performance")
			} else {
				fmt.Printf("\n Worst Overall Algorithm:  \033[0m")
				if bestOverall != "" {
					fmt.Print("\033[20;5;31m")
					fmt.Print(worstOverall)
					fmt.Print("\033[0m")
				} else {
					fmt.Print("N/A (algorithms have significant trade-offs)")
				}
				fmt.Println("\033[0m")

				fmt.Printf("Best Overall Algorithm:  \033[0m")
				if bestOverall != "" {
					fmt.Print("\033[20;5;32m")
					fmt.Print(bestOverall)
					fmt.Print("\033[0m")
				} else {
					fmt.Print("N/A (algorithms have significant trade-offs)")
				}
				fmt.Println("\033[0m")

			}
			fmt.Print("\n")

		case "Q", "q":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("\n\x1b[41m Invalid choice. Please enter a number from 1 to 5 or 'Q' to quit. \x1b[0m")
			continue
		}

		// Ask the user if they want to try other algorithms or exit
		fmt.Print("Do you want to try other algorithms or exit?")
		fmt.Print(" \x1b[42m Y \x1b[0m")
		fmt.Print(" or ")
		fmt.Print("\x1b[41m N \x1b[0m")
		fmt.Print(" : ")
		if !scanner.Scan() {
			break
		}

		response := strings.TrimSpace(scanner.Text())
		if response != "Y" && response != "y" {
			fmt.Print("\n\x1b[41m Exiting Application... \x1b[0m")
			return
		}
	}
}
