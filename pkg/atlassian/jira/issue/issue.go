package issue

import (
	"github.com/andygrunwald/go-jira"
	"time"
)

type Issue struct {
	MonitorIdentifier string
	Key               string
	Summary           string
	Description       string

	Link string

	Raw    *jira.Issue
	Create *time.Time
}
