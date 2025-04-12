package issue

import (
	"fmt"
	"github.com/atlassian/go-sentry-api"
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
)

func New(v *sentry.Issue) *Issue {
	return &Issue{
		MonitorIdentifier: fmt.Sprintf(
			"%s-%s",
			constant.SentryPrefix,
			*v.ID,
		),
		Project: v.Project.Name,
		Type:    *v.Type,
		Title:   *v.Title,
		Link:    *v.Permalink,
		Create:  v.FirstSeen,
		Raw:     v,
	}
}
