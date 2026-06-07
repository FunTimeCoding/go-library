package server

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/server"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/sweep"
)

func (s *Server) PostSweep(
	_ context.Context,
	_ server.PostSweepRequestObject,
) (server.PostSweepResponseObject, error) {
	result := sweep.Run(s.harborPath)

	return server.PostSweep200JSONResponse{
		Copied:  result.Copied,
		Updated: result.Updated,
		Skipped: result.Skipped,
	}, nil
}
