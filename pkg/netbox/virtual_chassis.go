package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
	"github.com/funtimecoding/go-library/pkg/netbox/virtual_chassis"
)

func (c *Client) VirtualChassis() []*virtual_chassis.Chassis {
	result, _, e := c.client.DcimAPI.DcimVirtualChassisList(
		c.context,
	).Limit(constant.PageLimit).Execute()
	errors.PanicOnError(e)

	return virtual_chassis.NewSlice(result.Results)
}
