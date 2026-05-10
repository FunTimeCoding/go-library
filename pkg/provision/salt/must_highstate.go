package salt

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/provision/salt/basic/response"
)

func (c *Client) MustHighstate(
	target string,
) map[string]response.LocalReturn {
	result, e := c.Highstate(target)
	errors.PanicOnError(e)

	return result
}
