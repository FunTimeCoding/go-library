package network

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"net"
)

func PhysicalAddress(s string) net.HardwareAddr {
	result, e := net.ParseMAC(s)
	errors.PanicOnError(e)

	return result
}
