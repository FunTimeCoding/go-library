package server

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) ResetMachine(
	_ context.Context,
	r server.ResetMachineRequestObject,
) (server.ResetMachineResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if e != nil {
		return server.ResetMachine400JSONResponse{ClientErrorJSONResponse: *clientError(e)}, nil
	}

	c, e := s.service.Client(instance)

	if e != nil {
		return server.ResetMachine500JSONResponse{
			ErrorJSONResponse: *s.captureFail(e, constant.UnexpectedError),
		}, nil
	}

	vm, e := findMachine(c, r.Identifier, r.Params.Node)

	if e != nil {
		return server.ResetMachine500JSONResponse{
			ErrorJSONResponse: *s.captureFail(e, constant.UnexpectedError),
		}, nil
	}

	if vm == nil {
		return nil, nil
	}

	task, e := c.ResetMachine(vm)

	if e != nil {
		return server.ResetMachine500JSONResponse{
			ErrorJSONResponse: *s.captureFail(e, constant.UnexpectedError),
		}, nil
	}

	return server.ResetMachine200JSONResponse{
		TaskId: string(task.UPID),
	}, nil
}
