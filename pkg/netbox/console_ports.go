package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/console_port"
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
)

func (c *Client) ConsolePorts() []*console_port.Port {
	result, _, e := c.client.DcimAPI.DcimConsolePortsList(
		c.context,
	).Limit(constant.PageLimit).Execute()
	errors.PanicOnError(e)

	return console_port.NewSlice(result.Results)
}
