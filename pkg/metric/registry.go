package metric

import "github.com/prometheus/client_golang/prometheus"

func (s *Server) Registry() *prometheus.Registry {
	return s.registry
}
