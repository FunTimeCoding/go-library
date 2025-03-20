package alert_processor

import (
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/advanced_parameter"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/alert_enricher"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/field_changer"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/label_filter"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/name_filter"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/rule_parser"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/statistic"
)

type Processor struct {
	parameter   *advanced_parameter.Parameter
	enricher    *alert_enricher.Enricher
	changer     *field_changer.Changer
	nameFilter  *name_filter.Filter
	labelFilter *label_filter.Filter
	rule        *rule_parser.Parser
	statistic   *statistic.Statistic
}
