package tenant_group

import "github.com/netbox-community/go-netbox/v4"

func New(v *netbox.TenantGroup) *Group {
	return &Group{Identifier: v.GetId(), Name: v.GetName(), Raw: v}
}
