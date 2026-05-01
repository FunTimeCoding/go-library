package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/tenant"
)

func (c *Client) MustTenantByName(n string) *tenant.Tenant {
	result, e := c.TenantByName(n)
	errors.PanicOnError(e)

	return result
}
