package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/rear_port"
)

func (c *Client) MustRearPorts() []*rear_port.Port {
	result, e := c.RearPorts()
	errors.PanicOnError(e)

	return result
}
