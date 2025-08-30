package custom_field_choice

import "github.com/netbox-community/go-netbox/v4"

type Choice struct {
	Identifier int32
	Name       string

	Raw *netbox.CustomFieldChoiceSet
}
