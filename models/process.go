package models

type Process struct {
	ID             int
	ArrivalTime    int
	BurstTime      int
	CompletionTime int
	TurnaroundTime int
	WaitingTime    int
}
