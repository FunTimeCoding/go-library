package jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) NewIssue(
	p *jira.MetaProject,
	projectKey string,
	issueType string,
	summary string,
	description string,
) *jira.Issue {
	result, e := jira.InitIssueWithMetaAndFields(
		p,
		p.GetIssueTypeWithName(issueType),
		map[string]string{
			"Project":     projectKey,
			"Issue Type":  issueType,
			"Summary":     summary,
			"Description": description,
		},
	)
	errors.PanicOnError(e)

	return result
}
