package salt

import (
	"github.com/daixijun/go-salt/v2"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) Keys() *salt.ListKeysReturn {
	result, e := c.client.ListKeys(c.context)
	errors.PanicOnError(e)

	return result
}
