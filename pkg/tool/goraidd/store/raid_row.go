package store

import "time"

type RaidRow struct {
	ID      uint
	Name    string
	Date    time.Time
	Fights  int
	Players int
}
