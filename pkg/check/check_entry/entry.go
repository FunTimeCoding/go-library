package check_entry

import "time"

type Entry struct {
	Time  time.Time
	Level string
	Text  string
}
