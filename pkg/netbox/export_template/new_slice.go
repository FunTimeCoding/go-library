package export_template

import "github.com/netbox-community/go-netbox/v4"

func NewSlice(v []netbox.ExportTemplate) []*Template {
	var result []*Template

	for _, e := range v {
		result = append(result, New(&e))
	}

	return result
}
