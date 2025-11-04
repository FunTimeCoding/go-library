package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/manufacturer"
)

func (c *Client) Manufacturers() []*manufacturer.Manufacturer {
	if len(c.cache.Manufacturers) != 0 {
		return c.cache.Manufacturers
	}

	result, r, e := c.client.DcimAPI.DcimManufacturersList(
		c.context,
	).Execute()
	errors.PanicOnWebError(r, e)
	c.cache.Manufacturers = manufacturer.NewSlice(result.Results)

	return c.cache.Manufacturers
}
