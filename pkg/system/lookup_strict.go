package system

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"net"
)

func LookupStrict(address string) []string {
	result, e := net.LookupAddr(address)
	errors.PanicOnError(e)

	return result
}
