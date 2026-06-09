package server

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/convert"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/server"
)

func (s *Server) ListContainerSnapshots(
	_ context.Context,
	r server.ListContainerSnapshotsRequestObject,
) (server.ListContainerSnapshotsResponseObject, error) {
	instance, e := s.resolveInstance(r.Params.Instance)

	if e != nil {
		return server.ListContainerSnapshots400JSONResponse{ClientErrorJSONResponse: *clientError(e)}, nil
	}

	c, e := s.service.Client(instance)

	if e != nil {
		return server.ListContainerSnapshots500JSONResponse{
			ErrorJSONResponse: *s.captureFail(e, constant.UnexpectedError),
		}, nil
	}

	ct, e := findContainer(c, r.Identifier, r.Params.Node)

	if e != nil {
		return server.ListContainerSnapshots500JSONResponse{
			ErrorJSONResponse: *s.captureFail(e, constant.UnexpectedError),
		}, nil
	}

	if ct == nil {
		return nil, nil
	}

	snapshots, e := c.ContainerSnapshots(ct)

	if e != nil {
		return server.ListContainerSnapshots500JSONResponse{
			ErrorJSONResponse: *s.captureFail(e, constant.UnexpectedError),
		}, nil
	}

	pointers := convert.ContainerSnapshots(snapshots)
	result := make(server.ListContainerSnapshots200JSONResponse, len(pointers))

	for i, snap := range pointers {
		result[i] = *snap
	}

	return result, nil
}
