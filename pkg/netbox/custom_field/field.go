package custom_field

import "github.com/netbox-community/go-netbox/v4"

type Field struct {
	Identifier int32
	Name       string

	Raw *netbox.CustomField
}
