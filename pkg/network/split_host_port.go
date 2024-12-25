package network

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/strings"
	"net"
)

func SplitHostPort(s string) (string, int) {
	host, port, e := net.SplitHostPort(s)
	errors.PanicOnError(e)

	return host, strings.ToIntegerStrict(port)
}
