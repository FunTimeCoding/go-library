package convert

import (
	"github.com/funtimecoding/go-library/pkg/netbox/site"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/generated/server"
)

func Sites(v []*site.Site) []server.Site {
	result := make([]server.Site, 0, len(v))

	for _, s := range v {
		result = append(result, Site(s))
	}

	return result
}
