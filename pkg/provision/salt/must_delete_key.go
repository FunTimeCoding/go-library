package salt

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) MustDeleteKey(minion string) []string {
	result, e := c.DeleteKey(minion)
	errors.PanicOnError(e)

	return result
}
