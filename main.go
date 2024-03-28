package main

import (
	"cpu-scheduling-algorithms/algorithms"
	"cpu-scheduling-algorithms/processes"
	"fmt"
	"os"
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

	fmt.Println("Process ID:", processID)
	fmt.Println("Arrival Time", arrivalTime)
	fmt.Println("Burst Time:", burstTime)
	fmt.Println("Priority Level:", priorityLevel)
	fmt.Println("Time Quantum:", timeQuantum)

	// Call FCFS scheduling function
	//algorithms.SRTF(processID, arrivalTime, burstTime)
	//algorithms.NPP(processID, arrivalTime, burstTime, priorityLevel)
	algorithms.RR(processID, arrivalTime, burstTime, timeQuantum)
}
