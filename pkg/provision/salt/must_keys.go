package salt

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/provision/salt/basic/response"
)

func (c *Client) MustKeys() *response.KeysReturn {
	result, e := c.Keys()
	errors.PanicOnError(e)

	return result
}
