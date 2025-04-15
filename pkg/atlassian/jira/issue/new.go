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
	result := Stub()
	result.MonitorIdentifier = fmt.Sprintf(
		"%s-%s",
		constant.JiraPrefix,
		v.Key,
	)
	result.Key = v.Key
	result.Summary = v.Fields.Summary
	result.Description = v.Fields.Description
	result.Initials = initialsField(v)
	result.Status = statusField(v)
	result.Type = v.Fields.Type.Name
	result.Create = &t
	result.Link = buildLink(o, v.Key)
	result.Raw = v

	return result
}
