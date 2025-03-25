package issue

import (
	"github.com/atlassian/go-sentry-api"
	"github.com/funtimecoding/go-library/pkg/face"
	"time"
)

type Issue struct {
	Project string
	Type    string
	Title   string
	Link    string
	Create  time.Time

	MonitorIdentifier string
	ageColor          face.SprintFunction

	Raw *sentry.Issue
}
