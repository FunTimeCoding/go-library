package alert_processor

import (
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/alert_filter"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/statistic"
)

func (p *Processor) Process(v []*alert.Alert) ([]*alert.Alert, *statistic.Statistic) {
	p.statistic.CountBeforeProcessing(v)

	if p.enricher != nil {
		v = p.enricher.Run(v)
	}

	if p.changer != nil {
		v = p.changer.Run(v)
	}

	if p.nameFilter != nil {
		v = p.nameFilter.Run(v)
	}

	if p.labelFilter != nil {
		v = p.labelFilter.Run(v)
	}

	if p != nil {
		v = alert_filter.New(p.parameter).Run(v)
	}

	for _, a := range v {
		a.Documentation = p.rule.Documentation(a.Name)
	}

	p.statistic.CountAfterProcessing(v)

	return alert.SortByAge(v, false), p.statistic
}
