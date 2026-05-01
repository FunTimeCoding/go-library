package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/device_type"
	"github.com/funtimecoding/go-library/pkg/netbox/manufacturer"
)

func (c *Client) MustCreateDeviceType(
	model string,
	m *manufacturer.Manufacturer,
) *device_type.Type {
	result, e := c.CreateDeviceType(model, m)
	errors.PanicOnError(e)

	return result
}
