package ranges

import (
	prometheus "github.com/prometheus/client_golang/api/prometheus/v1"
	"time"
)

func LastDay(stepMinutes int) prometheus.Range {
	return prometheus.Range{
		Start: time.Now().Add(-24 * time.Hour),
		End:   time.Now(),
		Step:  time.Duration(stepMinutes) * time.Minute,
	}
}
