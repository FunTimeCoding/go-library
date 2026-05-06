package runner

import "time"

const (
	SyncInterval  = 5 * time.Minute
	ApplyInterval = 30 * time.Minute
	Branch        = "main"
	RemoteBranch  = "origin/main"
)
