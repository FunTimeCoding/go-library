package location

import "github.com/netbox-community/go-netbox/v4"

func New(d *netbox.Location) *Location {
	return &Location{Identifier: d.GetId(), Name: d.GetDisplay(), Raw: d}
}
