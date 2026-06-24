package server

import (
	"context"
	"errors"
	"github.com/funtimecoding/go-library/pkg/errors/not_found"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) RollbackContainerSnapshot(
	_ context.Context,
	r server.RollbackContainerSnapshotRequestObject,
) (server.RollbackContainerSnapshotResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if e != nil {
		return server.RollbackContainerSnapshot400JSONResponse(*clientError(e)), nil
	}

	c, e := s.service.Client(instance)

	if e != nil {
		return server.RollbackContainerSnapshot500JSONResponse(*s.captureDetail(e)), nil
	}

	ct, e := findContainer(c, r.Identifier, r.Params.Node)

	if e != nil {
		if errors.Is(e, not_found.Sentinel) {
			return server.RollbackContainerSnapshot404JSONResponse{Error: e.Error()}, nil
		}

		return server.RollbackContainerSnapshot500JSONResponse(*s.captureDetail(e)), nil
	}

	task, e := c.RollbackContainerSnapshot(ct, r.Name)

	if e != nil {
		return server.RollbackContainerSnapshot500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.RollbackContainerSnapshot200JSONResponse{
		TaskId: string(task.UPID),
	}, nil
}
