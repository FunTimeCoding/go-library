package jira

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/basic/response"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) MustSearchV3(
	query string,
	a ...any,
) []*response.Issue {
	result, e := c.SearchV3(query, a...)
	errors.PanicOnError(e)

	return result
}
