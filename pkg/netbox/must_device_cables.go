package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/cable"
)

func (c *Client) MustDeviceCables(device string) []*cable.Cable {
	result, e := c.DeviceCables(device)
	errors.PanicOnError(e)

	return result
}
