package alert_processor

import "github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/alert"

func (p *Processor) Process(v []*alert.Alert) []*alert.Alert {
	if p.enricher != nil {
		v = p.enricher.Run(v)
	}

	return v
}
