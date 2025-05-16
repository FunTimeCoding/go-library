package jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) User() *jira.User {
	result, _, e := c.client.User.GetSelf()
	errors.PanicOnError(e)

	return result
}
