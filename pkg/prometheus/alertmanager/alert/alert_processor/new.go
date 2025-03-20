package alert_processor

import (
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/advanced_parameter"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/field_changer"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/label_filter"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/name_filter"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/rule_parser"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/statistic"
)

func New(
	p *advanced_parameter.Parameter,
	h *field_changer.Changer,
	n *name_filter.Filter,
	l *label_filter.Filter,
	r *rule_parser.Parser,
) *Processor {
	return &Processor{
		parameter:   p,
		changer:     h,
		nameFilter:  n,
		labelFilter: l,
		rule:        r,
		statistic:   statistic.New(),
	}
}
