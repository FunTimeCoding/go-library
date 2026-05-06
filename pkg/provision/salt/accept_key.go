package salt

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) AcceptKey(minion string) []string {
	result, e := c.client.AcceptKey(c.context, minion)
	errors.PanicOnError(e)

	return result
}
