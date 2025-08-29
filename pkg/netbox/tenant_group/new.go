package tenant_group

import "github.com/netbox-community/go-netbox/v4"

func New(d *netbox.TenantGroup) *Group {
	return &Group{Identifier: d.GetId(), Name: d.GetName(), Raw: d}
}
