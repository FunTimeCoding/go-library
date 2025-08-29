package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/tag"
)

func (c *Client) TagByName(n string) *tag.Tag {
	for _, element := range c.Tags() {
		if element.Name == n {
			return element
		}
	}

	errors.NotFound(n)

	return nil
}
