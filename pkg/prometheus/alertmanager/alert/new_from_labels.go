package alert

import "github.com/prometheus/alertmanager/api/v2/models"

func NewFromLabels(l models.LabelSet) *Alert {
	return &Alert{
		Labels: l,
		Raw:    &models.GettableAlert{Alert: models.Alert{Labels: l}},
	}
}
