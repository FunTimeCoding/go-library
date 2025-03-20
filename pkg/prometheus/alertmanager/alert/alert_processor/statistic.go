package alert_processor

import "github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/statistic"

func (p *Processor) Statistic() *statistic.Statistic {
	return p.statistic
}
