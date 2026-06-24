package server

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/generated/server"
)

func (s *Server) PostContext(
	_ context.Context,
	r server.PostContextRequestObject,
) (server.PostContextResponseObject, error) {
	s.service.AddContext(
		r.Body.Collection,
		r.Body.PathPrefix,
		r.Body.Description,
	)

	return server.PostContext200Response{}, nil
}
