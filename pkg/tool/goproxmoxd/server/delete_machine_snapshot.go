package server

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) DeleteMachineSnapshot(
	_ context.Context,
	r server.DeleteMachineSnapshotRequestObject,
) (server.DeleteMachineSnapshotResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if e != nil {
		return server.DeleteMachineSnapshot400JSONResponse{ClientErrorJSONResponse: *clientError(e)}, nil
	}

	c, e := s.service.Client(instance)

	if e != nil {
		return server.DeleteMachineSnapshot500JSONResponse{
			ErrorJSONResponse: *s.captureFail(e, constant.UnexpectedError),
		}, nil
	}

	vm, e := findMachine(c, r.Identifier, r.Params.Node)

	if e != nil {
		return server.DeleteMachineSnapshot500JSONResponse{
			ErrorJSONResponse: *s.captureFail(e, constant.UnexpectedError),
		}, nil
	}

	if vm == nil {
		return nil, nil
	}

	task, e := c.DeleteMachineSnapshot(vm, r.Name)

	if e != nil {
		return server.DeleteMachineSnapshot500JSONResponse{
			ErrorJSONResponse: *s.captureFail(e, constant.UnexpectedError),
		}, nil
	}

	return server.DeleteMachineSnapshot200JSONResponse{
		TaskId: string(task.UPID),
	}, nil
}
