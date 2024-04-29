package system

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/separator"
	"github.com/funtimecoding/go-library/pkg/strings"
	"net"
)

func LookupStrict(address string) []string {
	result, e := net.LookupAddr(CleanAddressStrict(address))
	errors.PanicOnError(e)

	return strings.SliceTrimSuffix(result, separator.Dot)
}
