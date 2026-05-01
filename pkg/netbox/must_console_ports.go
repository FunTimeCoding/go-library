package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/console_port"
)

func (c *Client) MustConsolePorts() []*console_port.Port {
	result, e := c.ConsolePorts()
	errors.PanicOnError(e)

	return result
}
