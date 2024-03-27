package algorithms

// SJFGantt generates a Gantt chart specifically for SJF scheduling
func SJFGantt(processID []string, completionTime []int) []TimeSlice {
	gantt := make([]TimeSlice, len(processID))
	var prevCompletionTime int = 0
	for i := range processID {
		gantt[i] = TimeSlice{PID: processID[i], Start: prevCompletionTime}
		prevCompletionTime = completionTime[i] // Update start time for next process based on previous completion
		gantt[i].Stop = prevCompletionTime
	}
	return gantt
}
