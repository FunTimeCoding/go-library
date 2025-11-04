package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
	"github.com/funtimecoding/go-library/pkg/netbox/virtual_device_context"
)

func (c *Client) VirtualDeviceContexts() []*virtual_device_context.Context {
	result, r, e := c.client.DcimAPI.DcimVirtualDeviceContextsList(
		c.context,
	).Limit(constant.PageLimit).Execute()
	errors.PanicOnWebError(r, e)

	return virtual_device_context.NewSlice(result.Results)
}
