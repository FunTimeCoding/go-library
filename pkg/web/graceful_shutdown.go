package web

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/errors"
	"log"
	"net/http"
	"time"
)

func GracefulShutdown(
	c context.Context,
	s *http.Server,
	verbose bool,
) {
	if verbose {
		log.Println("shutdown with timeout")
	}

	o, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()
	errors.PanicOnError(s.Shutdown(o))
	<-o.Done()

	if verbose {
		log.Println("shutdown complete")
	}
}
