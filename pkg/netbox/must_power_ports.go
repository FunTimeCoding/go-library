package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/power_port"
)

func (c *Client) MustPowerPorts() []*power_port.Port {
	result, e := c.PowerPorts()
	errors.PanicOnError(e)

	return result
}
