package github

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/github/code"
)

func (c *Client) SearchCode(
	query string,
	a ...any,
) []*code.Code {
	if len(a) > 0 {
		query = fmt.Sprintf(query, a...)
	}

	result, _, e := c.client.Search.Code(c.context, query, nil)
	errors.PanicOnError(e)

	return code.NewSlice(result.CodeResults)
}
