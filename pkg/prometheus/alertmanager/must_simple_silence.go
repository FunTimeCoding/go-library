package alertmanager

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) MustSimpleSilence(alert string) string {
	result, e := c.SimpleSilence(alert)
	errors.PanicOnError(e)

	return result
}
