package site

import "github.com/netbox-community/go-netbox/v4"

type Site struct {
	Identifier int32
	Name       string

	Raw *netbox.Site
}
