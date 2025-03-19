package ranges

import (
	prometheus "github.com/prometheus/client_golang/api/prometheus/v1"
	"time"
)

func LastHour(stepMinutes int) prometheus.Range {
	return prometheus.Range{
		Start: time.Now().Add(-time.Hour),
		End:   time.Now(),
		Step:  time.Duration(stepMinutes) * time.Minute,
	}
}
