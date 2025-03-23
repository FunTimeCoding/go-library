package alertmanager

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/advanced_option"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/alert_enricher"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/alert_processor"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/field_changer"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/label_filter"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/name_filter"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/rule_parser"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/statistic"
)

func (c *Client) AlertsAdvanced(
	p *advanced_option.Alert,
	n *alert_enricher.Enricher,
	h *field_changer.Changer,
	a *name_filter.Filter,
	f *label_filter.Filter,
) ([]*alert.Alert, *statistic.Statistic) {
	response, e := c.client.Alert.GetAlerts(nil)
	errors.PanicOnError(e)
	result, s := alert_processor.New(
		p,
		n,
		h,
		a,
		f,
		rule_parser.New(c.Rules()),
	).Process(alert.NewSlice(response.GetPayload(), c.host))

	return result, s
}
