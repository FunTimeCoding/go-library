package log

import "time"

type SeenDaysMember struct {
	Name   string
	Fights int
	Days   []time.Time
}
