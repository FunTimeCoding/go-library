package alert

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/openapi"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"
	"github.com/prometheus/alertmanager/api/v2/models"
	"slices"
)

func New(v *models.GettableAlert) *Alert {
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
		State:           state,
		Labels:          v.Labels,
		RemainingLabels: remaining,
		Raw:             v,
		Start:           openapi.ConvertTime(v.StartsAt),
	}
	extractKey(&remaining, constant.NameField, &result.Name)
	extractKey(&remaining, constant.SeverityField, &result.Severity)
	extractKey(&remaining, constant.SummaryField, &result.Summary)
	extractKey(&remaining, constant.MessageField, &result.Message)
	extractKey(&remaining, constant.PrometheusField, &result.Prometheus)

	return result
}
