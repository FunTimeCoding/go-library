package gitlab

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) MustQuery(
	q string,
	response any,
) {
	errors.PanicOnError(c.Query(q, response))
}
