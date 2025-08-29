package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
	"github.com/funtimecoding/go-library/pkg/netbox/device_type"
)

func (c *Client) DeviceTypes() []*device_type.DeviceType {
	if len(c.cache.DeviceTypes) != 0 {
		return c.cache.DeviceTypes
	}

	result, _, e := c.client.DcimAPI.DcimDeviceTypesList(
		c.context,
	).Limit(constant.PageLimit).Execute()
	errors.PanicOnError(e)
	c.cache.DeviceTypes = device_type.NewSlice(result.Results)

	return c.cache.DeviceTypes
}
