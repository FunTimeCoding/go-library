package system_number

import "github.com/netbox-community/go-netbox/v4"

func New(v *netbox.ASN) *Number {
	return &Number{Identifier: v.GetId(), Name: v.GetDisplay(), Raw: v}
}
