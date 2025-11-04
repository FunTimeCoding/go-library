package jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/go-library/pkg/integers"
)

func (c *Client) BoardSprints(identifier int) []*jira.Sprint {
	response, r, e := c.client.Board.GetAllSprints(
		integers.ToString(identifier),
	)
	panicOnError(r, e)
	var result []*jira.Sprint

	for _, s := range response {
		result = append(result, &s)
	}

	return result
}
