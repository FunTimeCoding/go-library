package metric

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/funtimecoding/go-library/pkg/web/location"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func New(
	port int,
	verbose bool,
) *Server {
	var address string

	if port == 0 {
		port = 9090
		address = web.MetricsAddress
	} else {
		address = fmt.Sprintf(":%d", port)
	}

	r := prometheus.NewRegistry()
	m := http.NewServeMux()
	m.Handle(
		location.Metrics,
		promhttp.HandlerFor(r, promhttp.HandlerOpts{Registry: r}),
	)

	return &Server{
		verbose:  verbose,
		port:     port,
		context:  context.Background(),
		registry: r,
		server:   &http.Server{Addr: address, Handler: m},
	}
}
