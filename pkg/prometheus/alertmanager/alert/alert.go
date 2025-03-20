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

	Entity        string
	Category      string
	Documentation string
	Link          string
	Remaining     models.LabelSet

	Raw *models.GettableAlert
}
