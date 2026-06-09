package server

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) CreateContainerSnapshot(
	_ context.Context,
	r server.CreateContainerSnapshotRequestObject,
) (server.CreateContainerSnapshotResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if e != nil {
		return server.CreateContainerSnapshot400JSONResponse{ClientErrorJSONResponse: *clientError(e)}, nil
	}

	c, e := s.service.Client(instance)

	if e != nil {
		return server.CreateContainerSnapshot500JSONResponse{
			ErrorJSONResponse: *s.captureFail(e, constant.UnexpectedError),
		}, nil
	}

	ct, e := findContainer(c, r.Identifier, r.Params.Node)

	if e != nil {
		return server.CreateContainerSnapshot500JSONResponse{
			ErrorJSONResponse: *s.captureFail(e, constant.UnexpectedError),
		}, nil
	}

	if ct == nil {
		return nil, nil
	}

	task, e := c.CreateContainerSnapshot(ct, r.Body.Name)

	if e != nil {
		return server.CreateContainerSnapshot500JSONResponse{
			ErrorJSONResponse: *s.captureFail(e, constant.UnexpectedError),
		}, nil
	}

	return server.CreateContainerSnapshot200JSONResponse{
		TaskId: string(task.UPID),
	}, nil
}
