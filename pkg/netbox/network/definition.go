package network

import "net"

type Definition struct {
	Name            string
	Type            string
	PhysicalAddress net.HardwareAddr
}
