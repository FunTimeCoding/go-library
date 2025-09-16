package jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) NewIssue(
	projectKey string,
	issueType string,
	summary string,
	description string,
) *jira.Issue {
	p := c.MetaProject(projectKey)
	result, e := jira.InitIssueWithMetaAndFields(
		p,
		p.GetIssueTypeWithName(issueType),
		map[string]string{
			constant.ProjectName:     projectKey,
			constant.IssueTypeName:   issueType,
			constant.SummaryName:     summary,
			constant.DescriptionName: description,
		},
	)
	errors.PanicOnError(e)

	return result
}
