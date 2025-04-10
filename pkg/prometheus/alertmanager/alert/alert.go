package alert

import (
	"github.com/funtimecoding/go-library/pkg/face"
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
	Tags      []string
	Runbook   string
	Link      string
	Remaining models.LabelSet

	instance face.StringAlias

	Raw *models.GettableAlert
}
