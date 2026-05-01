package convert

import (
	"github.com/funtimecoding/go-library/pkg/netbox/site"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/generated/server"
)

func Site(s *site.Site) server.Site {
	return server.Site{Identifier: s.Identifier, Name: s.Name}
}
