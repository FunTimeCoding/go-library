package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors/unexpected"
	"github.com/funtimecoding/go-library/pkg/netbox/prefix"
)

func (c *Client) PrefixByDescription(d string) *prefix.Prefix {
	var result []*prefix.Prefix

	for _, p := range c.Prefixes() {
		if p.Description == d {
			result = append(result, p)
		}
	}

	unexpected.Count(1, len(result))

	return result[0]
}
