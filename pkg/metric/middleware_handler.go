package metric

import (
	"github.com/slok/go-http-metrics/metrics/prometheus"
	"github.com/slok/go-http-metrics/middleware"
	"github.com/slok/go-http-metrics/middleware/std"
	"net/http"
)

func MiddlewareHandler(h http.Handler) http.Handler {
	return std.Handler(
		"",
		middleware.New(
			middleware.Config{
				Recorder: prometheus.NewRecorder(prometheus.Config{}),
			},
		),
		h,
	)
}
