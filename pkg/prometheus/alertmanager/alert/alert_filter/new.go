package alert_filter

import "github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/advanced_parameter"

func New(p *advanced_parameter.Parameter) *Filter {
	return &Filter{parameter: p}
}
