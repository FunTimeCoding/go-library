package server

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/convert"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) GetContainer(
	_ context.Context,
	r server.GetContainerRequestObject,
) (server.GetContainerResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if e != nil {
		return server.GetContainer400JSONResponse{ClientErrorJSONResponse: *clientError(e)}, nil
	}

	c, e := s.service.Client(instance)

	if e != nil {
		return server.GetContainer500JSONResponse{
			ErrorJSONResponse: *s.captureFail(e, constant.UnexpectedError),
		}, nil
	}

	ct, e := findContainer(c, r.Identifier, r.Params.Node)

	if e != nil {
		return server.GetContainer500JSONResponse{
			ErrorJSONResponse: *s.captureFail(e, constant.UnexpectedError),
		}, nil
	}

	if ct == nil {
		return nil, nil
	}

	return server.GetContainer200JSONResponse(
		*convert.ContainerDetail(ct),
	), nil
}
