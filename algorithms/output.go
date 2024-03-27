package algorithms

import (
	"fmt"
	"io"
	"strings"
)

// Function to print the Gantt chart with improved formatting
func outputGantt(w io.Writer, gantt []TimeSlice) {
	_, _ = fmt.Fprintln(w, "Gantt schedule")
	_, _ = fmt.Fprint(w, "|")
	maxLength := 0
	for _, ts := range gantt {
		pidStr := fmt.Sprint(ts.PID)
		if len(pidStr) > maxLength {
			maxLength = len(pidStr)
		}
	}
	format := strings.Repeat(" ", maxLength+2) + "%v |"
	for _, ts := range gantt {
		_, _ = fmt.Fprintf(w, format, ts.PID)
	}
	_, _ = fmt.Fprintln(w)

	_, _ = fmt.Fprint(w, "0")
	for _, ts := range gantt {
		_, _ = fmt.Fprintf(w, "\t%d", ts.Stop)
	}
	_, _ = fmt.Fprintln(w)
	_, _ = fmt.Fprintf(w, "\n\n")
}
