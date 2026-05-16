package convert

import (
	"github.com/funtimecoding/go-library/pkg/netbox/tunnel_group"
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/generated/server"
)

func TunnelGroup(g *tunnel_group.Group) *server.TunnelGroup {
	return &server.TunnelGroup{Identifier: g.Identifier, Name: g.Name}
}
