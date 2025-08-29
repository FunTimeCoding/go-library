package contact

import "github.com/netbox-community/go-netbox/v4"

type Contact struct {
	Identifier int32
	Name       string

	Raw *netbox.Contact
}
