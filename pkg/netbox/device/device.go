package device

import "github.com/netbox-community/go-netbox/v4"

type Device struct {
	Identifier     int32
	Name           string
	Role           string
	Type           string
	Site           string
	Tenant         string
	Comment        string
	PrimaryAddress string
	Tags           []string
	Link           string

	Raw *netbox.DeviceWithConfigContext
}
