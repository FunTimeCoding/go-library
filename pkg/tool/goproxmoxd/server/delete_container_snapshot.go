package server

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) DeleteContainerSnapshot(
	_ context.Context,
	r server.DeleteContainerSnapshotRequestObject,
) (server.DeleteContainerSnapshotResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if e != nil {
		return server.DeleteContainerSnapshot400JSONResponse{ClientErrorJSONResponse: *clientError(e)}, nil
	}

	c, e := s.service.Client(instance)

	if e != nil {
		return server.DeleteContainerSnapshot500JSONResponse{
			ErrorJSONResponse: *s.captureFail(e, constant.UnexpectedError),
		}, nil
	}

	ct, e := findContainer(c, r.Identifier, r.Params.Node)

	if e != nil {
		return server.DeleteContainerSnapshot500JSONResponse{
			ErrorJSONResponse: *s.captureFail(e, constant.UnexpectedError),
		}, nil
	}

	if ct == nil {
		return nil, nil
	}

	task, e := c.DeleteContainerSnapshot(ct, r.Name)

	if e != nil {
		return server.DeleteContainerSnapshot500JSONResponse{
			ErrorJSONResponse: *s.captureFail(e, constant.UnexpectedError),
		}, nil
	}

	return server.DeleteContainerSnapshot200JSONResponse{
		TaskId: string(task.UPID),
	}, nil
}
