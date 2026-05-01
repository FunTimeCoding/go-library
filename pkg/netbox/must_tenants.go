package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/tenant"
)

func (c *Client) MustTenants() []*tenant.Tenant {
	result, e := c.Tenants()
	errors.PanicOnError(e)

	return result
}
