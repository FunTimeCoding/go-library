package jira

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/basic/response"
	"github.com/funtimecoding/go-library/pkg/notation"
	"net/url"
)

func (c *Client) searchV3Page(
	maximumResults int,
	nextPageToken string,
	query string,
) *response.Search {
	var result response.Search
	status, r := c.basic.Get(
		fmt.Sprintf(
			"/rest/api/3/search/jql?fields=*all&maxResults=%d&nextPageToken=%s&jql=%s",
			maximumResults,
			nextPageToken,
			url.QueryEscape(query),
		),
	)
	notation.DecodeStrict(r, &result, true)

	if false {
		fmt.Printf("Response: %d %s\n", status, r)
	}

	return &result
}
