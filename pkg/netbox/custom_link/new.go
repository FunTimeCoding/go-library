package custom_link

import "github.com/netbox-community/go-netbox/v4"

func New(v *netbox.CustomLink) *Link {
	return &Link{Identifier: v.GetId(), Name: v.GetName(), Raw: v}
}
