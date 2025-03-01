package issue

import (
	"github.com/atlassian/go-sentry-api"
	"time"
)

type Issue struct {
	MonitorIdentifier string
	Project           string
	Type              string
	Title             string
	Link              string
	Create            time.Time

	Raw *sentry.Issue
}
