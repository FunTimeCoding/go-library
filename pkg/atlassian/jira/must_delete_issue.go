package jira

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) MustDeleteIssue(key string) {
	errors.PanicOnError(c.DeleteIssue(key))
}
