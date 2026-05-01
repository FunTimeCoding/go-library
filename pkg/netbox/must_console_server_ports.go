package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/console_server_port"
)

func (c *Client) MustConsoleServerPorts() []*console_server_port.Port {
	result, e := c.ConsoleServerPorts()
	errors.PanicOnError(e)

	return result
}
