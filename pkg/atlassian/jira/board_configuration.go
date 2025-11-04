package jira

import "github.com/andygrunwald/go-jira"

func (c *Client) BoardConfiguration(identifier int) *jira.BoardConfiguration {
	result, r, e := c.client.Board.GetBoardConfiguration(identifier)
	panicOnError(r, e)

	return result
}
