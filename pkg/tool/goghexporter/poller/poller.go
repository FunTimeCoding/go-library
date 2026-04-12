package poller

import (
	"github.com/funtimecoding/go-library/pkg/github"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

type Poller struct {
	client   *github.Client
	owner    string
	interval time.Duration
	gauge    *prometheus.GaugeVec
	stop     chan struct{}
}
