package store

import "time"

const (
	StatusPending = "pending"
	StatusRunning = "running"
	StatusSuccess = "success"
	StatusError   = "error"
	TriggerTimer  = "timer"
	TriggerManual = "manual"
	RetentionDays = 14
	RetentionAge  = RetentionDays * 24 * time.Hour
)
