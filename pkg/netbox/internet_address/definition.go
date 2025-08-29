package internet_address

import "net"

type Definition struct {
	Name    string
	Address net.IP
	Mask    net.IPMask
}
