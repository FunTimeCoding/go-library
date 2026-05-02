package worker

import "github.com/prometheus/client_golang/prometheus"

func newMetrics(r *prometheus.Registry) *metrics {
	m := &metrics{
		alertsTotal: prometheus.NewCounter(
			prometheus.CounterOpts{
				Name: "alertlog_alerts_total",
				Help: "Total number of alerts recorded.",
			},
		),
		alertsFiring: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "alertlog_alerts_firing",
				Help: "Number of currently firing alerts.",
			},
		),
		recordsTotal: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "alertlog_records_total",
				Help: "Total number of stored alert records.",
			},
		),
		pollDuration: prometheus.NewHistogram(
			prometheus.HistogramOpts{
				Name:    "alertlog_poll_duration_seconds",
				Help:    "Duration of a single poll cycle.",
				Buckets: prometheus.DefBuckets,
			},
		),
		lastPollTime: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "alertlog_last_poll_timestamp_seconds",
				Help: "Unix timestamp of the last successful poll.",
			},
		),
	}
	r.MustRegister(
		m.alertsTotal,
		m.alertsFiring,
		m.recordsTotal,
		m.pollDuration,
		m.lastPollTime,
	)

	return m
}
