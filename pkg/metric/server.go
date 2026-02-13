package metric

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"sync"
)

type Server struct {
	server   *http.Server
	registry *prometheus.Registry
	port     int
	context  context.Context
	verbose  bool
	wait     *sync.WaitGroup
}
