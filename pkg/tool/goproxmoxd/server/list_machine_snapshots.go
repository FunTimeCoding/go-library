package server

import (
	"context"
	"errors"
	"github.com/funtimecoding/go-library/pkg/errors/not_found"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/convert"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) ListMachineSnapshots(
	_ context.Context,
	r server.ListMachineSnapshotsRequestObject,
) (server.ListMachineSnapshotsResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if e != nil {
		return server.ListMachineSnapshots400JSONResponse(*clientError(e)), nil
	}

	c, e := s.service.Client(instance)

	if e != nil {
		return server.ListMachineSnapshots500JSONResponse(*s.captureDetail(e)), nil
	}

	vm, e := findMachine(c, r.Identifier, r.Params.Node)

	if e != nil {
		if errors.Is(e, not_found.Sentinel) {
			return server.ListMachineSnapshots404JSONResponse{Error: e.Error()}, nil
		}

		return server.ListMachineSnapshots500JSONResponse(*s.captureDetail(e)), nil
	}

	snapshots, e := c.MachineSnapshots(vm)

	if e != nil {
		return server.ListMachineSnapshots500JSONResponse(*s.captureDetail(e)), nil
	}

	pointers := convert.MachineSnapshots(snapshots)
	result := make(server.ListMachineSnapshots200JSONResponse, len(pointers))

	for i, snap := range pointers {
		result[i] = *snap
	}

	return result, nil
}
