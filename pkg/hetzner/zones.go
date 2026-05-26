package hetzner

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/hetzner/zone"
)

func (c *Client) Zones() []*zone.Zone {
	result, e := c.client.Zone.All(c.context)
	errors.PanicOnError(e)

	return zone.NewSlice(result)
}
