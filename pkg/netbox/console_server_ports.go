package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/console_server_port"
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
)

func (c *Client) ConsoleServerPorts() []*console_server_port.Port {
	result, r, e := c.client.DcimAPI.DcimConsoleServerPortsList(
		c.context,
	).Limit(constant.PageLimit).Execute()
	errors.PanicOnWebError(r, e)

	return console_server_port.NewSlice(result.Results)
}
