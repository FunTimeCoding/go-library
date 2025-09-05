package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/tag"
)

func (c *Client) TagByName(n string) *tag.Tag {
	for _, t := range c.Tags() {
		if t.Name == n {
			return t
		}
	}

	errors.NotFound(n)

	return nil
}
