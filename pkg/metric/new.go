package metric

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func New(
	port int,
	verbose bool,
) *Server {
	if port == 0 {
		port = 9090
	}

	r := prometheus.NewRegistry()
	m := http.NewServeMux()
	m.Handle(
		"/metrics",
		promhttp.HandlerFor(
			r,
			promhttp.HandlerOpts{Registry: r},
		),
	)

	return &Server{
		verbose:  verbose,
		port:     port,
		context:  context.Background(),
		registry: r,
		server: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: m,
		},
	}
}
