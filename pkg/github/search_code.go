package github

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/github/code"
)

func (c *Client) SearchCode(
	query string,
	a ...any,
) []*code.Code {
	if len(a) > 0 {
		query = fmt.Sprintf(query, a...)
	}

	result, r, e := c.client.Search.Code(c.context, query, nil)
	panicOnError(r, e)

	return code.NewSlice(result.CodeResults)
}
