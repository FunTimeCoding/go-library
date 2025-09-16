package jira

import "fmt"

func (c *Client) PrintIssueFields(
	projectKey string,
	issueType string,
) {
	for k, v := range c.IssueTypeFields(
		c.MetaIssueType(c.MetaProject(projectKey), issueType),
	) {
		fmt.Printf("Field: %s = %s\n", k, v)
	}
}
