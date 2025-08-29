package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors/unexpected"
	"github.com/funtimecoding/go-library/pkg/netbox/prefix"
)

func (c *Client) PrefixByDescription(d string) *prefix.Prefix {
	var result []*prefix.Prefix

	for _, element := range c.Prefixes() {
		if element.Description == d {
			result = append(result, element)
		}
	}

	unexpected.Count(1, len(result))

	return result[0]
}
