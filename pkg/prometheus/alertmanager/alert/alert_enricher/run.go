package alert_enricher

import "github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert"

func (e *Enricher) Run(v []*alert.Alert) []*alert.Alert {
	for _, l := range v {
		e.enrich(l)
	}

	return v
}
