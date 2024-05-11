package algorithms

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

	func TestFCFS(t *testing.T) {
		// Input data
		processID := []string{"1", "2", "3"}
		arrivalTime := []int{0, 1, 2}
		burstTime := []int{3, 5, 2}

		expectedGanttChart := []TimeSlice{
			{PID: "1", Start: 0, Stop: 3},
			{PID: "2", Start: 3, Stop: 8},
			{PID: "3", Start: 8, Stop: 10},
		}

		// Expected output
		expected := Result{
			Algorithm:         "FCFS",
			ProcessID:         processID,
			ArrivalTime:       arrivalTime,
			BurstTime:         burstTime,
			Priority:          nil,
			TimeQuantum:       0,
			CompletionTime:    []int{3, 8, 10},
			WaitingTime:       []int{0, 2, 6},
			TurnAroundTime:    []int{3, 7, 8},
			AvgWaitingTime:    2.67,
			AvgTurnAroundTime: 6,
			GanttChart:        expectedGanttChart,
			CPUUtilization:    100,
		}

		// Call the function
		result := FCFS(processID, arrivalTime, burstTime)

		// Round the AvgWaitingTime to two decimal places for comparison
		result.AvgWaitingTime = math.Round(result.AvgWaitingTime*100) / 100

		// Print the actual and expected results
		fmt.Println("Actual Result:", result)
		fmt.Println("Expected Result:", expected)

		// Compare the result with the expected output
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("FCFS result incorrect, got: %v, want: %v", result, expected)
		}
	}

	func TestSJF(t *testing.T) {
		// Input data
		processID := []string{"1", "2", "3"}
		arrivalTime := []int{0, 1, 2}
		burstTime := []int{3, 5, 2}
	
		expectedGanttChart := []TimeSlice{
			{PID: "1", Start: 0, Stop: 3},
			{PID: "2", Start: 3, Stop: 5},
			{PID: "3", Start: 5, Stop: 10},
		}
	
		// Expected output
		expected := Result{
			Algorithm:         "SJF",
			ProcessID:         processID,
			ArrivalTime:       arrivalTime,
			BurstTime:         burstTime,
			Priority:          nil,
			TimeQuantum:       0,
			CompletionTime:    []int{3, 5, 10},
			WaitingTime:       []int{0, 1, 4},
			TurnAroundTime:    []int{3, 4, 8},
			AvgWaitingTime:    1.67,
			AvgTurnAroundTime: 5,
			GanttChart:        expectedGanttChart,
			CPUUtilization:    100,
		}
	
		// Call the function
		result := SJF(processID, arrivalTime, burstTime)
	
		// Round the AvgWaitingTime to two decimal places for comparison
		result.AvgWaitingTime = math.Round(result.AvgWaitingTime*100) / 100
	
		// Print the actual and expected results
		fmt.Println("SJF Actual Result:", result)
		fmt.Println("SJF Expected Result:", expected)
	
		// Compare the result with the expected output
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("SJF result incorrect, got: %v, want: %v", result, expected)
		}
	}
	
