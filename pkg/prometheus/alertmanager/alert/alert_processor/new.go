package alert_processor

import (
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/advanced_parameter"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/label_filter"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/name_aliaser"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/name_filter"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/rule_parser"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/statistic"
)

func New(
	p *advanced_parameter.Parameter,
	a *name_aliaser.Aliaser,
	n *name_filter.Filter,
	l *label_filter.Filter,
	r *rule_parser.Parser,
) *Processor {
	return &Processor{
		parameter:   p,
		aliaser:     a,
		nameFilter:  n,
		labelFilter: l,
		rule:        r,
		statistic:   statistic.New(),
	}
}
