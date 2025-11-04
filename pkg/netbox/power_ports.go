package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
	"github.com/funtimecoding/go-library/pkg/netbox/power_port"
)

func (c *Client) PowerPorts() []*power_port.Port {
	result, r, e := c.client.DcimAPI.DcimPowerPortsList(
		c.context,
	).Limit(constant.PageLimit).Execute()
	errors.PanicOnWebError(r, e)

	return power_port.NewSlice(result.Results)
}
