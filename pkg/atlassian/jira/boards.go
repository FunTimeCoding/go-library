package jira

import "github.com/andygrunwald/go-jira"

func (c *Client) Boards() []*jira.Board {
	var result []*jira.Board
	response, r, e := c.client.Board.GetAllBoards(nil)
	panicOnError(r, e)

	for _, v := range response.Values {
		result = append(result, &v)
	}

	return result
}
