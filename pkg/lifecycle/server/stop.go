package server

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (s *Server) Stop(verbose bool) {
	web.GracefulShutdown(context.Background(), s.http, verbose)
}
