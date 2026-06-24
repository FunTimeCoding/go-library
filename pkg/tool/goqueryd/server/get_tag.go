package server

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/generated/server"
)

func (s *Server) GetTag(
	_ context.Context,
	r server.GetTagRequestObject,
) (server.GetTagResponseObject, error) {
	collection := ""

	if r.Params.Collection != nil {
		collection = *r.Params.Collection
	}

	sourceType := s.service.GetSourceType(collection, r.Params.Path)

	return server.GetTag200JSONResponse{
		SourceType: new(string(sourceType)),
	}, nil
}
