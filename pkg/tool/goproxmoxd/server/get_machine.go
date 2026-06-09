package server

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/convert"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) GetMachine(
	_ context.Context,
	r server.GetMachineRequestObject,
) (server.GetMachineResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if e != nil {
		return server.GetMachine400JSONResponse{ClientErrorJSONResponse: *clientError(e)}, nil
	}

	c, e := s.service.Client(instance)

	if e != nil {
		return server.GetMachine500JSONResponse{
			ErrorJSONResponse: *s.captureFail(e, constant.UnexpectedError),
		}, nil
	}

	vm, e := findMachine(c, r.Identifier, r.Params.Node)

	if e != nil {
		return server.GetMachine500JSONResponse{
			ErrorJSONResponse: *s.captureFail(e, constant.UnexpectedError),
		}, nil
	}

	if vm == nil {
		return nil, nil
	}

	return server.GetMachine200JSONResponse(
		*convert.MachineDetail(vm),
	), nil
}
