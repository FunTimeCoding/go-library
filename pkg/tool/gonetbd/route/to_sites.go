package route

import (
	"github.com/funtimecoding/go-library/pkg/netbox/site"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/server"
)

func toSites(sites []*site.Site) []server.Site {
	result := make([]server.Site, 0, len(sites))

	for _, s := range sites {
		result = append(
			result,
			server.Site{
			Identifier: s.Identifier,
			Name:       s.Name,
		})
	}

	return result
}
