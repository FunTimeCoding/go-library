package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/tenant"
	upstream "github.com/netbox-community/go-netbox/v4"
)

func (c *Client) CreateTenant(name string) *tenant.Tenant {
	q := upstream.NewTenantRequest(
		name,
		slug(name),
	)
	result, r, e := c.client.TenancyAPI.TenancyTenantsCreate(
		c.context,
	).TenantRequest(*q).Execute()
	errors.PanicOnWebError(r, e)
	c.cache.Tenants = nil

	return tenant.New(result)
}
