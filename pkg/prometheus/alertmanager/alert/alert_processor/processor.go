package alert_processor

import (
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/advanced_parameter"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/label_filter"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/name_aliaser"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/name_filter"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/rule_parser"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/statistic"
)

type Processor struct {
	parameter   *advanced_parameter.Parameter
	aliaser     *name_aliaser.Aliaser
	nameFilter  *name_filter.Filter
	labelFilter *label_filter.Filter
	rule        *rule_parser.Parser
	statistic   *statistic.Statistic
}
