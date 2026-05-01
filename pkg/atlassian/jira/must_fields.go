package jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) MustFields() []jira.Field {
	result, e := c.Fields()
	errors.PanicOnError(e)

	return result
}
