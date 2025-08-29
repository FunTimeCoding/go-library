package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/device"
	"github.com/funtimecoding/go-library/pkg/netbox/device_role"
	"github.com/funtimecoding/go-library/pkg/netbox/device_type"
	"github.com/funtimecoding/go-library/pkg/netbox/site"
	"github.com/funtimecoding/go-library/pkg/netbox/tenant"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) CreateDevice(
	name string,
	o *device_role.DeviceRole,
	tags []string,
	y *device_type.DeviceType,
	s *site.Site,
	n *tenant.Tenant,
) *device.Device {
	r := netbox.NewWritableDeviceWithConfigContextRequest(
		netbox.BriefDeviceTypeRequestAsDeviceBayTemplateRequestDeviceType(
			netbox.NewBriefDeviceTypeRequest(
				netbox.BriefManufacturerRequestAsBriefDeviceTypeRequestManufacturer(
					netbox.NewBriefManufacturerRequest(
						y.Raw.Manufacturer.Name,
						y.Raw.Manufacturer.Slug,
					),
				),
				y.Model,
				y.Raw.Slug,
			),
		),
		netbox.BriefDeviceRoleRequestAsDeviceWithConfigContextRequestRole(
			netbox.NewBriefDeviceRoleRequest(o.Name, o.Raw.Slug),
		),
		netbox.BriefSiteRequestAsDeviceWithConfigContextRequestSite(
			netbox.NewBriefSiteRequest(s.Name, s.Raw.Slug),
		),
	)
	r.SetTenant(
		netbox.BriefTenantRequestAsASNRangeRequestTenant(
			netbox.NewBriefTenantRequest(n.Name, n.Raw.Slug),
		),
	)
	r.SetName(name)

	if len(tags) > 0 {
		r.SetTags(c.tagsNestedRequest(tags))
	}

	r.SetStatus(device.ActiveStatus)
	result, _, e := c.client.DcimAPI.DcimDevicesCreate(
		c.context,
	).WritableDeviceWithConfigContextRequest(*r).Execute()
	errors.PanicOnError(e)

	return device.New(result)
}
