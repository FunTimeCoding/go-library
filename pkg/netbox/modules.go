package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
	"github.com/funtimecoding/go-library/pkg/netbox/module"
)

func (c *Client) Modules() []*module.Module {
	result, r, e := c.client.DcimAPI.DcimModulesList(
		c.context,
	).Limit(constant.PageLimit).Execute()
	errors.PanicOnWebError(r, e)

	return module.NewSlice(result.Results)
}
