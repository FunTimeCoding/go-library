package github

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/github/issue"
)

func (c *Client) Search(query string) []*issue.Issue {
	result, _, e := c.client.Search.Issues(c.context, query, nil)
	errors.PanicOnError(e)

	return issue.NewSlice(result.Issues)
}
