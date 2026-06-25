package technitium

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) MustCreateZone(
	name string,
	zoneType string,
) string {
	result, e := c.CreateZone(name, zoneType)
	errors.PanicOnError(e)

	return result
}
