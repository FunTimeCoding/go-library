package alert

import (
	"fmt"
	monitorConstant "github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/funtimecoding/go-library/pkg/openapi"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"
	"github.com/prometheus/alertmanager/api/v2/models"
	"net/url"
	"slices"
)

func New(
	v *models.GettableAlert,
	host string,
) *Alert {
	remaining := v.Labels
	state := *v.Status.State

	if state == "" {
		state = constant.None
	} else {
		if !slices.Contains(constant.AlertStates, state) {
			fmt.Printf("Unexpected state: %s\n", state)
		}
	}

	var receivers []string

	for _, receiver := range v.Receivers {
		receivers = append(receivers, *receiver.Name)
	}

	result := &Alert{
		MonitorIdentifier: fmt.Sprintf(
			"%s-%s",
			monitorConstant.AlertPrefix,
			*v.Fingerprint,
		),
		Fingerprint: *v.Fingerprint,
		State:       state,
		Labels:      v.Labels,
		Receivers:   receivers,
		Remaining:   remaining,
		Raw:         v,
		Start:       openapi.ConvertTime(v.StartsAt),
	}
	extractKey(&remaining, constant.AlertnameLabel, &result.Name)
	extractKey(&remaining, constant.SeverityLabel, &result.Severity)
	extractKey(&remaining, constant.SummaryLabel, &result.Summary)
	extractKey(&remaining, constant.MessageLabel, &result.Message)
	extractKey(&remaining, constant.PrometheusLabel, &result.Prometheus)
	result.Link = fmt.Sprintf(
		"https://%s/#/alerts?filter=%s",
		host,
		url.QueryEscape(
			fmt.Sprintf(
				"{%s=\"%s\"}",
				constant.AlertnameLabel,
				result.Name,
			),
		),
	)

	return result
}
