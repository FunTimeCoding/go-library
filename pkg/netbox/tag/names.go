package tag

import "github.com/netbox-community/go-netbox/v4"

func Names(v []netbox.NestedTag) []string {
	var result []string

	for _, element := range v {
		result = append(result, element.Name)
	}

	return result
}
