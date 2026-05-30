package server

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/convert"
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/generated/server"
)

func (s *Server) GetStats(
	_ context.Context,
	_ server.GetStatsRequestObject,
) (server.GetStatsResponseObject, error) {
	return server.GetStats200JSONResponse(
		*convert.Stats(s.habitica.MustUserStats()),
	), nil
}
