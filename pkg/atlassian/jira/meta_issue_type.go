package jira

import (
	"github.com/andygrunwald/go-jira"
	"log"
)

func (c *Client) MetaIssueType(
	p *jira.MetaProject,
	issueType string,
) *jira.MetaIssueType {
	result := p.GetIssueTypeWithName(issueType)

	if result == nil {
		log.Panicf("issue type not found: %s", issueType)
	}

	return result
}
