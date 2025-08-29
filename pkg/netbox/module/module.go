package module

import "github.com/netbox-community/go-netbox/v4"

type Module struct {
	Identifier  int32
	Name        string
	Description string

	Raw *netbox.Module
}
