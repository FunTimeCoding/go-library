package jira

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/basic/response"
	"github.com/funtimecoding/go-library/pkg/notation"
)

func (c *Client) searchV3Page(
	maximumResults int,
	nextPageToken string,
	query string,
) *response.Search {
	var result response.Search
	status, r := c.basic.Get(
		c.basic.Base().Copy().Base(
			"/rest/api/3",
		).Path("/search/jql").Set(
			"fields",
			"*all",
		).Set(
			"maxResults",
			fmt.Sprintf("%d", maximumResults),
		).Set(
			"nextPageToken",
			nextPageToken,
		).Set("jql", query).String(),
	)
	notation.DecodeStrict(r, &result, true)

	if false {
		fmt.Printf("Response: %d %s\n", status, r)
	}

	return &result
}
