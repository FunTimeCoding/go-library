package server

import (
	"context"
	"errors"
	"github.com/funtimecoding/go-library/pkg/errors/not_found"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) CreateMachineSnapshot(
	_ context.Context,
	r server.CreateMachineSnapshotRequestObject,
) (server.CreateMachineSnapshotResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if e != nil {
		return server.CreateMachineSnapshot400JSONResponse(*clientError(e)), nil
	}

	c, e := s.service.Client(instance)

	if e != nil {
		return server.CreateMachineSnapshot500JSONResponse(*s.captureDetail(e)), nil
	}

	vm, e := findMachine(c, r.Identifier, r.Params.Node)

	if e != nil {
		if errors.Is(e, not_found.Sentinel) {
			return server.CreateMachineSnapshot404JSONResponse{Error: e.Error()}, nil
		}

		return server.CreateMachineSnapshot500JSONResponse(*s.captureDetail(e)), nil
	}

	task, e := c.CreateMachineSnapshot(vm, r.Body.Name)

	if e != nil {
		return server.CreateMachineSnapshot500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.CreateMachineSnapshot200JSONResponse{
		TaskId: string(task.UPID),
	}, nil
}
