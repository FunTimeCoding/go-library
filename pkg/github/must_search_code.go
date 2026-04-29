package github

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/github/code"
)

func (c *Client) MustSearchCode(
	query string,
	a ...any,
) []*code.Code {
	result, e := c.SearchCode(query, a...)
	errors.PanicOnError(e)

	return result
}
