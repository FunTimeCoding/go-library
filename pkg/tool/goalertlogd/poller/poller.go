package poller

import (
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/store"
	"github.com/prometheus/client_golang/prometheus"
	"sync/atomic"
	"time"
)

type Poller struct {
	client    AlertSource
	store     *store.Store
	logger    *logger.Logger
	interval  time.Duration
	retention time.Duration
	firing    map[string]string
	stop      chan struct{}
	lastPoll  atomic.Value
	registry  *prometheus.Registry
	metrics   *metrics
}
