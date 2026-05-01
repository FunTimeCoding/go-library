package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/location"
)

func (c *Client) MustLocations() []*location.Location {
	result, e := c.Locations()
	errors.PanicOnError(e)

	return result
}
