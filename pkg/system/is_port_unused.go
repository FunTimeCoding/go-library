package system

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"net"
)

func IsPortUnused(port int) bool {
	listener, e := net.Listen(
		constant.Transmission,
		fmt.Sprintf("localhost:%d", port),
	)

	if e != nil {
		return false
	}

	errors.LogClose(listener)

	return true
}
