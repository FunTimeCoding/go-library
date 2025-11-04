package jira

import "github.com/andygrunwald/go-jira"

func (c *Client) Board(identifier int) *jira.Board {
	result, r, e := c.client.Board.GetBoard(identifier)
	panicOnError(r, e)

	return result
}
