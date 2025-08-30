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
	DeviceRoles       []*device_role.Role
	DeviceTypes       []*device_type.Type
	InternetAddresses []*internet_address.Address
	Manufacturers     []*manufacturer.Manufacturer
	PhysicalAddresses []*physical_address.Address
	Prefixes          []*prefix.Prefix
	Racks             []*rack.Rack
	Sites             []*site.Site
	Tags              []*tag.Tag
	Tenants           []*tenant.Tenant
}
