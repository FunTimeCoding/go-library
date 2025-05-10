package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"gitlab.com/gitlab-org/api/client-go"
)

func (c *Client) Query(
	q string,
	r any,
) *gitlab.Response {
	result, e := c.client.GraphQL.Do(
		c.context,
		gitlab.GraphQLQuery{Query: q},
		r,
	)
	errors.PanicOnError(e)

	return result
}
