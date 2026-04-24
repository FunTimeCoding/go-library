package poller

import "github.com/prometheus/client_golang/prometheus"

type metrics struct {
	alertsTotal  prometheus.Counter
	alertsFiring prometheus.Gauge
	recordsTotal prometheus.Gauge
	pollDuration prometheus.Histogram
	lastPollTime prometheus.Gauge
}
