package system

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/strings/split/key_value"
	"net"
	"strings"
)

func CleanAddress(s string) string {
	if strings.ContainsRune(s, ':') {
		host, _, e := net.SplitHostPort(s)
		errors.PanicOnError(e)

		s = host
	}

	if strings.ContainsRune(s, '/') {
		host, _ := key_value.Equals(s)

		s = host
	}

	return s
}
