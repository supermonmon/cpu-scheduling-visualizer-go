package algorithms

type TimeSlice struct {
	PID   string
	Start int
	Stop  int
}
type Result struct {
	Algorithm         string
	ProcessID         []string
	ArrivalTime       []int
	BurstTime         []int
	Priority          []int
	TimeQuantum       int
	CompletionTime    []int
	WaitingTime       []int
	TurnAroundTime    []int
	AvgWaitingTime    float64
	AvgTurnAroundTime float64
	GanttChart        []TimeSlice
	CPUUtilization    float64
}

type FCFSData struct {
	pid         string
	arrivalTime int
	burstTime   int
}

type SJFData struct {
	pid         string
	arrivalTime int
	burstTime   int
}
type SRTFData struct {
	pid         string
	arrivalTime int
	burstTime   int
	remaining   int
	completed   bool
}

type NPPData struct {
	pid         string
	arrivalTime int
	burstTime   int
	priority    int
	completed   bool
}