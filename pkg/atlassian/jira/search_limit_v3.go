package jira

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/basic_client/response"
)

func (c *Client) SearchLimitV3(
	limit int,
	query string,
	a ...any,
) []*response.Issue {
	if len(a) > 0 {
		query = fmt.Sprintf(query, a...)
	}

	return c.searchV3Page(limit, "", query).Issues
}
