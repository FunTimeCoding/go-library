package salt

import (
	"github.com/daixijun/go-salt/v2"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) MustLocal(
	glob string,
	function string,
	a []string,
) map[string]salt.LocalClientReturn {
	result, e := c.Local(glob, function, a)
	errors.PanicOnError(e)

	return result
}
