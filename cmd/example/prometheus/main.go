package main

import (
	alertmanager "github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/example"
	prometheus "github.com/funtimecoding/go-library/pkg/prometheus/example"
)

func main() {
	prometheus.Label()

	if false {
		prometheus.Query()
		alertmanager.Notify()
		alertmanager.Create()
		prometheus.Rule()
		alertmanager.Alert()
		alertmanager.Status()
		alertmanager.SetSilence()
		alertmanager.DeleteSilence()
		alertmanager.Silence()
	}
}
