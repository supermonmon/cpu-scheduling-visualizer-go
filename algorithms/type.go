package algorithms

type TimeSlice struct {
	PID   string
	Start int
	Stop  int
}
type FCFSResult struct {
	ProcessID         []string
	ArrivalTime       []int
	BurstTime         []int
	CompletionTime    []int
	WaitingTime       []int
	TurnAroundTime    []int
	AvgWaitingTime    float64
	AvgTurnAroundTime float64
	GanttChart        []TimeSlice
}

type SJFResult struct {
	ProcessID         []string
	ArrivalTime       []int
	BurstTime         []int
	CompletionTime    []int
	WaitingTime       []int
	TurnAroundTime    []int
	AvgWaitingTime    float64
	AvgTurnAroundTime float64
	GanttChart        []TimeSlice
}

type SRTFResult struct {
	ProcessID         []string
	ArrivalTime       []int
	BurstTime         []int
	CompletionTime    []int
	WaitingTime       []int
	TurnAroundTime    []int
	GanttChart        []TimeSlice
	AvgWaitingTime    float64
	AvgTurnAroundTime float64
}
type NPPResult struct {
	ProcessID         []string
	ArrivalTime       []int
	BurstTime         []int
	Priority          []int
	CompletionTime    []int
	WaitingTime       []int
	TurnAroundTime    []int
	GanttChart        []TimeSlice
	AvgWaitingTime    float64
	AvgTurnAroundTime float64
}

type RRResult struct {
	ProcessID         []string
	ArrivalTime       []int
	BurstTime         []int
	TimeQuantum       int
	CompletionTime    []int
	WaitingTime       []int
	TurnAroundTime    []int
	GanttChart        []TimeSlice
	AvgWaitingTime    float64
	AvgTurnAroundTime float64
}

type SchedulingResult struct {
	AvgTurnAroundTime float64
	AvgWaitingTime    float64
}
