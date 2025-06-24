package github

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/github/repository"
)

func (c *Client) SearchRepository(
	query string,
	a ...any,
) []*repository.Repository {
	if len(a) > 0 {
		query = fmt.Sprintf(query, a...)
	}

	result, _, e := c.client.Search.Repositories(c.context, query, nil)
	errors.PanicOnError(e)

	return repository.NewSlice(result.Repositories)
}
