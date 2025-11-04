package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
	"github.com/funtimecoding/go-library/pkg/netbox/module_bay"
)

func (c *Client) ModuleBays() []*module_bay.Bay {
	result, r, e := c.client.DcimAPI.DcimModuleBaysList(
		c.context,
	).Limit(constant.PageLimit).Execute()
	errors.PanicOnWebError(r, e)

	return module_bay.NewSlice(result.Results)
}
