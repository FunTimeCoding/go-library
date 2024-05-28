package system

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"net"
)

func TransmissionSocket(address string) net.Conn {
	result, e := net.Dial(Transmission, address)
	errors.PanicOnError(e)

	return result
}
