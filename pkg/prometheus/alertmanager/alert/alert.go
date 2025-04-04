package alert

import (
	"github.com/prometheus/alertmanager/api/v2/models"
	"time"
)

type Alert struct {
	MonitorIdentifier string
	Fingerprint       string
	Name              string
	State             string
	Severity          string
	Summary           string
	Message           string
	Prometheus        string
	Start             *time.Time
	Labels            models.LabelSet
	Receivers         []string

	Entity    string
	Category  string
	Runbook   string
	Link      string
	Remaining models.LabelSet

	Raw *models.GettableAlert
}
