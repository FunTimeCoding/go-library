package jira

import "github.com/andygrunwald/go-jira"

func (c *Client) IssueTypeFields(t *jira.MetaIssueType) (map[string]string, error) {
	return t.GetAllFields()
}
