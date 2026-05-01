package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/manufacturer"
)

func (c *Client) MustManufacturerByName(n string) *manufacturer.Manufacturer {
	result, e := c.ManufacturerByName(n)
	errors.PanicOnError(e)

	return result
}
