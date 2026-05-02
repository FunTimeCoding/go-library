package ranges

import (
	"github.com/prometheus/client_golang/api/prometheus/v1"
	"time"
)

func LastDay(stepMinutes int) v1.Range {
	return v1.Range{
		Start: time.Now().Add(-24 * time.Hour),
		End:   time.Now(),
		Step:  time.Duration(stepMinutes) * time.Minute,
	}
}
