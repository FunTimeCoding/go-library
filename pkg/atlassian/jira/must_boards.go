package jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) MustBoards() []*jira.Board {
	result, e := c.Boards()
	errors.PanicOnError(e)

	return result
}
