package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/manufacturer"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) CreateManufacturer(name string) *manufacturer.Manufacturer {
	q := netbox.NewManufacturerRequest(
		name,
		slug(name),
	)
	result, r, e := c.client.DcimAPI.DcimManufacturersCreate(
		c.context,
	).ManufacturerRequest(*q).Execute()
	errors.PanicOnWebError(r, e)
	c.cache.Manufacturers = nil

	return manufacturer.New(result)
}
