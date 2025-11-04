package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
	"github.com/funtimecoding/go-library/pkg/netbox/power_outlet"
)

func (c *Client) PowerOutlets() []*power_outlet.Outlet {
	result, r, e := c.client.DcimAPI.DcimPowerOutletsList(
		c.context,
	).Limit(constant.PageLimit).Execute()
	errors.PanicOnWebError(r, e)

	return power_outlet.NewSlice(result.Results)
}
