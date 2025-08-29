package service

import "github.com/netbox-community/go-netbox/v4"

type Service struct {
	Identifier int32
	Name       string

	Raw *netbox.Service
}
