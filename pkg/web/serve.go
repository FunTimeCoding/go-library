package web

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/system"
	"net/http"
)

func Serve(
	c context.Context,
	r http.Handler,
	port int,
	verbose bool,
) {
	s := ServerPort(r, port)
	ServeAsynchronous(s)
	system.KillSignalBlock()
	GracefulShutdown(c, s, verbose)
}
