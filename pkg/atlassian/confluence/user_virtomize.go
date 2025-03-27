package confluence

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/virtomize/confluence-go-api"
)

func (c *Client) UserVirtomize() *goconfluence.User {
	result, e := c.virtomize.CurrentUser()
	errors.PanicOnError(e)

	return result
}
