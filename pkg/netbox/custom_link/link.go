package custom_link

import "github.com/netbox-community/go-netbox/v4"

type Link struct {
	Identifier int32
	Name       string

	Raw *netbox.CustomLink
}
