package github

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/github/issue"
)

func (c *Client) MustSearchIssue(
	query string,
	a ...any,
) []*issue.Issue {
	result, e := c.SearchIssue(query, a...)
	errors.PanicOnError(e)

	return result
}
