package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/tenant"
)

func (c *Client) TenantByName(n string) *tenant.Tenant {
	for _, t := range c.Tenants() {
		if t.Name == n {
			return t
		}
	}

	errors.NotFound(n)

	return nil
}
