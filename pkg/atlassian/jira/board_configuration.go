package jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) BoardConfiguration(identifier int) *jira.BoardConfiguration {
	result, _, e := c.client.Board.GetBoardConfiguration(identifier)
	errors.PanicOnError(e)

	return result
}
