package web

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/system"
	"net/http"
)

func Serve(
	c context.Context,
	r http.Handler,
	port int,
) {
	s := &http.Server{Addr: fmt.Sprintf(":%d", port), Handler: r}
	ServeAsynchronous(s)
	system.KillSignalBlock()
	GracefulShutdown(c, s)
}
