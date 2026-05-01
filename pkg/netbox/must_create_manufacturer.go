package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/manufacturer"
)

func (c *Client) MustCreateManufacturer(name string) *manufacturer.Manufacturer {
	result, e := c.CreateManufacturer(name)
	errors.PanicOnError(e)

	return result
}
