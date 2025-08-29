package cable

import "github.com/netbox-community/go-netbox/v4"

func New(c *netbox.Cable) *Cable {
	return &Cable{
		Identifier:  c.GetId(),
		Name:        c.GetDisplay(),
		Description: c.GetDescription(),
		Raw:         c,
	}
}
