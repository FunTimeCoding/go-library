package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/rack_type"
)

func (c *Client) MustRackTypes() []*rack_type.Type {
	result, e := c.RackTypes()
	errors.PanicOnError(e)

	return result
}
