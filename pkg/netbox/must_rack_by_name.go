package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/rack"
)

func (c *Client) MustRackByName(n string) *rack.Rack {
	result, e := c.RackByName(n)
	errors.PanicOnError(e)

	return result
}
