package contact_role

import "github.com/netbox-community/go-netbox/v4"

func New(v *netbox.ContactRole) *Role {
	return &Role{Identifier: v.GetId(), Name: v.GetName(), Raw: v}
}
