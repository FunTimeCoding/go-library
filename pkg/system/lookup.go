package system

import (
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/funtimecoding/go-library/pkg/strings/separator"
	"net"
)

func Lookup(address string) []string {
	result, e := net.LookupAddr(CleanAddressStrict(address))

	if e != nil {
		return []string{}
	}

	return strings.SliceTrimSuffix(result, separator.Dot)
}
