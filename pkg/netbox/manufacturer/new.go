package manufacturer

import "github.com/netbox-community/go-netbox/v4"

func New(d *netbox.Manufacturer) *Manufacturer {
	return &Manufacturer{Identifier: d.GetId(), Name: d.GetName(), Raw: d}
}
