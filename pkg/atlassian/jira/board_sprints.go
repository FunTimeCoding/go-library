package jira

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) BoardSprints(identifier int) []*jira.Sprint {
	response, _, e := c.client.Board.GetAllSprints(
		fmt.Sprintf("%d", identifier),
	)
	errors.PanicOnError(e)
	var result []*jira.Sprint

	for _, r := range response {
		result = append(result, &r)
	}

	return result
}
