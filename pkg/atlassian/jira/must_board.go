package jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) MustBoard(identifier int) *jira.Board {
	result, e := c.Board(identifier)
	errors.PanicOnError(e)

	return result
}
