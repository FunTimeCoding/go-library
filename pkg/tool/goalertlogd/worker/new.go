package worker

import (
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/store"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

func New(
	a AlertSource,
	s *store.Store,
	l *logger.Logger,
	interval time.Duration,
	retention time.Duration,
	r *prometheus.Registry,
) *Worker {
	p := &Worker{
		client:    a,
		store:     s,
		logger:    l,
		interval:  interval,
		retention: retention,
		firing:    make(map[string]string),
		stop:      make(chan struct{}),
		registry:  r,
	}

	if r != nil {
		p.metrics = newMetrics(r)
	}

	return p
}
