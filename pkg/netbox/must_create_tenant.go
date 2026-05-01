package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/tenant"
)

func (c *Client) MustCreateTenant(name string) *tenant.Tenant {
	result, e := c.CreateTenant(name)
	errors.PanicOnError(e)

	return result
}
