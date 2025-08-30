package module_type_profile

import "github.com/netbox-community/go-netbox/v4"

func New(v *netbox.ModuleTypeProfile) *Profile {
	return &Profile{Identifier: v.GetId(), Name: v.GetName(), Raw: v}
}
