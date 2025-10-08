package collect

import "time"

type LastRun struct {
	Command string
	Time    time.Time
}
