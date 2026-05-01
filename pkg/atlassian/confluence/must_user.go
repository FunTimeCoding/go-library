package confluence

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/user"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) MustUser() *user.User {
	result, e := c.User()
	errors.PanicOnError(e)

	return result
}
