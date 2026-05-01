package jira

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) MustSearchFull(
	query string,
	a ...any,
) []*issue.Issue {
	result, e := c.SearchFull(query, a...)
	errors.PanicOnError(e)

	return result
}
