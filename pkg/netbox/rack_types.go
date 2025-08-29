package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
	"github.com/funtimecoding/go-library/pkg/netbox/rack_type"
)

func (c *Client) RackTypes() []*rack_type.Type {
	result, _, e := c.client.DcimAPI.DcimRackTypesList(
		c.context,
	).Limit(constant.PageLimit).Execute()
	errors.PanicOnError(e)

	return rack_type.NewSlice(result.Results)
}
