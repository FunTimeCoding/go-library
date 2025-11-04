package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
	"github.com/funtimecoding/go-library/pkg/netbox/tenant_group"
)

func (c *Client) TenantGroups() []*tenant_group.Group {
	result, r, e := c.client.TenancyAPI.TenancyTenantGroupsList(
		c.context,
	).Limit(constant.PageLimit).Execute()
	errors.PanicOnWebError(r, e)

	return tenant_group.NewSlice(result.Results)
}
