package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/virtual_chassis"
)

func (c *Client) MustVirtualChassis() []*virtual_chassis.Chassis {
	result, e := c.VirtualChassis()
	errors.PanicOnError(e)

	return result
}
