package jira

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
)

func (c *Client) MetaIssueType(
	p *jira.MetaProject,
	issueType string,
) (*jira.MetaIssueType, error) {
	result := p.GetIssueTypeWithName(issueType)

	if result == nil {
		return nil, fmt.Errorf("issue type not found: %s", issueType)
	}

	return result, nil
}
