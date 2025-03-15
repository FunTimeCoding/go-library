package alert

import (
	"github.com/prometheus/alertmanager/api/v2/models"
	"time"
)

type Alert struct {
	Name       string
	State      string
	Severity   string
	Summary    string
	Message    string
	Prometheus string
	Start      *time.Time
	Labels     models.LabelSet

	RemainingLabels models.LabelSet
	Documentation   string

	Raw *models.GettableAlert
}
