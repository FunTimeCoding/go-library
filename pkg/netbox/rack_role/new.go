package rack_role

import "github.com/netbox-community/go-netbox/v4"

func New(v *netbox.RackRole) *Role {
	return &Role{Identifier: v.GetId(), Name: v.GetName(), Raw: v}
}
