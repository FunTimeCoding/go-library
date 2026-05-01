package jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) MustBoardConfiguration(identifier int) *jira.BoardConfiguration {
	result, e := c.BoardConfiguration(identifier)
	errors.PanicOnError(e)

	return result
}
