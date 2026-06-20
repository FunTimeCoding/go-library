package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"net"
)

func (c *Client) MustPhysicalExists(a net.HardwareAddr) bool {
	result, e := c.PhysicalExists(a)
	errors.PanicOnError(e)

	return result
}
