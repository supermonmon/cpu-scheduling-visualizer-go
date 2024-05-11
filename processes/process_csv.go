package processes

import (
	"encoding/csv"
	"fmt"
	"strconv"
)

func ProcessCSVData(reader *csv.Reader) ([]string, []int, []int, []int, int, error) {
	var processID []string
	var arrivalTime, burstTime, priorityLevel []int
	var timeQuantumValue int

	records, err := reader.ReadAll()
	if err != nil {
		return nil, nil, nil, nil, 0, fmt.Errorf("error reading CSV file: %v", err)
	}

	// Extract time quantum
	timeQuantumValue, err = strconv.Atoi(records[0][len(records[0])-1])
	if err != nil {
		return nil, nil, nil, nil, 0, fmt.Errorf("error converting time quantum (%s) to integer: %v", records[0][len(records[0])-1], err)
	}

	// Loop through records, starting from index 0 to include the first process
	for _, record := range records {
		processID = append(processID, record[0])
		arrivalTimeValue, err := strconv.Atoi(record[1])
		if err != nil {
			return nil, nil, nil, nil, 0, fmt.Errorf("error converting arrival time (%s) to integer: %v", record[1], err)
		}
		burstTimeValue, err := strconv.Atoi(record[2])
		if err != nil {
			return nil, nil, nil, nil, 0, fmt.Errorf("error converting burst time (%s) to integer: %v", record[2], err)
		}
		priorityLevelValue, err := strconv.Atoi(record[3])
		if err != nil {
			return nil, nil, nil, nil, 0, fmt.Errorf("error converting priority level (%s) to integer: %v", record[3], err)
		}

		arrivalTime = append(arrivalTime, arrivalTimeValue)
		burstTime = append(burstTime, burstTimeValue)
		priorityLevel = append(priorityLevel, priorityLevelValue)
	}

	return processID, arrivalTime, burstTime, priorityLevel, timeQuantumValue, nil
}
