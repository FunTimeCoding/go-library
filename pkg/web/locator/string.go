package locator

import (
	"github.com/funtimecoding/go-library/pkg/system/join"
	"net"
	"net/url"
	"strconv"
)

func (l *Locator) String() string {
	result := &url.URL{}

	if l.host != "" {
		result.Host = l.host
	}

	if l.port != 0 {
		result.Host = net.JoinHostPort(l.host, strconv.Itoa(l.port))
	}

	if l.scheme != "" {
		result.Scheme = l.scheme
	}

	if l.basePath != "" && l.path != "" {
		result.Path = join.Absolute(l.basePath, l.path)
	} else if l.basePath != "" && l.path == "" {
		result.Path = l.basePath
	} else if l.path != "" {
		result.Path = l.path
	}

	result.RawQuery = l.values.Encode()

	if l.fragment != "" {
		result.Fragment = l.fragment
	}

	return result.String()
}
