package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
	"github.com/funtimecoding/go-library/pkg/netbox/location"
)

func (c *Client) Locations() []*location.Location {
	result, r, e := c.client.DcimAPI.DcimLocationsList(
		c.context,
	).Limit(constant.PageLimit).Execute()
	errors.PanicOnWebError(r, e)

	return location.NewSlice(result.Results)
}
