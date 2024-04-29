package system

import (
	"github.com/funtimecoding/go-library/pkg/separator"
	"github.com/funtimecoding/go-library/pkg/strings"
	"net"
)

func Lookup(address string) []string {
	result, e := net.LookupAddr(CleanAddressStrict(address))

	if e != nil {
		return []string{}
	}

	return strings.SliceTrimSuffix(result, separator.Dot)
}
