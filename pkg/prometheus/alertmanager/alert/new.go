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

	result := &Alert{
		MonitorIdentifier: fmt.Sprintf(
			"%s-%s",
			monitorConstant.AlertPrefix,
			*v.Fingerprint,
		),
		State:     state,
		Labels:    v.Labels,
		Remaining: remaining,
		Raw:       v,
		Start:     openapi.ConvertTime(v.StartsAt),
	}
	extractKey(&remaining, constant.AlertnameField, &result.Name)
	extractKey(&remaining, constant.SeverityField, &result.Severity)
	extractKey(&remaining, constant.SummaryField, &result.Summary)
	extractKey(&remaining, constant.MessageField, &result.Message)
	extractKey(&remaining, constant.PrometheusField, &result.Prometheus)
	result.Link = fmt.Sprintf(
		"https://%s/#/alerts?filter=%s",
		host,
		url.QueryEscape(
			fmt.Sprintf(
				"{%s=\"%s\"}",
				constant.AlertnameField,
				result.Name,
			),
		),
	)

	return result
}
