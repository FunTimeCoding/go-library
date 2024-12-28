package metric

import (
	"github.com/slok/go-http-metrics/metrics/prometheus"
	"github.com/slok/go-http-metrics/middleware"
	"github.com/slok/go-http-metrics/middleware/std"
	"net/http"
)

func MiddlewareServer(
	address string,
	m http.Handler,
) *http.Server {
	return &http.Server{
		Addr: address,
		Handler: std.Handler(
			"",
			middleware.New(
				middleware.Config{
					Recorder: prometheus.NewRecorder(prometheus.Config{}),
				},
			),
			m,
		),
	}
}
