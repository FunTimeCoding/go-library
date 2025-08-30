package wireless_link

import "github.com/netbox-community/go-netbox/v4"

func New(v *netbox.WirelessLink) *Link {
	return &Link{Identifier: v.GetId(), Name: v.GetDisplay(), Raw: v}
}
