package server

import (
	"context"
	"errors"
	"github.com/funtimecoding/go-library/pkg/errors/not_found"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/convert"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) ListContainerSnapshots(
	_ context.Context,
	r server.ListContainerSnapshotsRequestObject,
) (server.ListContainerSnapshotsResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if e != nil {
		return server.ListContainerSnapshots400JSONResponse(*clientError(e)), nil
	}

	c, e := s.service.Client(instance)

	if e != nil {
		return server.ListContainerSnapshots500JSONResponse(*s.captureDetail(e)), nil
	}

	ct, e := findContainer(c, r.Identifier, r.Params.Node)

	if e != nil {
		if errors.Is(e, not_found.Sentinel) {
			return server.ListContainerSnapshots404JSONResponse{Error: e.Error()}, nil
		}

		return server.ListContainerSnapshots500JSONResponse(*s.captureDetail(e)), nil
	}

	snapshots, e := c.ContainerSnapshots(ct)

	if e != nil {
		return server.ListContainerSnapshots500JSONResponse(*s.captureDetail(e)), nil
	}

	pointers := convert.ContainerSnapshots(snapshots)
	result := make(
		server.ListContainerSnapshots200JSONResponse,
		len(pointers),
	)

	for i, snap := range pointers {
		result[i] = *snap
	}

	return result, nil
}
