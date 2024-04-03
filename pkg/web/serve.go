package web

import (
	"context"
	"errors"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/system"
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
		if e := s.ListenAndServe(); e != nil {
			if !errors.Is(e, http.ErrServerClosed) {
				log.Fatalf("listen: %s\n", e)
			}
		}
	}()
	system.KillSignalBlock()
	GracefulShutdown(c, s)
}
