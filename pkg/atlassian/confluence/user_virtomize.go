package confluence

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	virtomize "github.com/virtomize/confluence-go-api"
)

func (c *Client) UserVirtomize() *virtomize.User {
	result, e := c.virtomize.CurrentUser()
	errors.PanicOnError(e)

	return result
}
