package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors/unexpected"
	"github.com/funtimecoding/go-library/pkg/netbox/rack"
)

func (c *Client) RackByName(n string) *rack.Rack {
	result := c.RacksByName(n)
	unexpected.Count(1, len(result))

	return result[0]
}
