package system

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/strings/separator"
	"github.com/funtimecoding/go-library/pkg/strings/slice"
	"net"
)

func MustLookup(address string) []string {
	result, e := net.LookupAddr(MustCleanAddress(address))
	errors.PanicOnError(e)

	return slice.TrimSuffix(result, separator.Dot)
}
