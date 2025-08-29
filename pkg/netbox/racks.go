package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
	"github.com/funtimecoding/go-library/pkg/netbox/rack"
)

func (c *Client) Racks() []*rack.Rack {
	if len(c.cache.Racks) != 0 {
		return c.cache.Racks
	}

	result, _, e := c.client.DcimAPI.DcimRacksList(
		c.context,
	).Limit(constant.PageLimit).Execute()
	errors.PanicOnError(e)
	c.cache.Racks = rack.NewSlice(result.Results)

	return c.cache.Racks
}
