package system

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"net"
)

func UnixSocket(address string) net.Conn {
	result, e := net.Dial("unix", address)
	errors.PanicOnError(e)

	return result
}
