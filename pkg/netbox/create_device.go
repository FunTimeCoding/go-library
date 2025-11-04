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
	o *device_role.Role,
	tags []string,
	y *device_type.Type,
	s *site.Site,
	n *tenant.Tenant,
) *device.Device {
	q := netbox.NewWritableDeviceWithConfigContextRequest(
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

	if n != nil {
		q.SetTenant(
			netbox.BriefTenantRequestAsASNRangeRequestTenant(
				netbox.NewBriefTenantRequest(n.Name, n.Raw.Slug),
			),
		)
	}

	q.SetName(name)

	if len(tags) > 0 {
		q.SetTags(c.tagsNestedRequest(tags))
	}

	q.SetStatus(device.ActiveStatus)
	result, r, e := c.client.DcimAPI.DcimDevicesCreate(
		c.context,
	).WritableDeviceWithConfigContextRequest(*q).Execute()
	errors.PanicOnWebError(r, e)

	return device.New(result)
}
