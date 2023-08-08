package system

import (
	"github.com/funtimecoding/go-library/errors"
	"net"
)

func PhysicalAddress(s string) net.HardwareAddr {
	result, e := net.ParseMAC(s)
	errors.PanicOnError(e)

	return result
}
