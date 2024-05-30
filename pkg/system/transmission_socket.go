package system

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"net"
)

func TransmissionSocket(address string) net.Conn {
	result, e := net.Dial(constant.Transmission, address)
	errors.PanicOnError(e)

	return result
}
