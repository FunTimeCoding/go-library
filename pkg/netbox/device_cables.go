package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/cable"
)

func (c *Client) DeviceCables(device string) []*cable.Cable {
	result, r, e := c.client.DcimAPI.DcimCablesList(
		c.context,
	).Device([]string{device}).Execute()
	errors.PanicOnWebError(r, e)

	return cable.NewSlice(result.Results)
}
