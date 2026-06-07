package server

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/server"
)

func (s *Server) PostRelease(
	_ context.Context,
	r server.PostReleaseRequestObject,
) (server.PostReleaseResponseObject, error) {
	identifier := s.service.ResolveByCallsign(r.Body.Callsign)

	if e := s.service.Release(identifier, r.Body.Callsign); e != nil {
		return server.PostRelease500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.PostRelease200Response{}, nil
}
