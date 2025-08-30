package system_number

import "github.com/netbox-community/go-netbox/v4"

type Number struct {
	Identifier int32
	Name       string

	Raw *netbox.ASN
}
