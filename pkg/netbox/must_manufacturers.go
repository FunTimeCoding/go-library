package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/manufacturer"
)

func (c *Client) MustManufacturers() []*manufacturer.Manufacturer {
	result, e := c.Manufacturers()
	errors.PanicOnError(e)

	return result
}
