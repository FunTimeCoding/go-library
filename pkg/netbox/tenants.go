package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
	"github.com/funtimecoding/go-library/pkg/netbox/tenant"
)

func (c *Client) Tenants() []*tenant.Tenant {
	if len(c.cache.Tenants) != 0 {
		return c.cache.Tenants
	}

	result, r, e := c.client.TenancyAPI.TenancyTenantsList(
		c.context,
	).Limit(constant.PageLimit).Execute()
	errors.PanicOnWebError(r, e)
	c.cache.Tenants = tenant.NewSlice(result.Results)

	return c.cache.Tenants
}
