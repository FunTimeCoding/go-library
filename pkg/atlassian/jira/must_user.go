package jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) MustUser() *jira.User {
	result, e := c.User()
	errors.PanicOnError(e)

	return result
}
