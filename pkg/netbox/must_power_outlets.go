package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/power_outlet"
)

func (c *Client) MustPowerOutlets() []*power_outlet.Outlet {
	result, e := c.PowerOutlets()
	errors.PanicOnError(e)

	return result
}
