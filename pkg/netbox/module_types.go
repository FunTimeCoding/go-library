package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
	"github.com/funtimecoding/go-library/pkg/netbox/module_type"
)

func (c *Client) ModuleTypes() []*module_type.Type {
	result, r, e := c.client.DcimAPI.DcimModuleTypesList(
		c.context,
	).Limit(constant.PageLimit).Execute()
	errors.PanicOnWebError(r, e)

	return module_type.NewSlice(result.Results)
}
