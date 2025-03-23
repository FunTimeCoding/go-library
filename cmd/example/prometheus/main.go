package main

import (
	alertmanager "github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/example"
	prometheus "github.com/funtimecoding/go-library/pkg/prometheus/example"
)

func main() {
	if true {
		alertmanager.Notify()
	}

	if false {
		alertmanager.Create()
		prometheus.Query()
		prometheus.Rule()
		alertmanager.Alert()
		alertmanager.Status()
		alertmanager.SetSilence()
		alertmanager.DeleteSilence()
		alertmanager.Silence()
	}
}
