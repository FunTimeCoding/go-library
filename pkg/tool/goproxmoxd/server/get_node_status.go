package server

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/convert"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) GetNodeStatus(
	_ context.Context,
	r server.GetNodeStatusRequestObject,
) (server.GetNodeStatusResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if e != nil {
		return server.GetNodeStatus400JSONResponse(*clientError(e)), nil
	}

	c, e := s.service.Client(instance)

	if e != nil {
		return server.GetNodeStatus500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	result, e := c.NodeStatus(r.Name)

	if e != nil {
		return server.GetNodeStatus500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	return server.GetNodeStatus200JSONResponse(*convert.NodeStatus(result)), nil
}
