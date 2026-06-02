package alertmanager

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) MustRunbook(name string) string {
	result, e := c.Runbook(name)
	errors.PanicOnError(e)

	return result
}
