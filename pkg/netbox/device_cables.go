package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/cable"
)

func (c *Client) DeviceCables(device string) []*cable.Cable {
	result, _, e := c.client.DcimAPI.DcimCablesList(
		c.context,
	).Device([]string{device}).Execute()
	errors.PanicOnError(e)

	return cable.NewSlice(result.Results)
}
