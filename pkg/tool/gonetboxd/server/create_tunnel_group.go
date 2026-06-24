package server

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateTunnelGroup(
	_ context.Context,
	r server.CreateTunnelGroupRequestObject,
) (server.CreateTunnelGroupResponseObject, error) {
	g, e := s.client.CreateTunnelGroup(r.Body.Name)

	if e != nil {
		return server.CreateTunnelGroup500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.CreateTunnelGroup201JSONResponse(*convert.TunnelGroup(g)), nil
}
