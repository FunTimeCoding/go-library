package custom_field_choice

import "github.com/netbox-community/go-netbox/v4"

func New(v *netbox.CustomFieldChoiceSet) *Choice {
	return &Choice{Identifier: v.GetId(), Name: v.GetName(), Raw: v}
}
