package issue

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/issue/option"
)

func New(
	v *jira.Issue,
	o *option.Issue,
) *Issue {
	return &Issue{
		Key:         v.Key,
		Summary:     v.Fields.Summary,
		Description: v.Fields.Description,
		Link:        buildLink(o, v.Key),
		Raw:         v,
	}
}
