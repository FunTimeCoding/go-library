package secure_shell

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"net"
)

func Accept(l net.Listener) net.Conn {
	result, e := l.Accept()
	errors.PanicOnError(e)

	return result
}
