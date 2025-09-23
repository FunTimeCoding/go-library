package main

import (
	alertmanager "github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/example"
	prometheus "github.com/funtimecoding/go-library/pkg/prometheus/example"
)

func main() {
	prometheus.Status()
	prometheus.Metric()
	prometheus.Label()
	prometheus.LabelName()
	prometheus.Query()
	prometheus.Rule()
	prometheus.Meta()

	if false {
		alertmanager.Notify()
		alertmanager.Create()
		alertmanager.Alert()
		alertmanager.Status()
		alertmanager.SetSilence()
		alertmanager.DeleteSilence()
		alertmanager.Silence()
	}
}
