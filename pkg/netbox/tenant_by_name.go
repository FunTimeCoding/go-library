package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/tenant"
)

func (c *Client) TenantByName(n string) *tenant.Tenant {
	for _, element := range c.Tenants() {
		if element.Name == n {
			return element
		}
	}

	errors.NotFound(n)

	return nil
}
