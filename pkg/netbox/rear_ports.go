package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
	"github.com/funtimecoding/go-library/pkg/netbox/rear_port"
)

func (c *Client) RearPorts() []*rear_port.Port {
	result, _, e := c.client.DcimAPI.DcimRearPortsList(
		c.context,
	).Limit(constant.PageLimit).Execute()
	errors.PanicOnError(e)

	return rear_port.NewSlice(result.Results)
}
