package custom_field

import "github.com/netbox-community/go-netbox/v4"

func New(v *netbox.CustomField) *Field {
	return &Field{Identifier: v.GetId(), Name: v.GetName(), Raw: v}
}
