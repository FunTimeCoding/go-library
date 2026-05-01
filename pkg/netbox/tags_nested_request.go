package netbox

import (
	"github.com/netbox-community/go-netbox/v4"
	"slices"
)

func (c *Client) tagsNestedRequest(v []string) ([]netbox.NestedTagRequest, error) {
	tags, e := c.Tags()

	if e != nil {
		return nil, e
	}

	var result []netbox.NestedTagRequest

	for _, t := range tags {
		if slices.Contains(v, t.Name) {
			result = append(result, *t.Nested)
		}
	}

	return result, nil
}
