package jira

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) MustWatchedIssueKeys() []string {
	result, e := c.WatchedIssueKeys()
	errors.PanicOnError(e)

	return result
}
