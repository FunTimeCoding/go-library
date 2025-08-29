package netbox

import (
	"github.com/netbox-community/go-netbox/v4"
	"slices"
)

func (c *Client) tagsNestedRequest(v []string) []netbox.NestedTagRequest {
	var result []netbox.NestedTagRequest

	for _, t := range c.Tags() {
		if slices.Contains(v, t.Name) {
			result = append(result, *t.Nested)
		}
	}

	return result
}
