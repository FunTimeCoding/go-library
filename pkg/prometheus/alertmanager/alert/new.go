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
		if !slices.Contains(constant.States, state) {
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

	ExtractKey(&remaining, constant.NameField, &result.Name)
	ExtractKey(&remaining, constant.SeverityField, &result.Severity)
	ExtractKey(&remaining, constant.SummaryField, &result.Summary)
	ExtractKey(&remaining, constant.MessageField, &result.Message)
	ExtractKey(&remaining, constant.PrometheusField, &result.Prometheus)

	return result
}

func ExtractKey(
	remaining *models.LabelSet,
	k string,
	to *string,
) {
	if v, ok := (*remaining)[k]; ok {
		*to = v
		delete(*remaining, k)
	} else {
		*to = constant.None
	}
}
