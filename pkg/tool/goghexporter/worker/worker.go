package worker

import (
	"github.com/funtimecoding/go-library/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/go-library/pkg/github"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

type Worker struct {
	client   *github.Client
	owner    string
	interval time.Duration
	gauge    *prometheus.GaugeVec
	recovery *recovery.Recovery
	stop     chan struct{}
}
