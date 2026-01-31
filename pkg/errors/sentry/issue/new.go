package issue

import (
	"github.com/atlassian/go-sentry-api"
	"github.com/funtimecoding/go-library/pkg/monitor/item/constant"
)

func New(v *sentry.Issue) *Issue {
	return &Issue{
		MonitorIdentifier: constant.GoSentry.StringIdentifier(*v.ID),
		Project:           v.Project.Name,
		Type:              *v.Type,
		Title:             *v.Title,
		Link:              *v.Permalink,
		Create:            v.FirstSeen,
		Raw:               v,
	}
}
