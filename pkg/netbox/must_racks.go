package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/rack"
)

func (c *Client) MustRacks() []*rack.Rack {
	result, e := c.Racks()
	errors.PanicOnError(e)

	return result
}
