package issue

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/issue/option"
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	"time"
)

func New(
	v *jira.Issue,
	o *option.Issue,
) *Issue {
	t := time.Time(v.Fields.Created)

	return &Issue{
		MonitorIdentifier: fmt.Sprintf(
			"%s-%s",
			constant.JiraPrefix,
			v.Key,
		),
		Key:         v.Key,
		Summary:     v.Fields.Summary,
		Description: v.Fields.Description,
		Create:      &t,
		Link:        buildLink(o, v.Key),
		Raw:         v,
	}
}
