package server

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListTunnelGroups(
	_ context.Context,
	_ server.ListTunnelGroupsRequestObject,
) (server.ListTunnelGroupsResponseObject, error) {
	groups, e := s.client.TunnelGroups()

	if e != nil {
		return server.ListTunnelGroups500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.ListTunnelGroups200JSONResponse(
		convert.TunnelGroups(groups),
	), nil
}
