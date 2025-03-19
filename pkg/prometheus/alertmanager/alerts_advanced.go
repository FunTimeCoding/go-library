package alertmanager

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/advanced_parameter"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/alert_filter"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/label_filter"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/name_aliaser"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/statistic"
)

func (c *Client) AlertsAdvanced(
	p *advanced_parameter.Parameter,
	l *name_aliaser.Aliaser,
	f *label_filter.Filter,
) ([]*alert.Alert, *statistic.Statistic) {
	response, e := c.client.Alert.GetAlerts(nil)
	errors.PanicOnError(e)
	result := alert.NewSlice(response.GetPayload(), c.host)
	s := statistic.New()

	if false {
		s.Count(result)
	}

	if l != nil {
		l.Run(result)
	}

	if f != nil {
		result = f.Run(result)
	}

	if p != nil {
		result = alert_filter.New(p).Run(result)
	}

	for _, a := range result {
		a.Documentation = c.Documentation(a.Name)
	}

	s.Relevant = len(result)

	return alert.SortByAge(result, false), s
}
