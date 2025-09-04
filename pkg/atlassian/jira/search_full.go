package jira

import "github.com/funtimecoding/go-library/pkg/atlassian/jira/issue"

func (c *Client) SearchFull(
	query string,
	a ...any,
) []*issue.Issue {
	// andygrunwald/go-jira does not yet support new search API
	// https://developer.atlassian.com/changelog/#CHANGE-2046
	// https://github.com/andygrunwald/go-jira/issues/715
	var result []*issue.Issue

	for _, i := range c.SearchV3(query, a...) {
		result = append(result, c.Issue(i.Key))
	}

	return c.enrichMany(result)
}
