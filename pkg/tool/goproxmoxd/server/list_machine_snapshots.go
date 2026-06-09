package server

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/convert"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) ListMachineSnapshots(
	_ context.Context,
	r server.ListMachineSnapshotsRequestObject,
) (server.ListMachineSnapshotsResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if e != nil {
		return server.ListMachineSnapshots400JSONResponse{ClientErrorJSONResponse: *clientError(e)}, nil
	}

	c, e := s.service.Client(instance)

	if e != nil {
		return server.ListMachineSnapshots500JSONResponse{
			ErrorJSONResponse: *s.captureFail(e, constant.UnexpectedError),
		}, nil
	}

	vm, e := findMachine(c, r.Identifier, r.Params.Node)

	if e != nil {
		return server.ListMachineSnapshots500JSONResponse{
			ErrorJSONResponse: *s.captureFail(e, constant.UnexpectedError),
		}, nil
	}

	if vm == nil {
		return nil, nil
	}

	snapshots, e := c.MachineSnapshots(vm)

	if e != nil {
		return server.ListMachineSnapshots500JSONResponse{
			ErrorJSONResponse: *s.captureFail(e, constant.UnexpectedError),
		}, nil
	}

	pointers := convert.MachineSnapshots(snapshots)
	result := make(server.ListMachineSnapshots200JSONResponse, len(pointers))

	for i, snap := range pointers {
		result[i] = *snap
	}

	return result, nil
}
