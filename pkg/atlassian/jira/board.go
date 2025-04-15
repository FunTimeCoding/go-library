package jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) Board(identifier int) *jira.Board {
	result, _, e := c.client.Board.GetBoard(identifier)
	errors.PanicOnError(e)

	return result
}
