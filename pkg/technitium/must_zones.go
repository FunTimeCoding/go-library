package technitium

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/technitium/zone"
)

func (c *Client) MustZones() []*zone.Zone {
	result, e := c.Zones()
	errors.PanicOnError(e)

	return result
}
