package store

import "time"

type TopRecord struct {
	Name            string
	Count           int
	AverageDuration time.Duration
	CurrentlyFiring int
	Severity        string
}
