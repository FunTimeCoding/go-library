package internet_address

import (
	"github.com/netbox-community/go-netbox/v4"
	"net"
)

type Address struct {
	Identifier       int32
	Name             string
	Address          net.IP
	Network          *net.IPNet
	ObjectType       string
	ObjectIdentifier int64

	Raw *netbox.IPAddress
}
