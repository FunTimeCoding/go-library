package netbox

import "github.com/netbox-community/go-netbox/v4"

func WithInterfaceTypes(v []netbox.InterfaceTypeValue) OptionFunc {
	return func(c *Client) {
		c.interfaceTypes = append(c.interfaceTypes, v...)
	}
}
