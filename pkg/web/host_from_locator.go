package web

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"net"
	"strings"
)

func HostFromLocator(s string) string {
	hostPort := HostPortFromLocator(s)

	if strings.ContainsRune(hostPort, ':') {
		host, _, e := net.SplitHostPort(hostPort)

		if e != nil {
			if !strings.Contains(e.Error(), "missing port in address") {
				errors.PanicOnError(e)
			}
		}

		return host
	}

	return hostPort
}
