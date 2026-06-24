package server

import (
	"context"
	"errors"
	"github.com/funtimecoding/go-library/pkg/errors/not_found"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) CreateContainerSnapshot(
	_ context.Context,
	r server.CreateContainerSnapshotRequestObject,
) (server.CreateContainerSnapshotResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if e != nil {
		return server.CreateContainerSnapshot400JSONResponse(*clientError(e)), nil
	}

	c, e := s.service.Client(instance)

	if e != nil {
		return server.CreateContainerSnapshot500JSONResponse(*s.captureDetail(e)), nil
	}

	ct, e := findContainer(c, r.Identifier, r.Params.Node)

	if e != nil {
		if errors.Is(e, not_found.Sentinel) {
			return server.CreateContainerSnapshot404JSONResponse{Error: e.Error()}, nil
		}

		return server.CreateContainerSnapshot500JSONResponse(*s.captureDetail(e)), nil
	}

	task, e := c.CreateContainerSnapshot(ct, r.Body.Name)

	if e != nil {
		return server.CreateContainerSnapshot500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.CreateContainerSnapshot200JSONResponse{
		TaskId: string(task.UPID),
	}, nil
}
