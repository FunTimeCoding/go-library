package service

import "github.com/netbox-community/go-netbox/v4"

func New(v *netbox.Service) *Service {
	return &Service{Identifier: v.GetId(), Name: v.GetName(), Raw: v}
}
