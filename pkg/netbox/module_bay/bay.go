package module_bay

import "github.com/netbox-community/go-netbox/v4"

type Bay struct {
	Identifier  int32
	Name        string
	Description string

	Raw *netbox.ModuleBay
}
