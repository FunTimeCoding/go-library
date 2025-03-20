package alertmanager

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/advanced_parameter"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/alert_processor"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/label_filter"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/name_aliaser"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/name_filter"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/rule_parser"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/statistic"
)

func (c *Client) AlertsAdvanced(
	p *advanced_parameter.Parameter,
	l *name_aliaser.Aliaser,
	n *name_filter.Filter,
	f *label_filter.Filter,
) ([]*alert.Alert, *statistic.Statistic) {
	response, e := c.client.Alert.GetAlerts(nil)
	errors.PanicOnError(e)
	result, s := alert_processor.New(
		p,
		l,
		n,
		f,
		rule_parser.New(c.Rules()),
	).Process(alert.NewSlice(response.GetPayload(), c.host))

	return result, s
}
