package jira

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/issue"
)

func (c *Client) SearchLimit(
	limit int,
	query string,
	a ...any,
) []*issue.Issue {
	if len(a) > 0 {
		query = fmt.Sprintf(query, a...)
	}

	page, r, e := c.client.Issue.SearchV2JQL(
		query,
		&jira.SearchOptionsV2{Fields: []string{"*all"}, MaxResults: limit},
	)
	panicOnError(r, e)

	result := issue.NewSlice(page, c.IssueOption())

	return c.enrichMany(result)
}
