package web

import (
	"github.com/funtimecoding/go-library/pkg/integers"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/funtimecoding/go-library/pkg/strings/separator"
)

func AddressHostPort(host string, port int) string {
	return join.Empty(host, separator.Colon, integers.ToString(port))
}
