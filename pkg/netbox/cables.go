package netbox

import (
	"github.com/funtimecoding/go-library/pkg/netbox/cable"
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
)

func (c *Client) Cables() ([]*cable.Cable, error) {
	result, _, e := c.client.DcimAPI.DcimCablesList(
		c.context,
	).Limit(constant.PageLimit).Execute()

	if e != nil {
		return nil, e
	}

	return cable.NewSlice(result.Results), nil
}
