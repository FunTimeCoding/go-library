package server

import (
	"context"
	"errors"
	"github.com/funtimecoding/go-library/pkg/errors/not_found"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) StopMachine(
	_ context.Context,
	r server.StopMachineRequestObject,
) (server.StopMachineResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if e != nil {
		return server.StopMachine400JSONResponse(*clientError(e)), nil
	}

	c, e := s.service.Client(instance)

	if e != nil {
		return server.StopMachine500JSONResponse(*s.captureDetail(e)), nil
	}

	vm, e := findMachine(c, r.Identifier, r.Params.Node)

	if e != nil {
		if errors.Is(e, not_found.Sentinel) {
			return server.StopMachine404JSONResponse{Error: e.Error()}, nil
		}

		return server.StopMachine500JSONResponse(*s.captureDetail(e)), nil
	}

	task, e := c.StopMachine(vm)

	if e != nil {
		return server.StopMachine500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.StopMachine200JSONResponse{
		TaskId: string(task.UPID),
	}, nil
}
