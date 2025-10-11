package main

import (
	grafana "github.com/funtimecoding/go-library/pkg/grafana/example"
	alertmanager "github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/example"
	prometheus "github.com/funtimecoding/go-library/pkg/prometheus/example"
	loki "github.com/funtimecoding/go-library/pkg/prometheus/loki/example"
)

func main() {
	prometheus.Target()
	prometheus.Series()

	if false {
		grafana.Read()

		loki.Load()
		loki.Write()

		prometheus.Status()
		prometheus.Metric()
		prometheus.Label()
		prometheus.LabelName()
		prometheus.Query()
		prometheus.Rule()
		prometheus.Meta()

		alertmanager.Notify()
		alertmanager.Create()
		alertmanager.Alert()
		alertmanager.Status()
		alertmanager.SetSilence()
		alertmanager.DeleteSilence()
		alertmanager.Silence()
	}
}
