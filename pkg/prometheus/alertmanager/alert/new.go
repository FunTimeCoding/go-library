package alert

import (
	"fmt"
	monitor "github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/funtimecoding/go-library/pkg/openapi"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"
	"github.com/prometheus/alertmanager/api/v2/models"
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
		if receiver.Name == nil {
			continue
		}

		n := *receiver.Name

		switch n {
		case "":
			continue
		case "null":
			continue
		}

		receivers = append(receivers, n)
	}

	result := &Alert{
		MonitorIdentifier: fmt.Sprintf(
			"%s-%s",
			monitor.AlertPrefix,
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
	result.Link = result.buildLink(host)

	return result
}
