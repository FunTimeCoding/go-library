package system

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"net"
)

func IsPortUnused(port int) bool {
	listener, e := net.Listen(
		"tcp",
		fmt.Sprintf("localhost:%d", port),
	)

	if e != nil {
		return false
	}

	errors.LogClose(listener)

	return true
}
