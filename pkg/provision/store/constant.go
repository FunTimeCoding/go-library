package store

import "time"

const (
	StatusRunning = "running"
	StatusSuccess = "success"
	StatusError   = "error"
	RetentionDays = 14
	RetentionAge  = RetentionDays * 24 * time.Hour
)
