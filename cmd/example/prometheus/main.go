package main

import (
	alertmanager "github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/example"
	prometheus "github.com/funtimecoding/go-library/pkg/prometheus/example"
)

func main() {
	if false {
		prometheus.Rule()
		alertmanager.Alert()
		alertmanager.Status()
		alertmanager.DeleteSilence()
	}

	if true {
		if false {
			alertmanager.CreateSilence()
		}

		alertmanager.Silence()
	}
}
