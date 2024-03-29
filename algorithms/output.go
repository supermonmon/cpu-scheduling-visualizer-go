package algorithms

import (
	"fmt"
	"io"
	"strings"
)

// Function to print the Gantt chart with improved formatting and color-coding
func outputGantt(w io.Writer, gantt []TimeSlice) {
	var colorMap = map[string]string{
		"P1": "\033[44m", // Blue
		"P2": "\033[42m", // Green
		"P3": "\033[43m", // Yellow
		"P4": "\033[41m", // Red
	}

	_, _ = fmt.Fprintln(w, "Gantt schedule")
	_, _ = fmt.Fprint(w, " ")

	maxLength := 0
	for _, ts := range gantt {
		pidStr := fmt.Sprint(ts.PID)
		if len(pidStr) > maxLength {
			maxLength = len(pidStr)
		}
	}
	format := strings.Repeat(" ", maxLength+2) + "%v  "
	for _, ts := range gantt {
		color := colorMap[ts.PID]
		if color == "" { // Wrap around colors for more processes
			color = colorMap["P"+fmt.Sprintf("%d", (int(ts.PID[1])-48)%4+1)] // Use modulo 4 for Blue-Green-Yellow-Red cycle
		}
		_, _ = fmt.Fprintf(w, color+format+"\033[0m", ts.PID) // Apply color and reset
	}
	_, _ = fmt.Fprintln(w)

	_, _ = fmt.Fprint(w, "0")
	for _, ts := range gantt {
		_, _ = fmt.Fprintf(w, "\t%d", ts.Stop)
	}
	_, _ = fmt.Fprintln(w)
	_, _ = fmt.Fprintf(w, "\n\n")
}
