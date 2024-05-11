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

	func TestSRTF(t *testing.T) {
		// Input data
		processID := []string{"1", "2", "3"}
		arrivalTime := []int{0, 1, 2}
		burstTime := []int{3, 5, 2}
	
		expectedGanttChart := []TimeSlice{
			{PID: "1", Start: 0, Stop: 3},
			{PID: "3", Start: 3, Stop: 5},
			{PID: "2", Start: 5, Stop: 10},
		}
	
		// Expected output
		expected := Result{
			Algorithm:         "SRTF",
			ProcessID:         processID,
			ArrivalTime:       arrivalTime,
			BurstTime:         burstTime,
			Priority:          nil,
			TimeQuantum:       0,
			CompletionTime:    []int{3, 10, 5},
			WaitingTime:       []int{0, 4, 1 },
			TurnAroundTime:    []int{3, 9, 3},
			AvgWaitingTime:    1.67,
			AvgTurnAroundTime: 5,
			GanttChart:        expectedGanttChart,
			CPUUtilization:    100,
		}
	
		// Call the function
		result := SRTF(processID, arrivalTime, burstTime)
	
		// Round the AvgWaitingTime to two decimal places for comparison
		result.AvgWaitingTime = math.Round(result.AvgWaitingTime*100) / 100
	
		// Print the actual and expected results
		fmt.Println("SRTF Actual Result:", result)
		fmt.Println("SRTF Expected Result:", expected)
	
		// Compare the result with the expected output
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("SRTF result incorrect, got: %v, want: %v", result, expected)
		}
	}

	func TestNPP(t *testing.T) {
		// Input data
		processID := []string{"1", "2", "3"}
		arrivalTime := []int{0, 1, 2}
		burstTime := []int{3, 5, 2}
		priority := []int{2, 1, 3} 
	
		expectedGanttChart := []TimeSlice{
			{PID: "1", Start: 0, Stop: 3},
			{PID: "2", Start: 3, Stop: 8},
			{PID: "3", Start: 8, Stop: 10},
		}
	
		// Expected output
		expected := Result{
			Algorithm:         "NPP",
			ProcessID:         processID,
			ArrivalTime:       arrivalTime,
			BurstTime:         burstTime,
			Priority:          priority,
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
		result := NPP(processID, arrivalTime, burstTime, priority)
	
		// Round the AvgWaitingTime to two decimal places for comparison
		result.AvgWaitingTime = math.Round(result.AvgWaitingTime*100) / 100
	
		// Print the actual and expected results
		fmt.Println("NPP Actual Result:", result)
		fmt.Println("NPP Expected Result:", expected)
	
		// Compare the result with the expected output
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("NPP result incorrect, got: %v, want: %v", result, expected)
		}
	}

func TestRR(t *testing.T) {
	// Input data
	processID := []string{"1", "2", "3", "4"}
	arrivalTime := []int{0, 1, 2, 3}
	burstTime := []int{5, 4, 3, 2}
	timeQuantum := 2

	expectedGanttChart := []TimeSlice{
		{PID: "1", Start: 0, Stop: 2},
		{PID: "2", Start: 2, Stop: 4},
		{PID: "3", Start: 4, Stop: 6},
		{PID: "4", Start: 6, Stop: 8},
		{PID: "1", Start: 8, Stop: 10},
		{PID: "2", Start: 10, Stop: 12},
		{PID: "3", Start: 12, Stop: 13},
		{PID: "1", Start: 13, Stop: 14},
	}

	// Expected output
	expected := Result{
		Algorithm:         "RR",
		ProcessID:         processID,
		ArrivalTime:       arrivalTime,
		BurstTime:         burstTime,
		Priority:          nil,
		TimeQuantum:       timeQuantum,
		CompletionTime:    []int{14, 12, 13, 8},
		WaitingTime:       []int{9, 7, 8, 3},
		TurnAroundTime:    []int{14, 11, 11, 5},
		AvgWaitingTime:    6.75,
		AvgTurnAroundTime: 10.25,
		GanttChart:        expectedGanttChart,
		CPUUtilization:    100,
	}

	result := RR(processID, arrivalTime, burstTime, timeQuantum)

	// Round the AvgWaitingTime to two decimal places for comparison
	result.AvgWaitingTime = math.Round(result.AvgWaitingTime*100) / 100

	// Print the actual and expected results
	fmt.Println("Round Robin Actual Result:", result)
	fmt.Println("Round Robin Expected Result:", expected)

	// Compare the result with the expected output
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Round Robin result incorrect, got: %v, want: %v", result, expected)
	}
}

	
