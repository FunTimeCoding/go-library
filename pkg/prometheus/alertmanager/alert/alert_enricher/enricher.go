package alert_enricher

import "github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/alert_enricher/enrichment"

type Enricher struct {
	enrichments []*enrichment.Enrichment
}
