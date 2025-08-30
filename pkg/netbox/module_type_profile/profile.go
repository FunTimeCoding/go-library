package module_type_profile

import "github.com/netbox-community/go-netbox/v4"

type Profile struct {
	Identifier int32
	Name       string

	Raw *netbox.ModuleTypeProfile
}
