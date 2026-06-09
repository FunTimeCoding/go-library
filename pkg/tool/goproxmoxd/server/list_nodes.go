package server

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/convert"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) ListNodes(
	_ context.Context,
	r server.ListNodesRequestObject,
) (server.ListNodesResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if e != nil {
		return server.ListNodes400JSONResponse{ClientErrorJSONResponse: *clientError(e)}, nil
	}

	c, e := s.service.Client(instance)

	if e != nil {
		return server.ListNodes500JSONResponse{
			ErrorJSONResponse: *s.captureFail(e, constant.UnexpectedError),
		}, nil
	}

	nodes, e := c.Nodes()

	if e != nil {
		return server.ListNodes500JSONResponse{
			ErrorJSONResponse: *s.captureFail(e, constant.UnexpectedError),
		}, nil
	}

	pointers := convert.Nodes(nodes)
	result := make(server.ListNodes200JSONResponse, len(pointers))

	for i, n := range pointers {
		result[i] = *n
	}

	return result, nil
}
