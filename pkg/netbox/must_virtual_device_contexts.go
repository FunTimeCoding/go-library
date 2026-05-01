package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/virtual_device_context"
)

func (c *Client) MustVirtualDeviceContexts() []*virtual_device_context.Context {
	result, e := c.VirtualDeviceContexts()
	errors.PanicOnError(e)

	return result
}
