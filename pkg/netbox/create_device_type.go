package netbox

import (
	"github.com/funtimecoding/go-library/pkg/netbox/device_type"
	"github.com/funtimecoding/go-library/pkg/netbox/manufacturer"
	upstream "github.com/netbox-community/go-netbox/v4"
)

func (c *Client) CreateDeviceType(
	model string,
	m *manufacturer.Manufacturer,
) (*device_type.Type, error) {
	q := upstream.NewWritableDeviceTypeRequest(
		upstream.BriefManufacturerRequestAsBriefDeviceTypeRequestManufacturer(
			upstream.NewBriefManufacturerRequest(m.Name, m.Raw.Slug),
		),
		model,
		slug(model),
	)
	result, _, e := c.client.DcimAPI.DcimDeviceTypesCreate(
		c.context,
	).WritableDeviceTypeRequest(*q).Execute()

	if e != nil {
		return nil, e
	}

	c.cache.DeviceTypes = nil

	return device_type.New(result), nil
}
