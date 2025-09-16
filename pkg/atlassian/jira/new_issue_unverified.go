package jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/issue"
)

func (c *Client) NewIssueUnverified(
	projectKey string,
	issueType string,
	summary string,
	description string,
) *jira.Issue {
	result := issue.RawStub()
	result.Fields.Reporter = c.User()
	result.Fields.Project = jira.Project{Key: projectKey}
	result.Fields.Type = jira.IssueType{Name: issueType}
	result.Fields.Summary = summary
	result.Fields.Description = description

	return result
}
