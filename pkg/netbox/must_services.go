package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/service"
)

func (c *Client) MustServices() []*service.Service {
	result, e := c.Services()
	errors.PanicOnError(e)

	return result
}
