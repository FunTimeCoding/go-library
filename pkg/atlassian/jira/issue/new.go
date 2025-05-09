package issue

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/issue/option"
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	"strings"
	"time"
)

func New(
	v *jira.Issue,
	o *option.Issue,
) *Issue {
	create := time.Time(v.Fields.Created)
	due := time.Time(v.Fields.Duedate)
	result := Stub()
	result.MonitorIdentifier = fmt.Sprintf(
		"%s-%s",
		constant.JiraPrefix,
		v.Key,
	)
	result.Key = v.Key
	result.Summary = strings.TrimSpace(v.Fields.Summary)
	result.Description = strings.TrimSpace(v.Fields.Description)
	result.Initials = initialsField(v)
	result.Status = statusField(v)
	result.Priority = priorityField(v)
	result.Type = v.Fields.Type.Name
	result.Create = &create
	result.Due = &due
	result.Link = BuildLink(o.Locator, v.Key)
	result.Labels = v.Fields.Labels
	result.fieldMap = o.FieldMap
	result.option = o
	result.Raw = v

	return result
}
