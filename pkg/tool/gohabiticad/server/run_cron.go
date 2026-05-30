package server

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/convert"
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/generated/server"
)

func (s *Server) RunCron(
	_ context.Context,
	_ server.RunCronRequestObject,
) (server.RunCronResponseObject, error) {
	return server.RunCron200JSONResponse(
		*convert.CronResult(s.habitica.MustCron()),
	), nil
}
