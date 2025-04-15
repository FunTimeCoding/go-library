package jira

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) WatchedIssueKeys() []string {
	issues, _, e := c.client.Issue.Search(
		"issue IN watchedIssues() ORDER BY key ASC",
		nil,
	)
	errors.PanicOnError(e)
	var result []string

	for _, i := range issues {
		result = append(result, i.Key)
	}

	return result
}
