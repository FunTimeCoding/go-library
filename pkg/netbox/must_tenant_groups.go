package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/tenant_group"
)

func (c *Client) MustTenantGroups() []*tenant_group.Group {
	result, e := c.TenantGroups()
	errors.PanicOnError(e)

	return result
}
