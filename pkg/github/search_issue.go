package github

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/github/issue"
)

func (c *Client) SearchIssue(
	query string,
	a ...any,
) []*issue.Issue {
	if len(a) > 0 {
		query = fmt.Sprintf(query, a...)
	}

	result, r, e := c.client.Search.Issues(c.context, query, nil)
	panicOnError(r, e)

	return issue.NewSlice(result.Issues)
}
