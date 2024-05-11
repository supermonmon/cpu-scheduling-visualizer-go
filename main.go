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
	defer file.Close()

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
		case "Q", "q":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("\n\x1b[41m Invalid choice. Please enter a number from 1 to 5 or 'Q' to quit. \x1b[0m")
			continue
		}

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
