package processes

import (
	"encoding/csv"
	"fmt"
	"strconv"
)

// Function to process CSV data and extract information
func ProcessCSVData(reader *csv.Reader) ([]string, []int, []int, error) {
	var processID []string
	var arrivalTime, burstTime []int

	records, err := reader.ReadAll()
	if err != nil {
		return nil, nil, nil, fmt.Errorf("error reading CSV file: %v", err)
	}

	for _, record := range records {
		processID = append(processID, record[0]) // No conversion to int
		arrivalTimeValue, err := strconv.Atoi(record[1])
		if err != nil {
			return nil, nil, nil, fmt.Errorf("error converting arrival time (%s) to integer: %v", record[1], err)
		}
		burstTimeValue, err := strconv.Atoi(record[2])
		if err != nil {
			return nil, nil, nil, fmt.Errorf("error converting burst time (%s) to integer: %v", record[2], err)
		}

		arrivalTime = append(arrivalTime, arrivalTimeValue)
		burstTime = append(burstTime, burstTimeValue)
	}

	return processID, arrivalTime, burstTime, nil
}
