package system

import (
	"github.com/funtimecoding/go-library/pkg/strings/separator"
	"github.com/funtimecoding/go-library/pkg/strings/slice"
	"net"
)

func Lookup(address string) []string {
	result, e := net.LookupAddr(CleanAddressStrict(address))

	if e != nil {
		return []string{}
	}

	return slice.TrimSuffix(result, separator.Dot)
}
