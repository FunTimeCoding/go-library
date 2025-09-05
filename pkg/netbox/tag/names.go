package tag

import "github.com/netbox-community/go-netbox/v4"

func Names(v []netbox.NestedTag) []string {
	var result []string

	for _, n := range v {
		result = append(result, n.Name)
	}

	return result
}
