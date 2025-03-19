package alert

import (
	"github.com/prometheus/alertmanager/api/v2/models"
	"time"
)

type Alert struct {
	MonitorIdentifier string
	Name              string
	State             string
	Severity          string
	Summary           string
	Message           string
	Prometheus        string
	Start             *time.Time
	Labels            models.LabelSet

	Remaining     models.LabelSet
	Documentation string
	Link          string

	Raw *models.GettableAlert
}
