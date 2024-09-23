package metric

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
)

type Server struct {
	server   *http.Server
	registry *prometheus.Registry
	port     int
	context  context.Context
	mux      *http.ServeMux
}
