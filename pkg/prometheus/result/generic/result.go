package generic

import "time"

type Result struct {
	Type   string
	Metric string
	Value  string
	Time   time.Time
}
