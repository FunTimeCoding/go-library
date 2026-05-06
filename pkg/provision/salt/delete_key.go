package salt

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) DeleteKey(minion string) []string {
	result, e := c.client.DeleteKey(c.context, minion)
	errors.PanicOnError(e)

	return result
}
