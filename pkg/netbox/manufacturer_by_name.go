package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/manufacturer"
)

func (c *Client) ManufacturerByName(n string) *manufacturer.Manufacturer {
	for _, m := range c.Manufacturers() {
		if m.Name == n {
			return m
		}
	}

	errors.NotFound(n)

	return nil
}
