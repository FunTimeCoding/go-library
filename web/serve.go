package web

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/system"
	"log"
	"net/http"
)

func Serve(
	c context.Context,
	r http.Handler,
	port int,
) {
	s := &http.Server{Addr: fmt.Sprintf(":%d", port), Handler: r}
	go func() {
		if err := s.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	system.KillSignalBlock()
	GracefulShutdown(c, s)
}
