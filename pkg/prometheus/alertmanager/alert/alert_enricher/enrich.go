package alert_enricher

import "github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert"

func (e *Enricher) enrich(a *alert.Alert) {
	for _, n := range e.enrichments {
		if n.Name == a.Name {
			a.Entity = n.Entity
			a.Category = n.Category
			a.Tag(n.Tags...)
		}
	}
}
