package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
	"github.com/funtimecoding/go-library/pkg/netbox/front_port"
)

func (c *Client) FrontPorts() []*front_port.Port {
	result, _, e := c.client.DcimAPI.DcimFrontPortsList(
		c.context,
	).Limit(constant.PageLimit).Execute()
	errors.PanicOnError(e)

	return front_port.NewSlice(result.Results)
}
