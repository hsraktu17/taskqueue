package model

type JobStatus string

const (
	StatusPending   JobStatus = "pending"
	StatusRunning   JobStatus = "running"
	StatusSuccess   JobStatus = "success"
	StatusFailed    JobStatus = "failed"
	StatusCancelled JobStatus = "cancelled"
	StatusScheduled JobStatus = "scheduled"
)
