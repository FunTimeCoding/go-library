package custom_field_choice

import "github.com/netbox-community/go-netbox/v4"

func NewSlice(v []netbox.CustomFieldChoiceSet) []*Choice {
	var result []*Choice

	for _, e := range v {
		result = append(result, New(&e))
	}

	return result
}
