package prefix

import "github.com/netbox-community/go-netbox/v4"

func New(p *netbox.Prefix) *Prefix {
	return &Prefix{
		Identifier:  p.GetId(),
		Name:        p.GetDisplay(),
		Description: p.GetDescription(),
		Raw:         p,
	}
}
