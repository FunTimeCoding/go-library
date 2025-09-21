package jira

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) DeleteIssue(key string) {
	_, e := c.client.Issue.Delete(key)
	errors.PanicOnError(e)
}
