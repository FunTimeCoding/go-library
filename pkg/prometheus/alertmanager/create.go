package alertmanager

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"
	"github.com/go-openapi/strfmt"
	raw "github.com/prometheus/alertmanager/api/v2/client/alert"
	"github.com/prometheus/alertmanager/api/v2/models"
	"time"
)

func (c *Client) Create(
	name string,
	instance string,
	summary string,
	description string,
	expression string,
) {
	start := time.Now()
	end := start.Add(3 * time.Minute)
	p := raw.NewPostAlertsParams()
	p.Alerts = append(
		p.Alerts,
		&models.PostableAlert{
			Alert: models.Alert{
				Labels: models.LabelSet{
					constant.AlertnameField: name,
					constant.SeverityField:  constant.CriticalSeverity,
					constant.InstanceField:  instance,
				},
				GeneratorURL: strfmt.URI(
					fmt.Sprintf(
						"https://%s/graph?g0.expr=%s",
						c.prometheus.Host(),
						expression,
					),
				),
			},
			Annotations: map[string]string{
				constant.SummaryField:     summary,
				constant.DescriptionField: description,
			},
			StartsAt: strfmt.DateTime(start),
			EndsAt:   strfmt.DateTime(end),
		},
	)
	_, e := c.client.Alert.PostAlerts(p)
	errors.PanicOnError(e)
}
