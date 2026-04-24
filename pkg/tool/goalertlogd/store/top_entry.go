package store

import "time"

type topEntry struct {
	name            string
	count           int
	totalDuration   time.Duration
	resolvedCount   int
	currentlyFiring int
	severity        string
}
