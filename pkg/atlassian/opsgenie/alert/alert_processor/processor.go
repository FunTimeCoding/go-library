package alert_processor

import "github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/alert/alert_enricher"

type Processor struct {
	enricher *alert_enricher.Enricher
}
