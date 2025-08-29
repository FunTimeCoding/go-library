package contact

import "github.com/netbox-community/go-netbox/v4"

func New(d *netbox.Contact) *Contact {
	return &Contact{Identifier: d.GetId(), Name: d.GetName(), Raw: d}
}
