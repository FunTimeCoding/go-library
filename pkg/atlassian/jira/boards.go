package jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) Boards() []*jira.Board {
	var result []*jira.Board
	response, _, e := c.client.Board.GetAllBoards(nil)
	errors.PanicOnError(e)

	for _, v := range response.Values {
		result = append(result, &v)
	}

	return result
}
