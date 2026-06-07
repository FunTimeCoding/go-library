package server

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/server"
)

func (s *Server) PostSessionExport(
	_ context.Context,
	r server.PostSessionExportRequestObject,
) (server.PostSessionExportResponseObject, error) {
	e := s.service.Resolve(r.Identifier)

	if e.Identifier == "" {
		return server.PostSessionExport404Response{}, nil
	}

	return server.PostSessionExport200JSONResponse{
		Paths: []string{s.exportSession(e, s.sessionExportPath)},
	}, nil
}
