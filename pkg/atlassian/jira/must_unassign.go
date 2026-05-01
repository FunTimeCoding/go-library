package jira

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) MustUnassign(key string) {
	errors.PanicOnError(c.Unassign(key))
}
