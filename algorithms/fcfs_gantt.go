package algorithms

// Function to create Gantt chart data structure
func FCFSGantt(processID []string, burstTime []int) []TimeSlice { // Updated processID type
	gantt := make([]TimeSlice, len(processID))
	var current int = 0
	for i := range processID {
		gantt[i] = TimeSlice{PID: processID[i], Start: current} // Use processID as string
		current += burstTime[i]
		gantt[i].Stop = current
	}
	return gantt
}
