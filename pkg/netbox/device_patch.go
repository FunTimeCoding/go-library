package netbox

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/netbox/device"
	"github.com/netbox-community/go-netbox/v4"
)

func devicePatch(d *device.Device) *netbox.PatchedWritableDeviceWithConfigContextRequest {
	result := netbox.NewPatchedWritableDeviceWithConfigContextRequest()
	result.SetName(d.Name)

	if s := d.Raw.Role.GetName(); s != "" {
		result.SetRole(
			netbox.BriefDeviceRoleRequestAsDeviceWithConfigContextRequestRole(
				netbox.NewBriefDeviceRoleRequest(
					s,
					d.Raw.Role.GetSlug(),
				),
			),
		)
	}

	if s := d.Raw.DeviceType.GetModel(); s != "" {
		result.SetDeviceType(
			netbox.BriefDeviceTypeRequestAsDeviceBayTemplateRequestDeviceType(
				netbox.NewBriefDeviceTypeRequest(
					netbox.BriefManufacturerRequestAsBriefDeviceTypeRequestManufacturer(
						netbox.NewBriefManufacturerRequest(
							d.Raw.DeviceType.Manufacturer.GetName(),
							d.Raw.DeviceType.Manufacturer.GetSlug(),
						),
					),
					s,
					d.Raw.DeviceType.GetSlug(),
				),
			),
		)
	}

	if s := d.Raw.Site.GetName(); s != "" {
		result.SetSite(
			netbox.BriefSiteRequestAsDeviceWithConfigContextRequestSite(
				netbox.NewBriefSiteRequest(s, d.Raw.Site.GetSlug()),
			),
		)
	}

	if s := d.Raw.PrimaryIp4.Get().GetAddress(); s != "" {
		result.SetPrimaryIp4(
			netbox.BriefIPAddressRequestAsDeviceWithConfigContextRequestPrimaryIp4(
				netbox.NewBriefIPAddressRequest(s),
			),
		)
	}

	if s := d.Raw.PrimaryIp6.Get().GetAddress(); s != "" {
		result.SetPrimaryIp6(
			netbox.BriefIPAddressRequestAsDeviceWithConfigContextRequestPrimaryIp4(
				netbox.NewBriefIPAddressRequest(s),
			),
		)
	}

	if s := d.Raw.GetSerial(); s != "" {
		if d.Serial == "" {
			fmt.Printf(
				"warning: not unsetting serial %s %s\n",
				d.Name,
				s,
			)
		} else if d.Serial != s {
			result.SetSerial(s)
		}
	}

	return result
}
