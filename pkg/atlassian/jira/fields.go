package jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) Fields() []jira.Field {
	result, _, e := c.client.Field.GetList()
	errors.PanicOnError(e)

	return result
}
