package runner

import "time"

const (
	SyncInterval      = 5 * time.Minute
	HighstateInterval = 30 * time.Minute
	KeySyncInterval   = 30 * time.Second
	Branch            = "main"
	RemoteBranch      = "origin/main"
	AllMinions        = "*"
)
