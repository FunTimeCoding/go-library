package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/cable"
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
)

func (c *Client) Cables() []*cable.Cable {
	result, r, e := c.client.DcimAPI.DcimCablesList(
		c.context,
	).Limit(constant.PageLimit).Execute()
	errors.PanicOnWebError(r, e)

	return cable.NewSlice(result.Results)
}
