package contact

import "github.com/netbox-community/go-netbox/v4"

func New(v *netbox.Contact) *Contact {
	return &Contact{Identifier: v.GetId(), Name: v.GetName(), Raw: v}
}
