package cache

import (
	"github.com/funtimecoding/go-library/pkg/netbox/device_role"
	"github.com/funtimecoding/go-library/pkg/netbox/device_type"
	"github.com/funtimecoding/go-library/pkg/netbox/internet_address"
	"github.com/funtimecoding/go-library/pkg/netbox/manufacturer"
	"github.com/funtimecoding/go-library/pkg/netbox/physical_address"
	"github.com/funtimecoding/go-library/pkg/netbox/prefix"
	"github.com/funtimecoding/go-library/pkg/netbox/rack"
	"github.com/funtimecoding/go-library/pkg/netbox/site"
	"github.com/funtimecoding/go-library/pkg/netbox/tag"
	"github.com/funtimecoding/go-library/pkg/netbox/tenant"
)

type Cache struct {
	DeviceTypes       []*device_type.Type
	DeviceRoles       []*device_role.Role
	Tags              []*tag.Tag
	Tenants           []*tenant.Tenant
	Sites             []*site.Site
	Prefixes          []*prefix.Prefix
	PhysicalAddresses []*physical_address.Address
	InternetAddresses []*internet_address.Address
	Racks             []*rack.Rack
	Manufacturers     []*manufacturer.Manufacturer
}
