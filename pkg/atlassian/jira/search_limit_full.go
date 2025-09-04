package jira

import "github.com/funtimecoding/go-library/pkg/atlassian/jira/issue"

func (c *Client) SearchLimitFull(
	limit int,
	query string,
	a ...any,
) []*issue.Issue {
	var result []*issue.Issue

	for _, i := range c.SearchLimitV3(limit, query, a...) {
		result = append(result, c.Issue(i.Key))
	}

	return c.enrichMany(result)
}
