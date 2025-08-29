package service

import "github.com/netbox-community/go-netbox/v4"

func New(d *netbox.Service) *Service {
	return &Service{Identifier: d.GetId(), Name: d.GetName(), Raw: d}
}
